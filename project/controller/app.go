package controller

import (
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/dto"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

//APPControllerRegister admin路由注册
func APPRegister(router *gin.RouterGroup) {
	admin := APPController{}
	router.GET("/app_list", admin.APPList)
	router.GET("/app_detail", admin.APPDetail)
	router.GET("/app_stat", admin.AppStatistics)
	router.GET("/app_delete", admin.APPDelete)
	router.POST("/app_add", admin.AppAdd)
	router.POST("/app_update", admin.AppUpdate)
}

type APPController struct {
}

func (admin *APPController) APPList(c *gin.Context) {
	params := &dto.APPListInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	info := &dao.App{}
	list, total, err := info.APPList(c, lib.GORMDefaultPool, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	outputList := []dto.APPListItemOutput{}
	for _, item := range list {
		serviceCounter, _ := public.FlowCounterHandler.GetCounter(public.FlowAPPPrefix + item.AppID)
		realQps := serviceCounter.GetQPS()
		realQpd, _ := serviceCounter.GetDayCount(time.Now())
		outputList = append(outputList, dto.APPListItemOutput{
			ID:       item.ID,
			AppID:    item.AppID,
			Name:     item.Name,
			Secret:   item.Secret,
			WhiteIPS: item.WhiteIPS,
			Qpd:      item.Qpd,
			Qps:      item.Qps,
			RealQpd:  realQpd,
			RealQps:  realQps,
		})
	}
	middleware.ResponseSuccess(c, map[string]interface{}{"list": outputList, "total": total})
	return
}

func (admin *APPController) APPDetail(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	detail, err := search.Find(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, detail)
	return
}

func (admin *APPController) APPDelete(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	info, err := search.Find(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	info.IsDelete = 1
	if err := info.Save(c, lib.GORMDefaultPool); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *APPController) AppAdd(c *gin.Context) {
	params := &dto.APPAddHttpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//验证service_name是否被占用
	search := &dao.App{
		AppID: params.AppID,
	}
	if _, err := search.Find(c, lib.GORMDefaultPool, search); err == nil {
		middleware.ResponseError(c, 2002, errors.New("租户ID被占用，请重新输入"))
		return
	}
	if params.Secret == "" {
		params.Secret = public.MD5(params.AppID)
	}
	tx := lib.GORMDefaultPool
	info := &dao.App{
		AppID:    params.AppID,
		Name:     params.Name,
		Secret:   params.Secret,
		WhiteIPS: params.WhiteIPS,
		Qps:      params.Qps,
		Qpd:      params.Qpd,
	}
	if err := info.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *APPController) AppUpdate(c *gin.Context) {
	params := &dto.APPUpdateHttpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.App{
		ID: params.ID,
	}
	info, err := search.Find(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	if params.Secret == "" {
		params.Secret = public.MD5(params.AppID)
	}
	info.Name = params.Name
	info.Secret = params.Secret
	info.WhiteIPS = params.WhiteIPS
	info.Qps = params.Qps
	info.Qpd = params.Qpd
	if err := info.Save(c, lib.GORMDefaultPool); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *APPController) AppStatistics(c *gin.Context) {
	params := &dto.APPDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	search := &dao.App{
		ID: params.ID,
	}
	detail, err := search.Find(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	counter, _ := public.FlowCounterHandler.GetCounter(public.FlowAPPPrefix + detail.AppID)

	//今日流量全天小时级访问统计
	todayStat := []int64{}
	for i := 0; i <= time.Now().In(lib.TimeLocation).Hour(); i++ {
		nowTime := time.Now()
		nowTime = time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		hourStat, _ := counter.GetHourCount(nowTime)
		todayStat = append(todayStat, hourStat)
	}

	//昨日流量全天小时级访问统计
	yesterdayStat := []int64{}
	for i := 0; i <= 23; i++ {
		nowTime := time.Now().AddDate(0, 0, -1)
		nowTime = time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), i, 0, 0, 0, lib.TimeLocation)
		hourStat, _ := counter.GetHourCount(nowTime)
		yesterdayStat = append(yesterdayStat, hourStat)
	}
	//yesterdayStat = []int64{
	//	12,
	//	20,
	//	23,
	//	57,
	//	25,
	//	48,
	//	76,
	//	69,
	//	140,
	//	200,
	//	250,
	//	345,
	//	500,
	//	550,
	//	780,
	//	670,
	//	650,
	//	500,
	//	488,
	//	480,
	//	440,
	//	360,
	//	200,
	//	105,
	//}
	//todayStat = []int64{
	//	5,
	//	10,
	//	20,
	//	48,
	//	50,
	//	55,
	//	60,
	//	80,
	//	100,
	//	180,
	//	200,
	//}
	middleware.ResponseSuccess(c, map[string][]int64{
		"today":     todayStat,
		"yesterday": yesterdayStat,
	})
	return
}

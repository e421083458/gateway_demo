package controller

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/dto"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"time"
)

//AdminRegister admin路由注册
func DashBoardRegister(router *gin.RouterGroup) {
	admin := DashBoardController{}
	router.GET("/panel_group_data", admin.PanelGroupData)
	router.GET("/flow_stat", admin.FlowStat)
	router.GET("/service_stat", admin.ServiceStat)
}

type DashBoardController struct {
}

func (admin *DashBoardController) PanelGroupData(c *gin.Context) {
	counter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
	if err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}

	serviceInfo := &dao.ServiceInfo{}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 1002, err)
		return
	}

	_, serviceNum, err := serviceInfo.ServiceList(c, tx, (&dto.ServiceListInput{PageNo: 1, PageSize: 1}))
	if err != nil {
		middleware.ResponseError(c, 1003, err)
		return
	}

	app := &dao.App{}
	_, appNum, err := app.APPList(c, tx, (&dto.APPListInput{PageNo: 1, PageSize: 1}))
	if err != nil {
		middleware.ResponseError(c, 1004, err)
		return
	}
	dayCount, _ := counter.GetDayCount(time.Now())
	middleware.ResponseSuccess(c, map[string]interface{}{
		"serviceNum":      serviceNum,
		"todayRequestNum": dayCount,
		"currentQps":      counter.GetQPS(),
		"appNum":          appNum,
	})
	return
}

func (admin *DashBoardController) FlowStat(c *gin.Context) {
	counter, _ := public.FlowCounterHandler.GetCounter(public.FlowTotal)

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
	middleware.ResponseSuccess(c, map[string][]int64{
		"today":     todayStat,
		"yesterday": yesterdayStat,
	})
	return
}

func (admin *DashBoardController) ServiceStat(c *gin.Context) {
	serviceInfo := &dao.ServiceInfo{}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 1002, err)
		return
	}

	loadTypes, err := serviceInfo.ServiceLoadType(c, tx)
	if err != nil {
		middleware.ResponseError(c, 1003, err)
		return
	}

	fmt.Println("loadTypes", loadTypes)
	ServiceStats := dto.ServiceStatOutput{}
	for _, loadType := range loadTypes {
		ServiceStats.Legend = append(ServiceStats.Legend, public.LoadTypeMap[loadType.LoadType])
		ServiceStats.Data = append(ServiceStats.Data, dto.ServiceStatItemOutput{
			Value: loadType.Num,
			Name:  public.LoadTypeMap[loadType.LoadType],
		})
	}
	middleware.ResponseSuccess(c, ServiceStats)
	return
}

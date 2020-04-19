package controller

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/dto"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"time"
)

//ServiceControllerRegister admin路由注册
func ServiceRegister(router *gin.RouterGroup) {
	admin := ServiceController{}
	router.GET("/service_list", admin.ServiceList)
	router.GET("/service_detail", admin.ServiceDetail)
	router.GET("/service_stat", admin.ServiceStatistics)
	router.GET("/service_lb", admin.ServiceLoadBalance)
	router.POST("/service_endpoint_open", admin.ServiceEndpointOpen)
	router.POST("/service_endpoint_close", admin.ServiceEndpointClose)
	router.GET("/service_delete", admin.ServiceDelete)
	router.POST("/service_add_http", admin.ServiceAddHttp)
	router.POST("/service_add_tcp", admin.ServiceAddTcp)
	router.POST("/service_add_grpc", admin.ServiceAddGrpc)
	router.POST("/service_update_http", admin.ServiceUpdateHttp)
	router.POST("/service_update_tcp", admin.ServiceUpdateTcp)
	router.POST("/service_update_grpc", admin.ServiceUpdateGrpc)
}

type ServiceController struct {
}

func (admin *ServiceController) ServiceList(c *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	roleInfo := &dao.ServiceInfo{}
	list, total, err := roleInfo.ServiceList(c, lib.GORMDefaultPool, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	outputList := []dto.ServiceListItemOutput{}
	for _, item := range list {
		detail, err := item.ServiceDetail(c, lib.GORMDefaultPool, &dao.ServiceInfo{
			ID: item.ID,
		})
		if err != nil {
			middleware.ResponseError(c, 2003, err)
			return
		}

		totalNode := len(detail.LoadBalance.GetIPListByModel())
		serviceCounter, _ := public.FlowCounterHandler.GetCounter(public.FlowCountServicePrefix+item.ServiceName)
		qps := serviceCounter.GetQPS()
		qpd, _ := serviceCounter.GetDayCount(time.Now())

		serviceIP := lib.GetStringConf("base.cluster.cluster_ip")
		servicePort := lib.GetStringConf("base.cluster.cluster_port")
		serviceSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")
		serviceHttpBaseURL := serviceIP + ":" + servicePort
		if detail.HttpRule.NeedHttps == 1 {
			serviceHttpBaseURL = serviceIP + ":" + serviceSSLPort
		}
		serviceAddr := "unknow"
		if item.LoadType == public.LoadTypeHTTP && detail.HttpRule.RuleType == 0 {
			serviceAddr = fmt.Sprintf("%s%s", serviceHttpBaseURL, detail.HttpRule.Rule)
		}
		if item.LoadType == public.LoadTypeHTTP && detail.HttpRule.RuleType == 1 {
			serviceAddr = detail.HttpRule.Rule
		}
		if item.LoadType == public.LoadTypeTCP {
			serviceAddr = fmt.Sprintf("%s:%d", serviceIP, detail.TcpRule.Port)
		}
		if item.LoadType == public.LoadTypeGRPC {
			serviceAddr = fmt.Sprintf("%s:%d", serviceIP, detail.GrpcRule.Port)
		}
		outputList = append(outputList, dto.ServiceListItemOutput{
			ID:          item.ID,
			LoadType:    item.LoadType,
			ServiceName: item.ServiceName,
			ServiceDesc: item.ServiceDesc,
			UpdatedAt:   item.UpdatedAt,
			CreatedAt:   item.CreatedAt,
			QPS:         qps,
			QPD:         qpd,
			TotalNode:   totalNode,
			ServiceAddr: serviceAddr,
		})
	}
	middleware.ResponseSuccess(c, map[string]interface{}{"list": outputList, "total": total})
	return
}

func (admin *ServiceController) ServiceDetail(c *gin.Context) {
	params := &dto.ServiceDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := search.ServiceDetail(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, detail)
	return
}

func (admin *ServiceController) ServiceStatistics(c *gin.Context) {
	params := &dto.ServiceDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	search := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := search.ServiceDetail(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	counter, _ := public.FlowCounterHandler.GetCounter(public.FlowCountServicePrefix + detail.Info.ServiceName)

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
	//	120,
	//	50,
	//	10,
	//	57,
	//	59,
	//	48,
	//	76,
	//	69,
	//	200,
	//	400,
	//	580,
	//	1500,
	//	2500,
	//	2300,
	//	1300,
	//	1700,
	//	1900,
	//	1000,
	//	800,
	//	570,
	//	500,
	//	360,
	//	200,
	//	105,
	//}
	//todayStat = []int64{
	//	78,
	//	23,
	//	78,
	//	123,
	//	325,
	//	378,
	//	456,
	//	478,
	//	500,
	//	800,
	//	760,
	//}
	middleware.ResponseSuccess(c, map[string][]int64{
		"today":     todayStat,
		"yesterday": yesterdayStat,
	})
	return
}

func (admin *ServiceController) ServiceLoadBalance(c *gin.Context) {
	params := &dto.ServiceDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := search.ServiceDetail(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	serviceIPList := detail.LoadBalance.GetIPListByModel()
	forbidIPList := detail.LoadBalance.GetForbidListByModel()
	serviceWeightList := detail.LoadBalance.GetWeightListByModel()
	//activeIPList := service.SysConfMgr.GetActiveIPList(detail.Info.ServiceName)
	middleware.ResponseSuccess(c, map[string][]string{
		"service_ip_list":     serviceIPList,
		"forbid_ip_list":      forbidIPList,
		"service_weight_list": serviceWeightList,
		//"active_ip_list":      activeIPList,
	})
	return
}

func (admin *ServiceController) ServiceDelete(c *gin.Context) {
	params := &dto.ServiceDetailInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.ServiceInfo{
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

func (admin *ServiceController) ServiceEndpointOpen(c *gin.Context) {
	params := &dto.ServiceEndpointCloseInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := search.ServiceDetail(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	if detail.LoadBalance == nil {
		middleware.ResponseError(c, 2003, errors.New("负载节点为空"))
		return
	}
	confIPList := detail.LoadBalance.GetIPListByModel()
	if !public.InStringList(params.Addr, confIPList) {
		middleware.ResponseError(c, 2004, errors.New("无法匹配有效地址"))
		return
	}
	newFbList := []string{}
	for _, item := range detail.LoadBalance.GetForbidListByModel() {
		if public.InStringList(item, confIPList) && params.Addr != item {
			newFbList = append(newFbList, item)
		}
	}
	detail.LoadBalance.ForbidList = strings.Join(newFbList, ",")
	if err := detail.LoadBalance.Save(c, lib.GORMDefaultPool); err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}

	//todo 内存动态更新
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *ServiceController) ServiceEndpointClose(c *gin.Context) {
	params := &dto.ServiceEndpointCloseInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	search := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := search.ServiceDetail(c, lib.GORMDefaultPool, search)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	if detail.LoadBalance == nil {
		middleware.ResponseError(c, 2003, errors.New("负载节点为空"))
		return
	}
	confIPList := detail.LoadBalance.GetIPListByModel()
	if !public.InStringList(params.Addr, confIPList) {
		middleware.ResponseError(c, 2004, errors.New("无法匹配有效地址"))
		return
	}
	newFbList := []string{}
	for _, item := range detail.LoadBalance.GetForbidListByModel() {
		if public.InStringList(item, confIPList) && !public.InStringList(params.Addr, newFbList) {
			newFbList = append(newFbList, item)
		}
	}
	if !public.InStringList(params.Addr, newFbList) {
		newFbList = append(newFbList, params.Addr)
	}
	detail.LoadBalance.ForbidList = strings.Join(newFbList, ",")
	if err := detail.LoadBalance.Save(c, lib.GORMDefaultPool); err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}

	//todo 内存动态更新
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *ServiceController) ServiceAddHttp(c *gin.Context) {
	params := &dto.ServiceAddHttpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//验证service_name是否被占用
	search := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
	}
	if _, err := search.Find(c, lib.GORMDefaultPool, search); err == nil {
		middleware.ResponseError(c, 2001, errors.New("服务名被占用，请重新输入"))
		return
	}

	//验证rule前缀是否被占用
	ruleSearch := &dao.HttpRule{
		Rule: params.Rule,
	}
	if _, err := ruleSearch.Find(c, lib.GORMDefaultPool, ruleSearch); err == nil {
		middleware.ResponseError(c, 2001, errors.New("服务前缀或域名被占用，请重新输入"))
		return
	}

	//验证rule_type=0时以/开头，rule_type=1时不能出现/
	if params.RuleType == 0 {
		matched, _ := regexp.Match(`^/\S+$`, []byte(params.Rule))
		if !matched {
			middleware.ResponseError(c, 2002, errors.New("路径接入时必须以/开头"))
			return
		}
	}
	if params.RuleType == 1 {
		matched, _ := regexp.Match(`^[0-9a-z-_\.]+$`, []byte(params.Rule))
		if !matched {
			middleware.ResponseError(c, 2002, errors.New("域名接入时，只支持数字、小写字母、中划线、下划线"))
			return
		}
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		middleware.ResponseError(c, 2002, errors.New("ip列表与权重设置不匹配"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()
	info := &dao.ServiceInfo{
		LoadType:    public.LoadTypeHTTP,
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2002, err)
		return
	}

	loadBalance := &dao.LoadBalance{
		ServiceID:              info.ID,
		RoundType:              params.RoundType,
		IpList:                 params.IpList,
		WeightList:             params.WeightList,
		ForbidList:             params.ForbidList,
		UpstreamConnectTimeout: params.UpstreamConnectTimeout,
		UpstreamHeaderTimeout:  params.UpstreamHeaderTimeout,
		UpstreamIdleTimeout:    params.UpstreamIdleTimeout,
		UpstreamMaxIdle:        params.UpstreamMaxIdle,
		CheckMethod:            public.DefaultCheckMethod,
		CheckTimeout:           public.DefaultCheckTimeout,
		CheckInterval:          public.DefaultCheckInterval,
	}
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, err)
		return
	}

	httpRule := &dao.HttpRule{
		ServiceID:      info.ID,
		RuleType:       params.RuleType,
		Rule:           params.Rule,
		NeedHttps:      params.RuleType,
		NeedWebsocket:  params.RuleType,
		NeedStripUri:   params.NeedStripUri,
		UrlRewrite:     params.UrlRewrite,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err := httpRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         info.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		WhiteHostName:     params.WhiteHostName,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
	}
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *ServiceController) ServiceUpdateHttp(c *gin.Context) {
	params := &dto.ServiceUpdateHttpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//验证rule_type=0时以/开头，rule_type=1时不能出现/
	if params.RuleType == 0 {
		matched, _ := regexp.Match(`^/\S+$`, []byte(params.Rule))
		if !matched {
			middleware.ResponseError(c, 2002, errors.New("路径接入时必须以/开头"))
			return
		}
	}
	if params.RuleType == 1 {
		matched, _ := regexp.Match(`^[0-9a-z-_\.]+$`, []byte(params.Rule))
		if !matched {
			middleware.ResponseError(c, 2002, errors.New("域名接入时，只支持数字、小写字母、中划线、下划线"))
			return
		}
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		middleware.ResponseError(c, 2002, errors.New("ip列表与权重设置不匹配"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()

	service := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := service.ServiceDetail(c, lib.GORMDefaultPool, service)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	info := detail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2002, err)
		return
	}

	loadBalance := &dao.LoadBalance{}
	if detail.LoadBalance != nil {
		loadBalance = detail.LoadBalance
	}
	fmt.Println("params.UpstreamConnectTimeout", params.UpstreamConnectTimeout)
	loadBalance.ServiceID = info.ID
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	loadBalance.UpstreamConnectTimeout = params.UpstreamConnectTimeout
	loadBalance.UpstreamHeaderTimeout = params.UpstreamHeaderTimeout
	loadBalance.UpstreamIdleTimeout = params.UpstreamIdleTimeout
	loadBalance.UpstreamMaxIdle = params.UpstreamMaxIdle
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, err)
		return
	}

	httpRule := &dao.HttpRule{}
	if detail.HttpRule != nil {
		httpRule = detail.HttpRule
	}
	httpRule.ServiceID = info.ID
	httpRule.RuleType = params.RuleType
	httpRule.Rule = params.Rule
	httpRule.NeedHttps = params.NeedHttps
	httpRule.NeedWebsocket = params.NeedWebsocket
	httpRule.NeedStripUri = params.NeedStripUri
	httpRule.UrlRewrite = params.UrlRewrite
	httpRule.HeaderTransfor = params.HeaderTransfor
	httpRule.HeaderTransfor = params.HeaderTransfor
	httpRule.HeaderTransfor = params.HeaderTransfor
	httpRule.HeaderTransfor = params.HeaderTransfor
	httpRule.HeaderTransfor = params.HeaderTransfor
	httpRule.HeaderTransfor = params.HeaderTransfor

	if err := httpRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, err)
		return
	}

	accessControl := &dao.AccessControl{}
	if detail.AccessControl != nil {
		accessControl = detail.AccessControl
	}
	accessControl.ServiceID = info.ID
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *ServiceController) ServiceAddTcp(c *gin.Context) {
	params := &dto.ServiceAddTcpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//验证 service_name 是否被占用
	infoSearch := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		IsDelete:    0,
	}
	if _, err := infoSearch.Find(c, lib.GORMDefaultPool, infoSearch); err == nil {
		middleware.ResponseError(c, 2002, errors.New("服务名被占用，请重新输入"))
		return
	}

	//验证端口是否被占用?
	tcpRuleSearch := &dao.TcpRule{
		Port: params.Port,
	}
	if _, err := tcpRuleSearch.Find(c, lib.GORMDefaultPool, tcpRuleSearch); err == nil {
		middleware.ResponseError(c, 2003, errors.New("服务端口被占用，请重新输入"))
		return
	}
	grpcRuleSearch := &dao.GrpcRule{
		Port: params.Port,
	}
	if _, err := grpcRuleSearch.Find(c, lib.GORMDefaultPool, grpcRuleSearch); err == nil {
		middleware.ResponseError(c, 2004, errors.New("服务端口被占用，请重新输入"))
		return
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		middleware.ResponseError(c, 2005, errors.New("ip列表与权重设置不匹配"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()
	info := &dao.ServiceInfo{
		LoadType:    public.LoadTypeTCP,
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}

	loadBalance := &dao.LoadBalance{
		ServiceID:     info.ID,
		RoundType:     params.RoundType,
		IpList:        params.IpList,
		WeightList:    params.WeightList,
		ForbidList:    params.ForbidList,
		CheckMethod:   public.DefaultCheckMethod,
		CheckTimeout:  public.DefaultCheckTimeout,
		CheckInterval: public.DefaultCheckInterval,
	}
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}

	httpRule := &dao.TcpRule{
		ServiceID: info.ID,
		Port:      params.Port,
	}
	if err := httpRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2008, err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         info.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		WhiteHostName:     params.WhiteHostName,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
	}
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2009, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *ServiceController) ServiceUpdateTcp(c *gin.Context) {
	params := &dto.ServiceUpdateTcpInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		middleware.ResponseError(c, 2002, errors.New("ip列表与权重设置不匹配"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()

	service := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := service.ServiceDetail(c, lib.GORMDefaultPool, service)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	info := detail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, err)
		return
	}

	loadBalance := &dao.LoadBalance{}
	if detail.LoadBalance != nil {
		loadBalance = detail.LoadBalance
	}
	loadBalance.ServiceID = info.ID
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, err)
		return
	}

	tcpRule := &dao.TcpRule{}
	if detail.TcpRule != nil {
		tcpRule = detail.TcpRule
	}
	tcpRule.ServiceID = info.ID
	tcpRule.Port = params.Port
	if err := tcpRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}

	accessControl := &dao.AccessControl{}
	if detail.AccessControl != nil {
		accessControl = detail.AccessControl
	}
	accessControl.ServiceID = info.ID
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *ServiceController) ServiceAddGrpc(c *gin.Context) {
	params := &dto.ServiceAddGrpcInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//验证 service_name 是否被占用
	infoSearch := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		IsDelete:    0,
	}
	if _, err := infoSearch.Find(c, lib.GORMDefaultPool, infoSearch); err == nil {
		middleware.ResponseError(c, 2002, errors.New("服务名被占用，请重新输入"))
		return
	}

	//验证端口是否被占用?
	tcpRuleSearch := &dao.TcpRule{
		Port: params.Port,
	}
	if _, err := tcpRuleSearch.Find(c, lib.GORMDefaultPool, tcpRuleSearch); err == nil {
		middleware.ResponseError(c, 2003, errors.New("服务端口被占用，请重新输入"))
		return
	}
	grpcRuleSearch := &dao.GrpcRule{
		Port: params.Port,
	}
	if _, err := grpcRuleSearch.Find(c, lib.GORMDefaultPool, grpcRuleSearch); err == nil {
		middleware.ResponseError(c, 2003, errors.New("服务端口被占用，请重新输入"))
		return
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		middleware.ResponseError(c, 2002, errors.New("ip列表与权重设置不匹配"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()
	info := &dao.ServiceInfo{
		LoadType:    public.LoadTypeGRPC,
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, err)
		return
	}

	loadBalance := &dao.LoadBalance{
		ServiceID:     info.ID,
		RoundType:     params.RoundType,
		IpList:        params.IpList,
		WeightList:    params.WeightList,
		ForbidList:    params.ForbidList,
		CheckMethod:   public.DefaultCheckMethod,
		CheckTimeout:  public.DefaultCheckTimeout,
		CheckInterval: public.DefaultCheckInterval,
	}
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}

	grpcRule := &dao.GrpcRule{
		ServiceID:      info.ID,
		Port:           params.Port,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err := grpcRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         info.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		WhiteHostName:     params.WhiteHostName,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
	}
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2007, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
	return
}

func (admin *ServiceController) ServiceUpdateGrpc(c *gin.Context) {
	params := &dto.ServiceUpdateGrpcInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	//ip与权重数量一致
	if len(strings.Split(params.IpList, ",")) != len(strings.Split(params.WeightList, ",")) {
		middleware.ResponseError(c, 2002, errors.New("ip列表与权重设置不匹配"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()

	service := &dao.ServiceInfo{
		ID: params.ID,
	}
	detail, err := service.ServiceDetail(c, lib.GORMDefaultPool, service)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	info := detail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2003, err)
		return
	}

	loadBalance := &dao.LoadBalance{}
	if detail.LoadBalance != nil {
		loadBalance = detail.LoadBalance
	}
	loadBalance.ServiceID = info.ID
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err := loadBalance.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2004, err)
		return
	}

	grpcRule := &dao.GrpcRule{}
	if detail.GrpcRule != nil {
		grpcRule = detail.GrpcRule
	}
	grpcRule.ServiceID = info.ID
	//grpcRule.Port = params.Port
	grpcRule.HeaderTransfor = params.HeaderTransfor
	if err := grpcRule.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2005, err)
		return
	}

	accessControl := &dao.AccessControl{}
	if detail.AccessControl != nil {
		accessControl = detail.AccessControl
	}
	accessControl.ServiceID = info.ID
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(c, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(c, 2006, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
	return
}

package controller

import (
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/gin-gonic/gin"
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
	middleware.ResponseSuccess(c, map[string]interface{}{
		"serviceNum":      5,
		"todayRequestNum": 2000,
		"currentQps":      13,
		"appNum":          5,
	})
	return
}

func (admin *DashBoardController) FlowStat(c *gin.Context) {
	yesterdayStat := []int64{
		120,
		50,
		10,
		57,
		59,
		48,
		76,
		69,
		200,
		400,
		580,
		1500,
		2500,
		2300,
		1300,
		1700,
		1900,
		1000,
		800,
		570,
		500,
		360,
		200,
		105,
	}
	todayStat := []int64{
		78,
		23,
		78,
		123,
		325,
		378,
		456,
		478,
		500,
		800,
		760,
	}
	middleware.ResponseSuccess(c, map[string][]int64{
		"today":     todayStat,
		"yesterday": yesterdayStat,
	})
	return
}

func (admin *DashBoardController) ServiceStat(c *gin.Context) {
	middleware.ResponseSuccess(c, map[string]interface{}{
		"legend": []string{"HTTP", "TCP", "GRPC"},
		"data": []map[string]interface{}{
			{"value": 1, "name": "HTTP",},
			{"value": 2, "name": "TCP",},
			{"value": 3, "name": "GRPC",},
		}})
	return
}

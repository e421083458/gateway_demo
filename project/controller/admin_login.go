package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/gateway_demo/project/dao"
	"github.com/e421083458/gateway_demo/project/dto"
	"github.com/e421083458/gateway_demo/project/middleware"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

//AdminRegister admin_login路由注册
func AdminLoginRegister(router *gin.RouterGroup) {
	admin := AdminLogin{}
	router.POST("/login", admin.Login)
	router.POST("/logout", admin.Logout)
}

type AdminLogin struct {
}

func (AdminLogin *AdminLogin) Login(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.GetValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	adminInfo, err := (&dao.GatewayAdmin{}).LoginCheck(c, lib.GORMDefaultPool, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	adminSession := &dto.AdminSession{
		ID:        adminInfo.ID,
		LoginTime: time.Now(),
		UserName:  adminInfo.UserName,
	}
	session := sessions.Default(c)
	adminBts, _ := json.Marshal(adminSession)
	session.Set(public.AdminInfoSessionKey, string(adminBts))
	session.Save()
	fmt.Println(session.Get(public.AdminInfoSessionKey))
	middleware.ResponseSuccess(c, map[string]string{"token": adminInfo.UserName})
	return
}

func (AdminLogin *AdminLogin) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	middleware.ResponseSuccess(c, "")
	return
}

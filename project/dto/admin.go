package dto

import (
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"time"
)

type AdminChangePwdInput struct {
	Password    string `json:"password" form:"password" comment:"新密码" validate:"required"`
}

func (params *AdminChangePwdInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"用户名" en_comment:"username" validate:"required,is-validuser"`
	Passpord string `json:"password" form:"password" comment:"密码" en_comment:"password" validate:"required"`
}

func (params *AdminLoginInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type AdminSession struct {
	ID        int64     `json:"id" form:"id" comment:"用户ID" en_comment:"id" validate:""`
	UserName  string    `json:"user_name" form:"user_name" comment:"用户名" en_comment:"user_name" validate:""`
	LoginTime time.Time `json:"login_time" form:"login_time" comment:"登陆时间" en_comment:"login_time" validate:""`
}

type AdminInfoOutput struct {
	ID           int64     `json:"id" form:"id" comment:"用户ID" en_comment:"id" validate:""`
	Name         string    `json:"name" form:"name" comment:"用户名" en_comment:"name" validate:""`
	Avatar       string    `json:"avatar" form:"avatar" comment:"头像" en_comment:"avatar" validate:""`
	Introduction string    `json:"introduction" form:"introduction" comment:"介绍" en_comment:"introduction" validate:""`
	Roles        []string  `json:"roles" form:"roles" comment:"角色" en_comment:"roles" validate:""`
	LoginTime    time.Time `json:"login_time" form:"login_time" comment:"登陆时间" en_comment:"login_time" validate:""`
}

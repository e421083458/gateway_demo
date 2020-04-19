package dao

import (
	"fmt"
	"github.com/e421083458/gateway_demo/project/dto"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"time"
)

//GatewayAdmin 管理员
type GatewayAdmin struct {
	ID        int64     `json:"id" toml:"-" orm:"column(id);auto" description:"自增主键"`
	UserName  string    `json:"user_name" toml:"user_name" validate:"required"  orm:"column(user_name);" description:"用户名"`
	Salt      string    `json:"salt" toml:"salt" validate:"required"  orm:"column(salt);" description:"盐"`
	Password  string    `json:"password" toml:"password" validate:"required"  orm:"column(password);" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" orm:"column(update_at);size(100)" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" orm:"column(create_at);size(100)" description:"添加时间"`
	IsDelete  int       `json:"is_delete" orm:"column(is_delete);size(100)" description:"是否删除"`
}

//TableName db table
func (o *GatewayAdmin) TableName() string {
	return "gateway_admin"
}

//验证登陆
func (o *GatewayAdmin) LoginCheck(c *gin.Context, tx *gorm.DB, params *dto.AdminLoginInput) (*GatewayAdmin, error) {
	trace := public.GetGinTraceContext(c)
	adminInfo := &GatewayAdmin{}
	if err := tx.SetCtx(trace).Where("user_name=? and is_delete=0", params.UserName).Find(&adminInfo).Error; err == gorm.ErrRecordNotFound {
		return adminInfo, errors.New("用户不存在")
	}

	salt := adminInfo.Salt
	loginPassword := public.GenSaltPassword(params.Passpord, salt)
	fmt.Println(loginPassword)
	if loginPassword != adminInfo.Password {
		return adminInfo, errors.New("密码错误")
	}
	return adminInfo, nil
}

func (t *GatewayAdmin) Find(c *gin.Context, tx *gorm.DB, search *GatewayAdmin) (*GatewayAdmin, error) {
	model := &GatewayAdmin{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(model).Error
	return model, err
}

func (t *GatewayAdmin) Save(c *gin.Context, tx *gorm.DB) error {
	if err := tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error; err != nil {
		return err
	}
	return nil
}

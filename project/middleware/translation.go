package middleware

import (
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		//设置支持语言
		en := en.New()
		zh := zh.New()

		//设置国际化翻译器
		uni := ut.New(zh, zh, en)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			zh_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			val.RegisterValidation("is-validuser", func(fl validator.FieldLevel) bool {
				return fl.Field().String() == "admin"
			})
			val.RegisterValidation("valid_service_name", func(fl validator.FieldLevel) bool {
				matched, _ := regexp.Match(`[a-b0-9_]+`, []byte(fl.Field().String()))
				return matched
			})
			val.RegisterValidation("valid_rule", func(fl validator.FieldLevel) bool {
				matched, _ := regexp.Match(`\S+`, []byte(fl.Field().String()))
				return matched
			})
			val.RegisterValidation("valid_url_rewrite", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				for _, item := range strings.Split(fl.Field().String(), ",") {
					itemPart := strings.Split(item, " ")
					if len(itemPart) != 2 {
						return false
					}
				}
				return true
			})
			val.RegisterValidation("valid_header_transfor", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				for _, item := range strings.Split(fl.Field().String(), ",") {
					itemPart := strings.Split(item, " ")
					if len(itemPart) != 3 {
						return false
					}
					if !public.InStringList(itemPart[0], []string{"add", "del", "edit"}) {
						return false
					}
					matched1, _ := regexp.Match(`\S+`, []byte(itemPart[1])) //head_name
					matched2, _ := regexp.Match(`\S+`, []byte(itemPart[2])) //head_value
					if !matched1 || !matched2 {
						return false
					}
				}
				return true
			})
			val.RegisterValidation("valid_ip_list", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				for _, item := range strings.Split(fl.Field().String(), ",") {
					matched, _ := regexp.Match(`\S+`, []byte(item)) //ip_addr
					if !matched {
						return false
					}
				}
				return true
			})
			val.RegisterValidation("valid_ip_port_list", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				for _, item := range strings.Split(fl.Field().String(), ",") {
					matched, _ := regexp.Match(`\S+`, []byte(item)) //ip_port
					if !matched {
						return false
					}
					itemPart := strings.Split(item, ":")
					if len(itemPart) != 2 {
						return false
					}
					port, err := strconv.ParseInt(itemPart[1], 10, 64)
					if err != nil || port <= 0 {
						return false
					}
				}
				return true
			})
			val.RegisterValidation("valid_weight_list", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				for _, item := range strings.Split(fl.Field().String(), ",") {
					matched, _ := regexp.Match(`\d+`, []byte(item)) //weight
					if !matched {
						return false
					}
				}
				return true
			})

			//自定义验证器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("is-validuser", trans, func(ut ut.Translator) error {
				return ut.Add("is-validuser", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("is-validuser", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_service_name", trans, func(ut ut.Translator) error {
				return ut.Add("valid_service_name", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_service_name", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_rule", trans, func(ut ut.Translator) error {
				return ut.Add("valid_rule", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_rule", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_url_rewrite", trans, func(ut ut.Translator) error {
				return ut.Add("valid_url_rewrite", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_url_rewrite", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_header_transfor", trans, func(ut ut.Translator) error {
				return ut.Add("valid_header_transfor", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_header_transfor", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_ip_list", trans, func(ut ut.Translator) error {
				return ut.Add("valid_ip_list", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_ip_list", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_ip_port_list", trans, func(ut ut.Translator) error {
				return ut.Add("valid_ip_list", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_ip_list", fe.Field())
				return t
			})
			val.RegisterTranslation("valid_weight_list", trans, func(ut ut.Translator) error {
				return ut.Add("valid_weight_list", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_weight_list", fe.Field())
				return t
			})
			break
		}

		c.Set(public.TranslatorKey, trans)
		c.Set(public.ValidatorKey, val)
		c.Next()
	}
}
package dao

import (
	"github.com/e421083458/gateway_demo/project/dto"
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"sync"
	"time"
)

type App struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	AppID     string    `json:"app_id" gorm:"column:app_id" description:"租户id	"`
	Name      string    `json:"name" gorm:"column:name" description:"租户名称	"`
	Secret    string    `json:"secret" gorm:"column:secret" description:"密钥"`
	WhiteIPS  string    `json:"white_ips" gorm:"column:white_ips" description:"ip白名单，支持前缀匹配"`
	Qpd       int64     `json:"qpd" gorm:"column:qpd" description:"日请求量限制"`
	Qps       int64     `json:"qps" gorm:"column:qps" description:"每秒请求量限制"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间	"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	IsDelete  int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

func (t *App) TableName() string {
	return "gateway_app"
}

func (t *App) Find(c *gin.Context, tx *gorm.DB, search *App) (*App, error) {
	model := &App{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(model).Error
	return model, err
}

func (t *App) Save(c *gin.Context, tx *gorm.DB) error {
	if err := tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error; err != nil {
		return err
	}
	return nil
}

func (t *App) APPList(c *gin.Context, tx *gorm.DB, params *dto.APPListInput) ([]App, int64, error) {
	var list []App
	var count int64
	pageNo := params.PageNo
	pageSize := params.PageSize

	//limit offset,pagesize
	offset := (pageNo - 1) * pageSize
	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName()).Select("*")
	query = query.Where("is_delete=?", 0)
	if params.Info != "" {
		query = query.Where(" (name like ? or app_id like ?)", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Count(&count).Error
	if errCount != nil {
		return nil, 0, err
	}
	return list, count, nil
}

var AppHandler *AppManager

type AppManager struct {
	appMap     map[string]*App
	appSlice   []*App
	appMapLock sync.RWMutex
	init       sync.Once
	err        error
}

func init() {
	AppHandler = NewAppManager()
}

func NewAppManager() *AppManager {
	return &AppManager{
		appMap:     make(map[string]*App),
		appSlice:   []*App{},
		appMapLock: sync.RWMutex{},
	}
}

func (s *AppManager) GetAppList() []*App {
	return s.appSlice
}

func (s *AppManager) LoadOnce() error {
	s.init.Do(func() {
		model := App{}
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Set("trace", lib.NewTrace())
		params := &dto.APPListInput{PageSize: 99999, PageNo: 1}
		list, _, err := model.APPList(c, lib.GORMDefaultPool, params)
		if err != nil {
			s.err = err
			return
		}
		s.appMapLock.Lock()
		defer s.appMapLock.Unlock()
		for _, item := range list {
			tmp:=item
			s.appMap[item.AppID] = &tmp
			s.appSlice = append(s.appSlice, &tmp)
		}
		return
	})
	return s.err
}

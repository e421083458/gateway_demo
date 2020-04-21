package dto

import (
	"github.com/e421083458/gateway_demo/project/public"
	"github.com/gin-gonic/gin"
	"time"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"查找信息" validate:""`
	PageSize int    `json:"page_size" form:"page_size" comment:"页数" validate:"required,min=1,max=999"`
	PageNo   int    `json:"page_no" form:"page_no" comment:"页码" validate:"required,min=1,max=999"`
}

type ServiceLoadTypeStat struct {
	LoadType int `json:"load_type" form:"load_type" comment:"负载类型"`
	Num      int `json:"num" form:"num" comment:"服务数量"`
}

type ServiceStatOutput struct {
	Legend []string                `json:"legend" form:"legend" comment:"负载数量"`
	Data   []ServiceStatItemOutput `json:"data" form:"data" comment:"统计详情"`
}

type ServiceStatItemOutput struct {
	Value int    `json:"value" form:"value" comment:"负载数量"`
	Name  string `json:"name" form:"name" comment:"负载类型"`
}

func (params *ServiceListInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceListItemOutput struct {
	ID          int64     `json:"id" gorm:"primary_key"`
	LoadType    int       `json:"load_type" gorm:"column:load_type" description:"账号创建时间"`
	ServiceAddr string    `json:"service_addr" gorm:"column:service_addr" description:"服务地址"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	UpdatedAt   time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间	"`
	CreatedAt   time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	QPS         int64     `json:"qps" description:"每秒请求数"`
	QPD         int64     `json:"qpd" description:"每天请求数"`
	TotalNode   int       `json:"total_node" description:"总节点数"`
}

type ServiceDetailInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" validate:"required"`
}

func (params *ServiceDetailInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceEndpointCloseInput struct {
	ID   int64  `json:"id" form:"id" comment:"服务ID" validate:"required"`
	Addr string `json:"addr" form:"addr" comment:"服务地址" validate:"required"`
}

func (params *ServiceEndpointCloseInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceAddHttpInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	RuleType          int    `json:"rule_type" form:"rule_type" comment:"接入方式" validate:""`
	Rule              string `json:"rule" form:"rule" comment:"域名或者前缀" validate:"required,valid_rule"`
	NeedHttps         int    `json:"need_https" form:"need_https" comment:"支持https" validate:""`
	NeedStripUri      int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启动strip_uri" validate:""`
	NeedWebsocket     int    `json:"need_websocket" form:"need_websocket" comment:"支持websocket" validate:""`
	UrlRewrite        string `json:"url_rewrite" form:"url_rewrite" comment:"url重写" validate:"valid_url_rewrite"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"header头转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ClientIPFlowLimit int64  `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int64  `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ip_port_list"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weight_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`

	UpstreamConnectTimeout int `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时，单位s，0表示无限制" validate:""`
	UpstreamHeaderTimeout  int `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时，单位s，0表示无限制" validate:""`
	UpstreamIdleTimeout    int `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间，单位s，0表示无限制" validate:""`
	UpstreamMaxIdle        int `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数，0表示无限制" validate:""`
}

func (params *ServiceAddHttpInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceUpdateHttpInput struct {
	ID                int64  `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	RuleType          int    `json:"rule_type" form:"rule_type" comment:"接入方式" validate:""`
	Rule              string `json:"rule" form:"rule" comment:"域名或者前缀" validate:"required,valid_rule"`
	NeedHttps         int    `json:"need_https" form:"need_https" comment:"支持https" validate:""`
	NeedStripUri      int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启动strip_uri" validate:""`
	NeedWebsocket     int    `json:"need_websocket" form:"need_websocket" comment:"支持websocket" validate:""`
	UrlRewrite        string `json:"url_rewrite" form:"url_rewrite" comment:"url重写" validate:"valid_url_rewrite"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"header头转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ClientIPFlowLimit int64  `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int64  `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ip_port_list"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weight_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`

	UpstreamConnectTimeout int `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"建立连接超时，单位s，0表示无限制" validate:""`
	UpstreamHeaderTimeout  int `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"获取header超时，单位s，0表示无限制" validate:""`
	UpstreamIdleTimeout    int `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"链接最大空闲时间，单位s，0表示无限制" validate:""`
	UpstreamMaxIdle        int `json:"upstream_max_idle" form:"upstream_max_idle" comment:"最大空闲链接数，0表示无限制" validate:""`
}

func (params *ServiceUpdateHttpInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceAddGrpcInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ClientIPFlowLimit int64  `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int64  `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ip_port_list"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weight_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`
}

func (params *ServiceAddGrpcInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceUpdateGrpcInput struct {
	ID                int64  `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ClientIPFlowLimit int64  `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int64  `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ip_port_list"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weight_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`
}

func (params *ServiceUpdateGrpcInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceAddTcpInput struct {
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	HeaderTransfor    string `json:"header_transfor" form:"header_transfor" comment:"header头转换" validate:"valid_header_transfor"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ClientIPFlowLimit int64  `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int64  `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ip_port_list"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weight_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`
}

func (params *ServiceAddTcpInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

type ServiceUpdateTcpInput struct {
	ID                int64  `json:"id" form:"id" comment:"服务ID" validate:"required"`
	ServiceName       string `json:"service_name" form:"service_name" comment:"服务名称" validate:"required,valid_service_name"`
	ServiceDesc       string `json:"service_desc" form:"service_desc" comment:"服务描述" validate:"required"`
	Port              int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" validate:"required,min=8001,max=8999"`
	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限验证" validate:""`
	BlackList         string `json:"black_list" form:"black_list" comment:"黑名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"白名单IP，以逗号间隔，白名单优先级高于黑名单" validate:"valid_ip_list"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ClientIPFlowLimit int64  `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端IP限流" validate:""`
	ServiceFlowLimit  int64  `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" validate:""`
	RoundType         int    `json:"round_type" form:"round_type" comment:"轮询策略" validate:""`
	IpList            string `json:"ip_list" form:"ip_list" comment:"IP列表" validate:"required,valid_ip_port_list"`
	WeightList        string `json:"weight_list" form:"weight_list" comment:"权重列表" validate:"required,valid_weight_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`
}

func (params *ServiceUpdateTcpInput) GetValidParams(c *gin.Context) error {
	return public.DefaultGetValidParams(c, params)
}

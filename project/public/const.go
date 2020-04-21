package public

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"

	//flow_count_key
	RedisFlowCountDayKey  = "flow_count_day"
	RedisFlowCountHourKey = "flow_count_hour"

	//flow_limit_prefix
	FlowServicePrefix = "service_"
	FlowAPPPrefix     = "app_"
	FlowTotal         = "total"

	//sessionKey
	AdminInfoSessionKey = "admin_info"

	//load_type
	LoadTypeHTTP = 0
	LoadTypeTCP  = 1
	LoadTypeGRPC = 2

	//default check setting
	DefaultCheckMethod    = 0
	DefaultCheckTimeout   = 2
	DefaultCheckMaxErrNum = 2
	DefaultCheckInterval  = 5
)

var LoadTypeMap = map[int]string{
	LoadTypeHTTP: "HTTP",
	LoadTypeTCP:  "TCP",
	LoadTypeGRPC: "GRPC",
}

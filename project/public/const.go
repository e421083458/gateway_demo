package public

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"

	//flow_count_key
	RedisFlowCountDayKey  = "flow_count_day"
	RedisFlowCountHourKey = "flow_count_hour"

	//flow_count_prefix
	FlowCountServicePrefix = "service_"
	FlowCountAPPPrefix     = "app_"

	//sessionKey
	AdminInfoSessionKey = "admin_info"

	//load_type
	LoadTypeHTTP = 0
	LoadTypeTCP  = 1
	LoadTypeGRPC = 2

	//default check setting
	DefaultCheckMethod   = 0
	DefaultCheckTimeout  = 2
	DefaultCheckInterval = 5
)

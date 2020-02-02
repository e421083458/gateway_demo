package load_balance

type LoadBalance interface {
	Add(...string) error
	Get(string) (string, error)

	//后期服务发现补充
	Update()
}

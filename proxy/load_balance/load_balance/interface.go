package load_balance

type LoadBalance interface {
	Add(...string) error
	Get(string) (string, error)
}
package common

type LoadBalance interface {
	Add(...string) error
	Get(string) (string, error)
}
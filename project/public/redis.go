package public

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/garyburd/redigo/redis"
)

func RedisConfPipline(name string, pip ...func(c redis.Conn)) error {
	c, err := lib.RedisConnFactory(name)
	if err != nil {
		return err
	}
	defer c.Close()
	for _, f := range pip {
		f(c)
	}
	c.Flush()
	return nil
}
package public

import (
	"github.com/garyburd/redigo/redis"
)

func RedisConfPipline(pip ...func(c redis.Conn)) error {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
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

func RedisConfDo(commandName string, args ...interface{}) (interface{}, error) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return nil,err
	}
	defer c.Close()
	return c.Do(commandName, args...)
}

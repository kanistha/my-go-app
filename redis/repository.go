package redis

import (
	"mygoapp"
	"github.com/gomodule/redigo/redis"
)

type redisRepository struct {

}

func NewRepository() *redisRepository {
	return &redisRepository{ }
}

func (repo *redisRepository) Get(id string) (mygoapp.User, error){
	conn := newRedis().Get()
	defer conn.Close()

	s, err := conn.Do("GET", id)
	if err != nil {
		return  mygoapp.User{},err
	}
	user := s.(mygoapp.User)
	return user, err
}



func newRedis() redis.Pool{
	pool := redis.Pool{
		MaxIdle:80,
		MaxActive:12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
	return pool
}


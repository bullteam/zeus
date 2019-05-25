package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
	"time"
)

var (
	// DefaultKey the collection name of redis
	DefaultKey = "zeus_admin"
	Cache      *Redis
)

type Redis struct {
	pool     *redis.Pool
	connInfo string
	dbNum    int
	key      string
	password string
	maxIdle  int
}

// NewRedisPool create new redis pool with default collection name.
func NewRedis() *Redis {
	return &Redis{}
}

// actually do the redis cmds, args[0] must be the key name.
func (r *Redis) Do(cmdName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}
	args[0] = r.associate(args[0])
	c := r.pool.Get()
	defer c.Close()

	return c.Do(cmdName, args...)
}

// associate with config key.
func (r *Redis) associate(originKey interface{}) string {
	return fmt.Sprintf("%s:%s", r.key, originKey)
}

// Get from redis.
func (r *Redis) Get(key string) interface{} {
	if v, err := r.Do("GET", key); err == nil {
		return v
	}

	return nil
}

// GetMulti get multi from redis.
func (r *Redis) GetMulti(keys []string) []interface{} {
	c := r.pool.Get()
	defer c.Close()
	var args []interface{}
	for _, key := range keys {
		args = append(args, r.associate(key))
	}
	values, err := redis.Values(c.Do("MGET", args...))
	if err != nil {
		return nil
	}

	return values
}

// Put set to redis.
func (r *Redis) Put(key string, val interface{}, timeout time.Duration) error {
	_, err := r.Do("SETEX", key, int64(timeout/time.Second), val)

	return err
}

// Delete delete in redis.
func (r *Redis) Delete(key string) error {
	_, err := r.Do("DEL", key)

	return err
}

// IsExist check key existence in redis.
func (r *Redis) IsExist(key string) bool {
	v, err := redis.Bool(r.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return v
}

// Incr increase counter in redis.
func (r *Redis) Incr(key string) error {
	_, err := redis.Bool(r.Do("INCRBY", key, 1))

	return err
}

// Decr decrease counter in redis.
func (r *Redis) Decr(key string) error {
	_, err := redis.Bool(r.Do("INCRBY", key, -1))

	return err
}

// ClearAll clean all key in redis. delete this redis collection.
func (r *Redis) ClearAll() error {
	c := r.pool.Get()
	defer c.Close()
	cachedKeys, err := redis.Strings(c.Do("KEYS", r.key+":*"))
	if err != nil {
		return err
	}
	for _, str := range cachedKeys {
		if _, err = c.Do("DEL", str); err != nil {
			return err
		}
	}
	return err
}

// connect to redis.
func (r *Redis) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", r.connInfo)
		if err != nil {
			return nil, err
		}

		if r.password != "" {
			if _, err := c.Do("AUTH", r.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", r.dbNum)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	r.pool = &redis.Pool{
		MaxIdle:     r.maxIdle,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

func NewRedisPool(config string) (*Redis, error) {
	r := NewRedis()

	var cf map[string]string
	_ = json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}
	if _, ok := cf["conn"]; !ok {
		return nil, errors.New("config has no conn key")
	}

	// Format redis://<password>@<host>:<port>
	cf["conn"] = strings.Replace(cf["conn"], "redis://", "", 1)
	if i := strings.Index(cf["conn"], "@"); i > -1 {
		cf["password"] = cf["conn"][0:i]
		cf["conn"] = cf["conn"][i+1:]
	}

	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	if _, ok := cf["maxIdle"]; !ok {
		cf["maxIdle"] = "3"
	}
	r.key = cf["key"]
	r.connInfo = cf["conn"]
	r.dbNum, _ = strconv.Atoi(cf["dbNum"])
	r.password = cf["password"]
	r.maxIdle, _ = strconv.Atoi(cf["maxIdle"])

	r.connectInit()

	c := r.pool.Get()
	defer c.Close()

	return r, c.Err()
}

// init redis pool
func RedisInit() {
	redisConn := beego.AppConfig.String("redis_conn")
	redisPort := beego.AppConfig.String("redis_port")
	redisWwd := beego.AppConfig.String("redis_pwd")
	beego.Info(redisConn)
	redisConf := fmt.Sprintf(`{"key":"%s","conn":"%s:%s","dbNum":"%d","password":"%s"}`,
		"zeus_admin",
		redisConn,
		redisPort,
		0,
		redisWwd,
	)
	var err error
	Cache, err = NewRedisPool(redisConf)
	if err != nil {
		beego.Info(redisConf)
		beego.Error("Redis connection fail:" + err.Error())
	}
}

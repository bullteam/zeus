package tests

import (
	"fmt"
	"testing"
	"zeus/pkg/components"
	"zeus/pkg/utils"
)

var Cache *components.Redis

func Test_GetUserPerms(t *testing.T) {

}

func Test_Redis(t *testing.T) {
	redisConf := fmt.Sprintf(`{"key":"%s","conn":"%s:%s","dbNum":"%d"}`,
		"zeus_admin",
		"hyman_redis",
		"6379",
		0,
	)
	var err error
	Cache, err = components.NewRedisPool(redisConf)
	if err != nil {
		t.Log(redisConf)
		t.Log("Redis connection fail:" + err.Error())
	}

	data := Cache.Get("test11")
	fmt.Println(utils.ByteToString(data.([]byte)))
}

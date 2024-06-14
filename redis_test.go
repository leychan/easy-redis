package easyredis

import (
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

type a struct {
	Name string
}

var easyrdb = New(redis.Options{
	Addr:     "127.0.0.1:6379",
	Password: "",
	DB:       0,
})

func TestEasyRedis_GetStruct(t *testing.T) {
	var b a
	fmt.Println(easyrdb.GetStruct("test", &b))
	fmt.Println(b)
}

func TestEasyRedis_SetStruct(t *testing.T) {
	b := a{
		Name: "test",
	}

	easyrdb.SetStruct("test", b, time.Second * 300)
}

func TestEasyRedis_Set(t *testing.T) {
	
	easyrdb.SetEasy("test_set1", "test", time.Second*300)
	easyrdb.SetEasy("test_set2", 2, time.Second*300)
	easyrdb.SetEasy("test_set3", a{Name: "chenlei"}, time.Second*300)
}

func TestEasyRedis_Get(t *testing.T) {
	easyrdb.SetEasy("test_set1", "test", time.Second*300)
	easyrdb.SetEasy("test_set2", 2, time.Second*300)
	easyrdb.SetEasy("test_set3", true, time.Second*300)
	// easy.Set("test_set3", a{Name: "chenlei"}, time.Second*300)
	fmt.Println(easyrdb.GetEasy("test_set1",))
	fmt.Println(easyrdb.GetEasy("test_set2"))
	fmt.Println(easyrdb.GetBool("test_set3"))
	str, exist, _ := easyrdb.GetEasy("test_set3")
	if !exist {
		fmt.Println("不存在")
		return
	}
	fmt.Println(str)

}


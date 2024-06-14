package easyredis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type EasyRedis redis.Client

func New(rOpt redis.Options) *EasyRedis {
    client := redis.NewClient(&rOpt)
    return (*EasyRedis)(client)
}

// Lock 设置锁
// return true: 设置锁成功 false: 设置锁失败
func (r *EasyRedis) Lock(key string, expire time.Duration) bool {
    err := r.SetNX(context.Background(), key, 1, expire).Err()
    return err == nil
}

// Unlock 解锁
// return true: 解锁成功 false: 解锁失败
func (r *EasyRedis) Unlock(key string) bool {
    err := r.Del(context.Background(), key).Err()
    return err == nil
}

// SetStruct 保存结构体
func (r *EasyRedis) SetStruct(key string, value any, expire time.Duration) error {
    jsonBytes, err := json.Marshal(value)
    if err != nil {
        return err
    }
    err = r.Set(context.Background(), key, string(jsonBytes), expire).Err()
    return err
}

// SetEasy 保存
func (r *EasyRedis) SetEasy(key string, value any, expire time.Duration) error {
    return r.Set(context.Background(), key, value, expire).Err()
}

// GetStruct 获取结构体 
// return (bool, error) 
// bool true: 有值 false: 无值
func (r *EasyRedis) GetStruct(key string, value any) (bool, error) {
    val, err := r.Get(context.Background(), key).Result()
    if err == redis.Nil {
        return false, nil
    }
    if err != nil {
        return false, err
    }
    err = json.Unmarshal([]byte(val), value)
    return true, err
}

// Get 获取
// return (string, bool, error)
// bool true: 有值 false: 无值
func (r *EasyRedis) GetEasy(key string) (any, bool, error) {
    val, err := r.Get(context.Background(), key).Result()
    if err == redis.Nil {
        return nil, false, nil
    }
    if err != nil {
        return nil, false, err
    }
    return val, true, nil
}

// GetInt 获取int
// return (int64, bool, error) 
// bool true: 有值 false: 无值
func (r *EasyRedis) GetInt64(key string) (int64, bool, error) {
    val, err := r.Get(context.Background(), key).Int64()
    if err == redis.Nil {
        return 0, false, nil
    }
    if err != nil {
        return 0, false, err
    }
    return val, true, nil
}

// GetBool 获取bool
// return (bool, bool, error)
func (r *EasyRedis) GetBool(key string) (bool, bool, error) {
    val, err := r.Get(context.Background(), key).Bool()
    if err == redis.Nil {
        return false, false, nil
    }

    if err != nil {
        return false, false, err
    }
    return val, true, nil
}

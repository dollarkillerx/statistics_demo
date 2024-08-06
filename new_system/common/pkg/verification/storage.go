package verification

import (
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	key = "captcha:"
)

type RedisStore struct {
	conn *redis.Client
}

func (r *RedisStore) Set(id string, value string) error {
	// 验证码5分钟
	if err := r.conn.SetEx(context.TODO(), key+id, value, time.Second*60*5).Err(); err != nil {
		log.Error().Msgf("store captcha to redis error: %v", err)
		return err
	}

	return nil
}

func (r *RedisStore) Get(id string, clear bool) string {
	k := key + id
	v, err := r.conn.Get(context.TODO(), k).Result()

	if err != nil {
		return ""
	}

	if clear {
		r.conn.Del(context.TODO(), k)
	}

	return v
}

func (r *RedisStore) Verify(id, answer string, clear bool) bool {
	val := r.Get(id, clear)
	return val == strings.TrimSpace(answer)
}

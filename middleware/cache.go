package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type bodyCacheWriter struct {
	gin.ResponseWriter
	cache      *cache.Cache
	requestURI string
}

func (w bodyCacheWriter) Writer(b []byte) (int, error) {
	status := w.Status()
	if 200 <= status && status <= 299 {
		w.cache.Set(&cache.Item{
			Ctx:   &gin.Context{},
			Key:   w.requestURI,
			Value: b,
			TTL:   time.Hour,
		})
	}

	return w.ResponseWriter.Write(b)
}
func CacheCheck() gin.HandlerFunc {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server": os.Getenv("REDIS_ADDR_PORT"),
		},
	})

	myCache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return func(ctx *gin.Context) {
		var response interface{}
		fmt.Println(ctx.Request.RequestURI)
		err := myCache.Get(&gin.Context{}, ctx.Request.RequestURI, &response)
		if err != nil {
			bcw := &bodyCacheWriter{cache: myCache, requestURI: ctx.Request.RequestURI, ResponseWriter: ctx.Writer}
			ctx.Writer = bcw
			ctx.Next()
		} else {
			res := helper.ResponseFormatter(http.StatusOK, "success", "", response.([]byte))
			ctx.JSON(http.StatusOK, res)
			ctx.Abort()
		}
	}
}

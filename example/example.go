package main

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	njgw "github.com/mich4s/nestjsgolangworker"
)

func main() {
	service := njgw.NewService(&redis.Options{
		Addr: "localhost:6379",
	})
	service.MessageHandler("hello", func(ctx *njgw.Context) {
		fmt.Println(ctx.GetData())
		ctx.Send("Hello World!")
	})
	service.MessageHandler("toFail", func(ctx *njgw.Context) {
		fmt.Println(ctx.GetData())
		ctx.SendError("test error")
	})
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

#nestjsgolangworker
#####This project allow to easily connect Golang worker with NestJS monolith application by Reddis microservice

I've made this package because i needed fast worker for some of compute-heavy tasks for my monolith app. Becase configuring redis every time is waste of time
this library creates small abstraction for redis library and allow code to be executed "more reactive".

At this moment library allows only this kind of configuration
`this.client.send<string>({ cmd: 'hello' }, { test: 'test' })`

First big TODO is for sure creating universal standard for channel declaration and some code refactor. At this moment this works fine for me so I've posted this here so maybe someone will find this useful.



##Example
```package main
   
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
   		fmt.Println(ctx)
   		// ctx.SendError("test error")
   		ctx.Send("Hello World!")
   	})
   	wg := sync.WaitGroup{}
   	wg.Add(1)
   	wg.Wait()
   }
```

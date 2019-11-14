package nestjsredis

import "github.com/go-redis/redis"

type Handler = func(ctx *Context)

type ResponseWriter = func(response string)

type Service struct {
	client       *redis.Client
	subscription map[string]Handler
}

type Context struct {
	id             string
	message        *message
	responseWriter ResponseWriter
}

type channel struct {
	income  string
	outcome string
}

type message struct {
	Data    interface{}       `json:"data"`
	Id      string            `json:"id"`
	Pattern map[string]string `json:"pattern"`
}

type responseError struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type response struct {
	Err        *responseError `json:"err"`
	Id         string         `json:"id"`
	IsDisposed bool           `json:"isDisposed"`
	Response   interface{}    `json:"response"`
}

package nestjsredis

type Handler = func(ctx *Context)

type ResponseWriter = func(response string)

type subscription struct {
	channel string
	handler Handler
}

type Pattern struct {
	Cmd string `json:"cmd"`
}

type Channel struct {
	income  string
	outcome string
}

type Message struct {
	Data    interface{}       `json:"data"`
	Id      string            `json:"id"`
	Pattern map[string]string `json:"pattern"`
}

type ResponseError struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Response struct {
	Err        *ResponseError `json:"err"`
	Id         string         `json:"id"`
	IsDisposed bool           `json:"isDisposed"`
	Response   interface{}    `json:"response"`
}

type Context struct {
	id      string
	message *Message
	//response *Response
	//responseError  *ResponseError
	responseWriter ResponseWriter
}

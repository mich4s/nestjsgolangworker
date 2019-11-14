package nestjsredis

import (
	"encoding/json"
	"fmt"
)

func (ctx *Context) Send(response interface{}) {
	ctx.send(&Response{
		Err:        nil,
		Id:         ctx.id,
		IsDisposed: true,
		Response:   response,
	})
}

func (ctx *Context) SendError(message string) {
	ctx.send(&Response{
		Err: &ResponseError{
			Message: message,
			Status:  "error",
		},
		Id:         ctx.id,
		IsDisposed: true,
		Response:   nil,
	})
}

func (ctx *Context) send(response *Response) {
	result, err := json.Marshal(response)
	fmt.Println(string(result))
	if err != nil {
		fmt.Println(err)
	} else {
		ctx.responseWriter(string(result))
	}

}

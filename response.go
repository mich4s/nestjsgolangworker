package nestjsredis

import (
	"encoding/json"
	"fmt"
)

func (ctx *Context) Send(responseData interface{}) {
	ctx.send(&response{
		Err:        nil,
		Id:         ctx.id,
		IsDisposed: true,
		Response:   responseData,
	})
}

func (ctx *Context) SendError(message string) {
	ctx.send(&response{
		Err: &responseError{
			Message: message,
			Status:  "error",
		},
		Id:         ctx.id,
		IsDisposed: true,
		Response:   nil,
	})
}

func (ctx *Context) send(response *response) {
	result, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	} else {
		ctx.responseWriter(string(result))
	}

}

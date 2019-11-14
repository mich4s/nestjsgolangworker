package nestjsredis

func (ctx *Context) GetData() interface{} {
	return ctx.message.Data
}

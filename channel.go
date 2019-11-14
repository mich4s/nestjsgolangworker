package nestjsredis

import "fmt"

func createChannel(command string) *channel {
	return &channel{
		income:  fmt.Sprintf("{\"cmd\":\"%s\"}_ack", command),
		outcome: fmt.Sprintf("{\"cmd\":\"%s\"}_res", command),
	}
}

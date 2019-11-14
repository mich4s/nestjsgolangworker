package nestjsredis

import "fmt"

func createChannel(command string) *Channel {
	return &Channel{
		income:  fmt.Sprintf("{\"cmd\":\"%s\"}_ack", command),
		outcome: fmt.Sprintf("{\"cmd\":\"%s\"}_res", command),
	}
}

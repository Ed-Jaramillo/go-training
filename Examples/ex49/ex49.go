package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel int
	isUrgent bool
}

type groupMessage struct {
	groupName string
	messageContent string
	priorityLevel int
}

type systemAlert struct {
	alertCode string
	messageContent string
}

func (d directMessage) importance() int{
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

func (s systemAlert) importance() int {
	return 100
}

func main() {
	notifications := []notification{
		directMessage{senderUsername: "alice", messageContent: "Hi!", priorityLevel: 10, isUrgent: true},
		groupMessage{groupName: "devs", messageContent: "Standup at 9", priorityLevel: 20},
		systemAlert{alertCode: "E500", messageContent: "Server down"},
	}

	for _, n := range notifications {
		fmt.Printf("%T - importance: %d\n", n, n.importance())
	}
}


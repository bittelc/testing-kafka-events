package main

import (
	"github.com/satori/go.uuid"
)

type Event struct {
	AccId string
	Type  string
}

type CreateEvent struct {
	Event
	AccName string
}

func NewCreateAccountEvent(name string) CreateEvent {
	event := new(CreateEvent)
	event.Type = "CreateEvent"
	event.AccId = uuid.NewV4().String()
	event.AccName = name
	return *event
}

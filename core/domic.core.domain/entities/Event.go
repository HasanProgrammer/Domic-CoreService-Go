package entities

import "time"

type Event struct {
	Id          string
	AggregateId string
	Type        string // Name of Event
	Service     string // Name Of Service
	Payload     string
	Table       string
	Action      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	IsActive    bool
	IsDeleted   bool
}

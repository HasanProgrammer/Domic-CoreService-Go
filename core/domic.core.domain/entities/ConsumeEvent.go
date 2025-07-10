package entities

import "time"

// for command of [ CQRS ]
type ConsumeEvent struct {
	Id           string
	Type         string
	CountOfRetry int32
	CreatedAt    time.Time
}

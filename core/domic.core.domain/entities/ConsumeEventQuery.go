package entities

import "time"

// for query of [ CQRS ]
type ConsumeEventQuery struct {
	Id           string
	Type         string
	CountOfRetry int32
	CreatedAt    time.Time
}

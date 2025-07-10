package entities

import "time"

type SystemException struct {
	Id        string
	Service   string
	Action    string
	Message   string
	Exception string
	CreatedAt time.Time
}

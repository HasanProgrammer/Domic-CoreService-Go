package entities

import "time"

type SystemRequest struct {
	Id        string
	IpClient  string
	Service   string
	Action    string
	Header    string
	Payload   string
	CreatedAt time.Time
}

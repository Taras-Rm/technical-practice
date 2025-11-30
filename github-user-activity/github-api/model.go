package githubapi

type EventType string

const (
	PushEventType   EventType = "PushEvent"
	CreateEventType EventType = "CreateEvent"
)

type PublicEvent struct {
	Id    string    `json:"id"`
	Type  EventType `json:"type"`
	Actor struct {
		Id    int64  `json:"id"`
		Login string `json:"login"`
	} `json:"actor"`
	Repo struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"repo"`
}

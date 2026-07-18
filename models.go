package audited

import "time"

// Event An event represents an action perform by an actor towards a target or targets at a given time.
type Event struct {
	Action     string                  `json:"action"`
	Actor      Actor                   `json:"actor"`
	Context    Context                 `json:"context"`
	SourceID   string                  `json:"source_id"`
	Targets    []Target                `json:"targets"`
	Version    int                     `json:"version"`
	Metadata   *map[string]interface{} `json:"metadata,omitempty"`
	OccurredAt time.Time               `json:"occurred_at"`
}

// Actor An actor represents the person or entity that performed an action.
type Actor struct {
	Id       string                  `json:"id"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	Name     *string                 `json:"name,omitempty"`
	Type     string                  `json:"type"`
}

// Context The context holds details such as the IP and user agent of the person or entity that performed an action.
type Context struct {
	Location  string  `json:"location"`
	UserAgent *string `json:"user_agent,omitempty"`
}

// Target It represents an event target.
type Target struct {
	ID       string                  `json:"id"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	Name     *string                 `json:"name,omitempty"`
	Type     string                  `json:"type"`
}

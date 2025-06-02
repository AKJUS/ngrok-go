package ngrok

import "time"

// EventType represents the type of event that occurred
type EventType int

const (
	EventTypeAgentConnectSucceeded EventType = iota
	EventTypeAgentDisconnected
	EventTypeAgentHeartbeatReceived
)

func (t EventType) String() string {
	return [...]string{
		"AgentConnectSucceeded",
		"AgentDisconnected",
		"AgentHeartbeatReceived",
	}[t]
}

// Event is the interface implemented by all event types
type Event interface {
	EventType() EventType
	Timestamp() time.Time
}

// baseEvent provides common functionality for all events
type baseEvent struct {
	Type       EventType
	OccurredAt time.Time
}

func (e baseEvent) EventType() EventType { return e.Type }
func (e baseEvent) Timestamp() time.Time { return e.OccurredAt }

// EventHandler is the function type for event callbacks. EventHandlers must not
// block. If you would like to run operations on an event that will block or
// fail, instead write your handler to either non-blockingly push the Event onto
// a channel or spin up a goroutine.
type EventHandler func(Event)

// EventAgentConnectSucceeded is emitted when an agent successfully establishes a connection
type EventAgentConnectSucceeded struct {
	baseEvent
	Agent   Agent
	Session AgentSession
}

// EventAgentDisconnected is emitted when an agent disconnects
type EventAgentDisconnected struct {
	baseEvent
	Agent   Agent
	Session AgentSession
	Error   error
}

// EventAgentHeartbeatReceived is emitted when a heartbeat is successful
type EventAgentHeartbeatReceived struct {
	baseEvent
	Agent   Agent
	Session AgentSession
	Latency time.Duration
}

// newAgentConnectSucceeded creates a new EventAgentConnectSucceeded event
func newAgentConnectSucceeded(agent Agent, session AgentSession) *EventAgentConnectSucceeded {
	return &EventAgentConnectSucceeded{
		baseEvent: baseEvent{
			Type:       EventTypeAgentConnectSucceeded,
			OccurredAt: time.Now(),
		},
		Agent:   agent,
		Session: session,
	}
}

// newAgentDisconnected creates a new EventAgentDisconnected event
func newAgentDisconnected(agent Agent, session AgentSession, err error) *EventAgentDisconnected {
	return &EventAgentDisconnected{
		baseEvent: baseEvent{
			Type:       EventTypeAgentDisconnected,
			OccurredAt: time.Now(),
		},
		Agent:   agent,
		Session: session,
		Error:   err,
	}
}

// newAgentHeartbeatReceived creates a new EventAgentHeartbeatReceived event
func newAgentHeartbeatReceived(agent Agent, session AgentSession, latency time.Duration) *EventAgentHeartbeatReceived {
	return &EventAgentHeartbeatReceived{
		baseEvent: baseEvent{
			Type:       EventTypeAgentHeartbeatReceived,
			OccurredAt: time.Now(),
		},
		Agent:   agent,
		Session: session,
		Latency: latency,
	}
}

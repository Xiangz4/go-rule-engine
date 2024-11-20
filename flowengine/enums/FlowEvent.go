package enums

type FlowEvent interface {
	// GetEventName returns the name of the event
	GetEventName() string

	// GetEventType returns the type of the event
	GetEventType() string
}

package events

type Consumer interface {
	Start() error
}

type Producer interface {
	ProducerSyncDynamicEvent(dnc DynamicEvent) error
}

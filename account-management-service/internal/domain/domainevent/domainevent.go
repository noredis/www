package domainevent

type DomainEvent interface {
	GetEventName() string
}

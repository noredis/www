package entity

import (
	"account-management-service/internal/domainevent"
	"github.com/google/uuid"
)

type entity struct {
	id           uuid.UUID
	domainEvents []domainevent.DomainEvent
}

func (e *entity) ID() uuid.UUID {
	return e.id
}

func (e *entity) raiseDomainEvent(event domainevent.DomainEvent) {
	e.domainEvents = append(e.domainEvents, event)
}

func (e *entity) ClearDomainEvents() {
	e.domainEvents = make([]domainevent.DomainEvent, 0)
}

func (e *entity) DomainEvents() []domainevent.DomainEvent {
	return e.domainEvents
}

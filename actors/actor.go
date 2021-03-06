package actors

import (
	"fmt"
)

// Actor is the base representation of actor in actor model
type Actor struct {
	identifier     string
	actionsHandler *ActionsHandler
	ch             chan Event
}

// makeActor used to instantiate runner instance
func makeActor(nsp string, actable Actable) *Actor {
	actor := &Actor{
		identifier:     fmt.Sprintf("%v:%v", nsp, actable.GetID()),
		actionsHandler: makeActionsHandler(),
		ch:             make(chan Event, 5),
	}

	actable.RegisterActions(actor.actionsHandler)

	actor.start()

	return actor
}

// Identifier used to identify mailbox's address of actor
func (actor *Actor) Identifier() string {
	return actor.identifier
}

func (actor *Actor) start() {
	go func() {
		for event := range actor.ch {
			actor.actionsHandler.handleEvent(event)
		}
	}()
}

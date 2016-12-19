// Generated by: genem (https://github.com/shiwano/genem)

package context

import (
	"reflect"
	"sync"

	"github.com/shiwano/submarine/server/battle/lib/navmesh"
)

// EventEmitter represents an event emitter.
type EventEmitter struct {
	actorCreateEventMu            *sync.Mutex
	actorCreateEventListeners     []func(actor Actor)
	actorCreateEventListenersOnce []func(actor Actor)

	actorDestroyEventMu            *sync.Mutex
	actorDestroyEventListeners     []func(actor Actor)
	actorDestroyEventListenersOnce []func(actor Actor)

	actorAddEventMu            *sync.Mutex
	actorAddEventListeners     []func(actor Actor)
	actorAddEventListenersOnce []func(actor Actor)

	actorMoveEventMu            *sync.Mutex
	actorMoveEventListeners     []func(actor Actor)
	actorMoveEventListenersOnce []func(actor Actor)

	actorRemoveEventMu            *sync.Mutex
	actorRemoveEventListeners     []func(actor Actor)
	actorRemoveEventListenersOnce []func(actor Actor)

	actorChangeVisibilityEventMu            *sync.Mutex
	actorChangeVisibilityEventListeners     []func(actor Actor, teamLayer navmesh.LayerMask)
	actorChangeVisibilityEventListenersOnce []func(actor Actor, teamLayer navmesh.LayerMask)

	actorUsePingerEventMu            *sync.Mutex
	actorUsePingerEventListeners     []func(actor Actor, finished bool)
	actorUsePingerEventListenersOnce []func(actor Actor, finished bool)
}

// NewEventEmitter creates an event emitter.
func NewEventEmitter() *EventEmitter {
	return &EventEmitter{

		actorCreateEventMu: new(sync.Mutex),

		actorDestroyEventMu: new(sync.Mutex),

		actorAddEventMu: new(sync.Mutex),

		actorMoveEventMu: new(sync.Mutex),

		actorRemoveEventMu: new(sync.Mutex),

		actorChangeVisibilityEventMu: new(sync.Mutex),

		actorUsePingerEventMu: new(sync.Mutex),
	}
}

// EmitActorCreateEvent emits the specified event.
func (_e *EventEmitter) EmitActorCreateEvent(actor Actor) {
	_e.actorCreateEventMu.Lock()
	listeners := make([]func(actor Actor), len(_e.actorCreateEventListeners))
	copy(listeners, _e.actorCreateEventListeners)
	listenersOnce := _e.actorCreateEventListenersOnce
	_e.actorCreateEventListenersOnce = make([]func(actor Actor), 0)
	_e.actorCreateEventMu.Unlock()
	for _, l := range listeners {
		l(actor)
	}
	for _, l := range listenersOnce {
		l(actor)
	}
}

// AddActorCreateEventListener registers the specified event listener.
func (_e *EventEmitter) AddActorCreateEventListener(listener func(actor Actor)) {
	_e.actorCreateEventMu.Lock()
	_e.actorCreateEventListeners = append(_e.actorCreateEventListeners, listener)
	_e.actorCreateEventMu.Unlock()
}

// AddActorCreateEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *EventEmitter) AddActorCreateEventListenerOnce(listener func(actor Actor)) {
	_e.actorCreateEventMu.Lock()
	_e.actorCreateEventListenersOnce = append(_e.actorCreateEventListenersOnce, listener)
	_e.actorCreateEventMu.Unlock()
}

// RemoveActorCreateEventListener removes the event listener previously registered.
func (_e *EventEmitter) RemoveActorCreateEventListener(listener func(actor Actor)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.actorCreateEventMu.Lock()
	listeners := _e.actorCreateEventListeners[:0]
	for _, l := range _e.actorCreateEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.actorCreateEventListeners = listeners
	listenersOnce := _e.actorCreateEventListenersOnce[:0]
	for _, l := range _e.actorCreateEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.actorCreateEventListenersOnce = listenersOnce
	_e.actorCreateEventMu.Unlock()
}

// EmitActorDestroyEvent emits the specified event.
func (_e *EventEmitter) EmitActorDestroyEvent(actor Actor) {
	_e.actorDestroyEventMu.Lock()
	listeners := make([]func(actor Actor), len(_e.actorDestroyEventListeners))
	copy(listeners, _e.actorDestroyEventListeners)
	listenersOnce := _e.actorDestroyEventListenersOnce
	_e.actorDestroyEventListenersOnce = make([]func(actor Actor), 0)
	_e.actorDestroyEventMu.Unlock()
	for _, l := range listeners {
		l(actor)
	}
	for _, l := range listenersOnce {
		l(actor)
	}
}

// AddActorDestroyEventListener registers the specified event listener.
func (_e *EventEmitter) AddActorDestroyEventListener(listener func(actor Actor)) {
	_e.actorDestroyEventMu.Lock()
	_e.actorDestroyEventListeners = append(_e.actorDestroyEventListeners, listener)
	_e.actorDestroyEventMu.Unlock()
}

// AddActorDestroyEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *EventEmitter) AddActorDestroyEventListenerOnce(listener func(actor Actor)) {
	_e.actorDestroyEventMu.Lock()
	_e.actorDestroyEventListenersOnce = append(_e.actorDestroyEventListenersOnce, listener)
	_e.actorDestroyEventMu.Unlock()
}

// RemoveActorDestroyEventListener removes the event listener previously registered.
func (_e *EventEmitter) RemoveActorDestroyEventListener(listener func(actor Actor)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.actorDestroyEventMu.Lock()
	listeners := _e.actorDestroyEventListeners[:0]
	for _, l := range _e.actorDestroyEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.actorDestroyEventListeners = listeners
	listenersOnce := _e.actorDestroyEventListenersOnce[:0]
	for _, l := range _e.actorDestroyEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.actorDestroyEventListenersOnce = listenersOnce
	_e.actorDestroyEventMu.Unlock()
}

// EmitActorAddEvent emits the specified event.
func (_e *EventEmitter) EmitActorAddEvent(actor Actor) {
	_e.actorAddEventMu.Lock()
	listeners := make([]func(actor Actor), len(_e.actorAddEventListeners))
	copy(listeners, _e.actorAddEventListeners)
	listenersOnce := _e.actorAddEventListenersOnce
	_e.actorAddEventListenersOnce = make([]func(actor Actor), 0)
	_e.actorAddEventMu.Unlock()
	for _, l := range listeners {
		l(actor)
	}
	for _, l := range listenersOnce {
		l(actor)
	}
}

// AddActorAddEventListener registers the specified event listener.
func (_e *EventEmitter) AddActorAddEventListener(listener func(actor Actor)) {
	_e.actorAddEventMu.Lock()
	_e.actorAddEventListeners = append(_e.actorAddEventListeners, listener)
	_e.actorAddEventMu.Unlock()
}

// AddActorAddEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *EventEmitter) AddActorAddEventListenerOnce(listener func(actor Actor)) {
	_e.actorAddEventMu.Lock()
	_e.actorAddEventListenersOnce = append(_e.actorAddEventListenersOnce, listener)
	_e.actorAddEventMu.Unlock()
}

// RemoveActorAddEventListener removes the event listener previously registered.
func (_e *EventEmitter) RemoveActorAddEventListener(listener func(actor Actor)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.actorAddEventMu.Lock()
	listeners := _e.actorAddEventListeners[:0]
	for _, l := range _e.actorAddEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.actorAddEventListeners = listeners
	listenersOnce := _e.actorAddEventListenersOnce[:0]
	for _, l := range _e.actorAddEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.actorAddEventListenersOnce = listenersOnce
	_e.actorAddEventMu.Unlock()
}

// EmitActorMoveEvent emits the specified event.
func (_e *EventEmitter) EmitActorMoveEvent(actor Actor) {
	_e.actorMoveEventMu.Lock()
	listeners := make([]func(actor Actor), len(_e.actorMoveEventListeners))
	copy(listeners, _e.actorMoveEventListeners)
	listenersOnce := _e.actorMoveEventListenersOnce
	_e.actorMoveEventListenersOnce = make([]func(actor Actor), 0)
	_e.actorMoveEventMu.Unlock()
	for _, l := range listeners {
		l(actor)
	}
	for _, l := range listenersOnce {
		l(actor)
	}
}

// AddActorMoveEventListener registers the specified event listener.
func (_e *EventEmitter) AddActorMoveEventListener(listener func(actor Actor)) {
	_e.actorMoveEventMu.Lock()
	_e.actorMoveEventListeners = append(_e.actorMoveEventListeners, listener)
	_e.actorMoveEventMu.Unlock()
}

// AddActorMoveEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *EventEmitter) AddActorMoveEventListenerOnce(listener func(actor Actor)) {
	_e.actorMoveEventMu.Lock()
	_e.actorMoveEventListenersOnce = append(_e.actorMoveEventListenersOnce, listener)
	_e.actorMoveEventMu.Unlock()
}

// RemoveActorMoveEventListener removes the event listener previously registered.
func (_e *EventEmitter) RemoveActorMoveEventListener(listener func(actor Actor)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.actorMoveEventMu.Lock()
	listeners := _e.actorMoveEventListeners[:0]
	for _, l := range _e.actorMoveEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.actorMoveEventListeners = listeners
	listenersOnce := _e.actorMoveEventListenersOnce[:0]
	for _, l := range _e.actorMoveEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.actorMoveEventListenersOnce = listenersOnce
	_e.actorMoveEventMu.Unlock()
}

// EmitActorRemoveEvent emits the specified event.
func (_e *EventEmitter) EmitActorRemoveEvent(actor Actor) {
	_e.actorRemoveEventMu.Lock()
	listeners := make([]func(actor Actor), len(_e.actorRemoveEventListeners))
	copy(listeners, _e.actorRemoveEventListeners)
	listenersOnce := _e.actorRemoveEventListenersOnce
	_e.actorRemoveEventListenersOnce = make([]func(actor Actor), 0)
	_e.actorRemoveEventMu.Unlock()
	for _, l := range listeners {
		l(actor)
	}
	for _, l := range listenersOnce {
		l(actor)
	}
}

// AddActorRemoveEventListener registers the specified event listener.
func (_e *EventEmitter) AddActorRemoveEventListener(listener func(actor Actor)) {
	_e.actorRemoveEventMu.Lock()
	_e.actorRemoveEventListeners = append(_e.actorRemoveEventListeners, listener)
	_e.actorRemoveEventMu.Unlock()
}

// AddActorRemoveEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *EventEmitter) AddActorRemoveEventListenerOnce(listener func(actor Actor)) {
	_e.actorRemoveEventMu.Lock()
	_e.actorRemoveEventListenersOnce = append(_e.actorRemoveEventListenersOnce, listener)
	_e.actorRemoveEventMu.Unlock()
}

// RemoveActorRemoveEventListener removes the event listener previously registered.
func (_e *EventEmitter) RemoveActorRemoveEventListener(listener func(actor Actor)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.actorRemoveEventMu.Lock()
	listeners := _e.actorRemoveEventListeners[:0]
	for _, l := range _e.actorRemoveEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.actorRemoveEventListeners = listeners
	listenersOnce := _e.actorRemoveEventListenersOnce[:0]
	for _, l := range _e.actorRemoveEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.actorRemoveEventListenersOnce = listenersOnce
	_e.actorRemoveEventMu.Unlock()
}

// EmitActorChangeVisibilityEvent emits the specified event.
func (_e *EventEmitter) EmitActorChangeVisibilityEvent(actor Actor, teamLayer navmesh.LayerMask) {
	_e.actorChangeVisibilityEventMu.Lock()
	listeners := make([]func(actor Actor, teamLayer navmesh.LayerMask), len(_e.actorChangeVisibilityEventListeners))
	copy(listeners, _e.actorChangeVisibilityEventListeners)
	listenersOnce := _e.actorChangeVisibilityEventListenersOnce
	_e.actorChangeVisibilityEventListenersOnce = make([]func(actor Actor, teamLayer navmesh.LayerMask), 0)
	_e.actorChangeVisibilityEventMu.Unlock()
	for _, l := range listeners {
		l(actor, teamLayer)
	}
	for _, l := range listenersOnce {
		l(actor, teamLayer)
	}
}

// AddActorChangeVisibilityEventListener registers the specified event listener.
func (_e *EventEmitter) AddActorChangeVisibilityEventListener(listener func(actor Actor, teamLayer navmesh.LayerMask)) {
	_e.actorChangeVisibilityEventMu.Lock()
	_e.actorChangeVisibilityEventListeners = append(_e.actorChangeVisibilityEventListeners, listener)
	_e.actorChangeVisibilityEventMu.Unlock()
}

// AddActorChangeVisibilityEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *EventEmitter) AddActorChangeVisibilityEventListenerOnce(listener func(actor Actor, teamLayer navmesh.LayerMask)) {
	_e.actorChangeVisibilityEventMu.Lock()
	_e.actorChangeVisibilityEventListenersOnce = append(_e.actorChangeVisibilityEventListenersOnce, listener)
	_e.actorChangeVisibilityEventMu.Unlock()
}

// RemoveActorChangeVisibilityEventListener removes the event listener previously registered.
func (_e *EventEmitter) RemoveActorChangeVisibilityEventListener(listener func(actor Actor, teamLayer navmesh.LayerMask)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.actorChangeVisibilityEventMu.Lock()
	listeners := _e.actorChangeVisibilityEventListeners[:0]
	for _, l := range _e.actorChangeVisibilityEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.actorChangeVisibilityEventListeners = listeners
	listenersOnce := _e.actorChangeVisibilityEventListenersOnce[:0]
	for _, l := range _e.actorChangeVisibilityEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.actorChangeVisibilityEventListenersOnce = listenersOnce
	_e.actorChangeVisibilityEventMu.Unlock()
}

// EmitActorUsePingerEvent emits the specified event.
func (_e *EventEmitter) EmitActorUsePingerEvent(actor Actor, finished bool) {
	_e.actorUsePingerEventMu.Lock()
	listeners := make([]func(actor Actor, finished bool), len(_e.actorUsePingerEventListeners))
	copy(listeners, _e.actorUsePingerEventListeners)
	listenersOnce := _e.actorUsePingerEventListenersOnce
	_e.actorUsePingerEventListenersOnce = make([]func(actor Actor, finished bool), 0)
	_e.actorUsePingerEventMu.Unlock()
	for _, l := range listeners {
		l(actor, finished)
	}
	for _, l := range listenersOnce {
		l(actor, finished)
	}
}

// AddActorUsePingerEventListener registers the specified event listener.
func (_e *EventEmitter) AddActorUsePingerEventListener(listener func(actor Actor, finished bool)) {
	_e.actorUsePingerEventMu.Lock()
	_e.actorUsePingerEventListeners = append(_e.actorUsePingerEventListeners, listener)
	_e.actorUsePingerEventMu.Unlock()
}

// AddActorUsePingerEventListenerOnce registers the specified event listener that is invoked only once.
func (_e *EventEmitter) AddActorUsePingerEventListenerOnce(listener func(actor Actor, finished bool)) {
	_e.actorUsePingerEventMu.Lock()
	_e.actorUsePingerEventListenersOnce = append(_e.actorUsePingerEventListenersOnce, listener)
	_e.actorUsePingerEventMu.Unlock()
}

// RemoveActorUsePingerEventListener removes the event listener previously registered.
func (_e *EventEmitter) RemoveActorUsePingerEventListener(listener func(actor Actor, finished bool)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	_e.actorUsePingerEventMu.Lock()
	listeners := _e.actorUsePingerEventListeners[:0]
	for _, l := range _e.actorUsePingerEventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	_e.actorUsePingerEventListeners = listeners
	listenersOnce := _e.actorUsePingerEventListenersOnce[:0]
	for _, l := range _e.actorUsePingerEventListenersOnce {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listenersOnce = append(listenersOnce, l)
		}
	}
	_e.actorUsePingerEventListenersOnce = listenersOnce
	_e.actorUsePingerEventMu.Unlock()
}

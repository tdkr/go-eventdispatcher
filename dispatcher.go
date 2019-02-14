/*
@Time : 2018/10/18 10:46
@Author : RonanLuo
*/
package event

type Event string

type EventHandler func(...interface{})

type EventLisenter struct {
	handler EventHandler
}

type EventDispatcher struct {
	listeners map[Event][]*EventLisenter
}

func NewEventListener(handler EventHandler) *EventLisenter {
	return &EventLisenter{
		handler: handler,
	}
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		listeners: map[Event][]*EventLisenter{},
	}
}

func (dispatcher *EventDispatcher) Register(e Event, listener *EventLisenter) {
	if _, ok := dispatcher.listeners[e]; !ok {
		dispatcher.listeners[e] = []*EventLisenter{}
	}
	dispatcher.listeners[e] = append(dispatcher.listeners[e], listener)
}

func (dispatcher *EventDispatcher) Unregister(e Event, listener *EventLisenter) {
	if arr, ok := dispatcher.listeners[e]; ok {
		for i, v := range arr {
			if v == listener {
				dispatcher.listeners[e] = append(dispatcher.listeners[e][:i], dispatcher.listeners[e][i+1:]...)
				return
			}
		}
	}
}

func (dispatcher *EventDispatcher) Dispatch(e Event, args ...interface{}) {
	if listeners, ok := dispatcher.listeners[e]; ok {
		for _, listener := range listeners {
			listener.handler(args...)
		}
	}
}

package topic

import (
	"sync"

	"github.com/Rayvid/go_firstattempt/internal/message"
)

// Topic defines logic around actual topic
type Topic struct {
	Name      string
	messages  []message.Message
	listeners []chan<- message.Message
	syncRoot  sync.Locker
}

// Internal topics list management stuff
var syncRoot sync.Mutex
var topics map[string]*Topic = make(map[string]*Topic)

// Subscribe to the messages on this topic
func (topic *Topic) Subscribe(ch chan<- message.Message) {
	topic.syncRoot.Lock()
	defer topic.syncRoot.Unlock()

	topic.listeners = append(topic.listeners, ch)
	for _, msg := range topic.messages {
		// Do not block when already under lock
		go func(listener chan<- message.Message, msg message.Message) { listener <- msg }(ch, msg)
	}
}

// Unsubscribe to this topic
func (topic *Topic) Unsubscribe(ch chan<- message.Message) {
	topic.syncRoot.Lock()
	defer topic.syncRoot.Unlock()

	for i, listener := range topic.listeners {
		if listener == ch {
			topic.listeners[i] = nil
			topic.listeners = append(topic.listeners[:i], topic.listeners[i+1:]...)

			if len(topic.listeners) == 0 { // Abandod topic if no listeners
				syncRoot.Lock()
				defer syncRoot.Unlock()

				// Do not keep empty topic
				delete(topics, topic.Name)
			}

			break
		}
	}
}

// Post message to all listeners
func (topic *Topic) Post(msg message.Message) {
	topic.syncRoot.Lock()
	defer topic.syncRoot.Unlock()

	if len(topic.listeners) > 0 {
		topic.messages = append(topic.messages, msg)
		for _, listener := range topic.listeners {
			// Do not block when already under lock
			go func(listener chan<- message.Message, msg message.Message) { listener <- msg }(listener, msg)
		}
	} else {
		syncRoot.Lock()
		defer syncRoot.Unlock()

		// Do not keep empty topic
		delete(topics, topic.Name)
	}
}

// GetOrCreate performs autocreate if needed
func GetOrCreate(name string) (result *Topic) {
	syncRoot.Lock()
	defer syncRoot.Unlock()

	t, ok := topics[name]
	if !ok {
		t = &Topic{Name: name, messages: make([](message.Message), 0), listeners: make([](chan<- message.Message), 0), syncRoot: &sync.Mutex{}}
		topics[name] = t
	}

	return t
}

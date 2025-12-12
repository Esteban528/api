package events

import (
	"errors"
)

var pubs = make(map[string]any)

func init() {
	pubs = make(map[string]any)
}

type Publisher[T any] struct {
	Subs []Subscriber[T]
}

type Subscriber[T any] interface {
	Update(T)
}

func (p *Publisher[T]) NotifyAll(newStatus T) {
	for _, sub := range p.Subs {
		sub.Update(newStatus)
	}
}

func AddPublisher[T any](key string) error {
	if _, exists := pubs[key]; !exists {
		pubs[key] = &Publisher[T]{Subs: []Subscriber[T]{}}
		return nil
	}

	return errors.New("key is already a subscriber")
}

func Get[T any](key string) *Publisher[T] {
	value := pubs[key].(*Publisher[T])
	return value
}

func Subscribe[T any](key string, sub Subscriber[T]) error {
	if _, exists := pubs[key]; !exists {
		return errors.New("key doesn't exist")
	}

	pub := pubs[key].(*Publisher[T])
	pub.Subs = append(pub.Subs, sub)
	return nil
}

func NotifyAll[T any](key string, value T) error {
	if _, exists := pubs[key]; !exists {
		return errors.New("key doesn't exist")
	}

	go Get[T](key).NotifyAll(value)
	return nil
}

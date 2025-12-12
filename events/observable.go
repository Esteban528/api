package events

import "errors"

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
		pubs[key] = make([]Subscriber[T], 0)
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

	subs := pubs[key].([]Subscriber[T])
	subs = append(subs, sub)
	pubs[key] = subs
	return nil
}

func NotifyAll[T any](key string, value T) error {
	if _, exists := pubs[key]; !exists {
		return errors.New("key doesn't exist")
	}

	Get[T](key).NotifyAll(value)
	return nil
}

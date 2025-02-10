package data

type FakePublisher struct {
	messages []any
}

func NewFakePublisher() *FakePublisher {
	return &FakePublisher{}
}

func (w *FakePublisher) Publish(message any) error {
	w.messages = append(w.messages, message)
	return nil
}

func (w *FakePublisher) Messages() []any {
	return w.messages
}

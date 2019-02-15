package broker

type MessageHandler func([]byte) error

type Publisher interface {
	Publish(topic string, message []byte) error
}

type Subscriber interface {
	Subscribe(topic string, handler MessageHandler) error
}

type Broker interface {
	Publisher
	Subscriber
}

package nats

import (
	"github.com/ConsenSys/broker"
	xnats "github.com/nats-io/go-nats"
)

type natsBroker struct {
	nc *xnats.Conn
}

func NewBroker(url string) broker.Broker{
	nb := &natsBroker{}

	go func() {
		nc, _ := xnats.Connect(url)
		nb.nc = nc
	}()

	return nb
}

func (n natsBroker) Publish(topic string, message []byte) error {
	return n.nc.Publish(topic, message)
}

func (n natsBroker) Subscribe(topic string, handler broker.MessageHandler) error {
	_, err := n.nc.Subscribe(topic, func(msg *xnats.Msg) {
		handler(msg.Data)
	})
	return err
}



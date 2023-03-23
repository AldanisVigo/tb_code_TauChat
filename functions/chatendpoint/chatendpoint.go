package lib

import (
	"fmt"
	"github.com/taubyte/go-sdk/event"
	"github.com/taubyte/go-sdk/pubsub/node"
)

func chat(event event.Event) error {
	h, err := event.HTTP()
	
	if err != nil {
		return err
	}

	channel, err := node.Channel("/chatendpoint")
	if err != nil {
		return err
	}

	url, err := channel.WebSocket().Url()
	if err != nil {
		return err
	}

	_, err = h.Write([]byte("{\"socket\":"\"" + url.Path + "\"}"))
	if err != nil {
		return err
	}

	return nil

}

func chatendpoint(e event.Event) uint32 {
	h, err :=  event.HTTP()

	if err != nil {
		h.Write([]byte(fmt.Sprintf("ERROR: %s\n", err)))
	}

	err = chat(e)
	if err != nil {
		return 1
	}

	return 0
}
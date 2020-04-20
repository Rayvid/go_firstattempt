package get

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Rayvid/go_firstattempt/internal/message"
	"github.com/Rayvid/go_firstattempt/internal/topic"
)

// HandleRequest handles event stream post requests
func HandleRequest(w http.ResponseWriter, r *http.Request, topicName string) (err error) {
	t := topic.GetOrCreate(topicName)

	ch := make(chan message.Message)
	defer t.Unsubscribe(ch)

	timeout := make(chan struct{})
	t.Subscribe(ch)

	go func() {
		defer close(timeout)
		time.Sleep(30 * time.Second) // TODO make this configurable
	}()

loop:
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				ch = nil
			} else {
				_, err = fmt.Fprintf(w, "id: %v\nevent: msg\ndata: %v\n", msg.ID, msg.Text)
			}
		case _, _ = <-timeout:
			_, err = fmt.Fprintf(w, "event: timeout\ndata: 30s\n")
			ch = nil

			if ch == nil {
				break loop
			}
		}
	}

	return err
}

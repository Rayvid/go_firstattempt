package post

import (
	"io/ioutil"
	"net/http"
	"sync/atomic"

	"github.com/Rayvid/go_firstattempt/internal/message"
	"github.com/Rayvid/go_firstattempt/internal/topic"
)

var autoincrement int64

// HandleRequest handles event stream get requests
func HandleRequest(w http.ResponseWriter, r *http.Request, topicName string) (err error) {
	t := topic.GetOrCreate(topicName)
	content, err := ioutil.ReadAll(r.Body)

	if err == nil {
		t.Post(message.Message{ID: atomic.AddInt64(&autoincrement, 1), Text: string(content)})
	}

	return err
}

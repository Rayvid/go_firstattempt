package post

import (
	"fmt"
	"net/http"
)

// HandleRequest handles event stream get requests
func HandleRequest(w http.ResponseWriter, r *http.Request, topic string) (err error) {
	_, err = fmt.Fprintf(w, "Handling post of `%v`", topic)

	return err
}

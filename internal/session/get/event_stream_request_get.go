package get

import (
	"fmt"
	"net/http"
)

// HandleRequest handles event stream post requests
func HandleRequest(w http.ResponseWriter, r *http.Request, topic string) (err error) {
	_, err = fmt.Fprintf(w, "Handling get of `%v`", topic)

	return err
}

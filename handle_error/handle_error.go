package handle_error

import (
	"errors"
	"net/http"

	"github.com/rayhan889/rss-aggr/handle_json"
)

func HandleError(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Something went wrong")
	handle_json.RespondWithError(w, http.StatusBadRequest, err)
}
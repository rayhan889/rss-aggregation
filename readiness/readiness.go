package readiness

import (
	"net/http"

	"github.com/rayhan889/rss-aggr/handle_json"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	handle_json.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

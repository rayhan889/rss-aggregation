package readiness

import (
	"net/http"

	"github.com/rayhan889/rss-aggr/json"
)

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	json.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

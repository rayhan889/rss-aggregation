package feeds

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/rayhan889/rss-aggr/handle_json"
	"github.com/rayhan889/rss-aggr/internal/database"
	"github.com/rayhan889/rss-aggr/models/feed"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apf *ApiConfig) HandleCreateNewFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type feedParamaters struct {
		Name    string `json:"name"`
		Url     string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := feedParamaters{}
	err := decoder.Decode(&params)

	if err != nil {
		handle_json.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	feedDt, err := apf.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})
	if err != nil {
		handle_json.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	handle_json.RespondWithJSON(w, 201, feed.HandleFeedToFeedCustomModel(feedDt))

}

func (apf *ApiConfig) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apf.DB.GetFeeds(r.Context())
	if err != nil {
		handle_json.RespondWithError(w, 404, err)
		return
	}
	
	handle_json.RespondWithJSON(w, 200, feed.HandleFeedsToFeedsCustomModel(feeds))
}

func (apf *ApiConfig) HandleGetFeedsByUserID(w http.ResponseWriter, r *http.Request, userDT database.User) {
	userIDStr := chi.URLParam(r, "userID")
	userID, parseErr := uuid.Parse(userIDStr)
	if parseErr != nil {
		handle_json.RespondWithError(w, 400, parseErr)
		return
	}

	if userDT.ID != userID {
		handle_json.RespondWithError(w, 403, errors.New("You are not authorized to access this resource"))
		return
	}

	feeds, err := apf.DB.GetFeedsByUserID(r.Context(), userID)
	if err != nil {
		handle_json.RespondWithError(w, 404, err)
		return
	}

	handle_json.RespondWithJSON(w, 200, feed.HandleFeedsToFeedsCustomModel(feeds))
}
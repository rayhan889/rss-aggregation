package users

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/rayhan889/rss-aggr/handle_json"
	"github.com/rayhan889/rss-aggr/internal/database"
	"github.com/rayhan889/rss-aggr/models/user"
)

type ApiConfig struct {
	DB *database.Queries
}

func (apf *ApiConfig) HandleCreateNewUser(w http.ResponseWriter, r *http.Request) {
	type userParamaters struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := userParamaters{}
	err := decoder.Decode(&params)

	if err != nil {
		handle_json.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	userDt, err := apf.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Email: params.Email,
		Name: params.Name,
		Password: params.Password,
	})
	if err != nil {
		handle_json.RespondWithError(w, http.StatusInternalServerError, err)
		return
	}

	handle_json.RespondWithJSON(w, http.StatusOK, user.HandleUserToUserCustomModel(userDt))

}
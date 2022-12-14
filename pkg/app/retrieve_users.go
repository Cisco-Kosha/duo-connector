package app

import (
	"net/http"

	"github.com/kosha/duo-connector/pkg/httpclient"
)

// getRetreive Users godoc
// @Summary Get Returns a paged list of users
// @Description Retrieve Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Router /admin/v1/users [get]
func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	generate_url := a.Cfg.GetDuoSecurityURL() + "/admin/v1/users"
	data, err := httpclient.GetRequest(generate_url, a.Cfg.GetPersonalAccessToken())
	if err != nil {
		a.Log.Errorf("Error in Duo API call getUsers", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, data)
}

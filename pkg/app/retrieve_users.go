package app

import (
	"net/http"

	"github.com/kosha/duo-connector/pkg/httpclient"
)

// getRetreive godoc
// @Summary Get Returns a paged list of users
// @Description all the users
// @Tags specification
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v2/admin/v1/users [get]
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
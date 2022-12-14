package app


import (
	"github.com/gorilla/mux"
	"github.com/kosha/duo-connector/pkg/httpclient"
	"net/http"
)

// getRetreive Users godoc
// @Summary Get Returns a paged list of users
// @Description Retrieve Users
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200
// @Router /admin/v1/users [get]
func (a *App) getBases(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	generate_url := a.Cfg.GetDuoSecurityURL + "/admin/v1/users"
	bases := httpclient.GetRequest(generate_url, a.Cfg.GetPersonalAccessToken())
	respondWithJSON(w, http.StatusOK, bases)
}

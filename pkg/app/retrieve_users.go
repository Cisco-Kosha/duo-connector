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
// @Success 200
// @Router /api/v1/admin/v1/users [get]
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

// To generate swagger.json and swagger.yaml files based on the API documentation, simple run -

// go install github.com/swaggo/swag/cmd/swag@latest
// swag init -g main.go --parseDependency --parseInternal
// To generate OpenAPISpec version 3 from Swagger 2.0 specification, run -

// npm i api-spec-converter
// npx api-spec-converter --from=swagger_2 --to=openapi_3 --syntax=json ./docs/swagger.json > openapi.json

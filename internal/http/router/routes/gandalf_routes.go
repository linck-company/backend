package router

import (
	"backendv1/internal/http/handlers/gandalfhandlers"
	"fmt"
	"net/http"
)

func InitGandalfRoutes(mux *http.ServeMux, dsn string, apiPrefix string) {
	gho := gandalfhandlers.NewGandalfHandler(dsn)

	mux.HandleFunc(fmt.Sprintf("GET %s/gandalf/entity/details", apiPrefix), gho.GetEntityDetails)
	mux.HandleFunc(fmt.Sprintf("POST %s/gandalf/entity/create/entity", apiPrefix), gho.CreateEntityDetails)
	mux.HandleFunc(fmt.Sprintf("GET %s/gandalf/user/registered_entity", apiPrefix), gho.GetUserRegisteredEntity)
	mux.HandleFunc(fmt.Sprintf("POST %s/gandalf/entity/legacy_holders", apiPrefix), gho.GetLegacyHolders)
}

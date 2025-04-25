package router

import (
	odysseyhandler "backendv1/internal/http/handlers/odysseyhandlers"
	"fmt"
	"net/http"
)

func InitOdysseyRoutes(mux *http.ServeMux, dsn string, apiPrefix string) {
	oho := odysseyhandler.NewOdysseyHandler()

	mux.HandleFunc(fmt.Sprintf("GET %s/odyssey/report/details", apiPrefix), oho.GetReportDetails)
}

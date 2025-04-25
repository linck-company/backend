package router

import (
	"backendv1/internal/http/handlers/hecatehandlers"
	"fmt"
	"net/http"
)

func InitHecateRoutes(mux *http.ServeMux, dsn string, apiPrefix string) {
	hho := hecatehandlers.NewHecateHandlers(dsn)

	mux.HandleFunc(fmt.Sprintf("GET %s/hecate/events/details", apiPrefix), hho.GetEventDetails)
	mux.HandleFunc(fmt.Sprintf("POST %s/hecate/events/register", apiPrefix), hho.RegisterForEvent)
	mux.HandleFunc(fmt.Sprintf("POST %s/hecate/events/unregister", apiPrefix), hho.UnRegisterForEvent)
	mux.HandleFunc(fmt.Sprintf("POST %s/hecate/events/create/event", apiPrefix), hho.CreateEventDetails)
	mux.HandleFunc(fmt.Sprintf("POST %s/hecate/students/event/records", apiPrefix), hho.GetStudentRecords)
}

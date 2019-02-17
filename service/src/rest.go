package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

//InitRestHandler Supported REST route handing
func InitRestHandler() {
	prefix := "rest"
	router := mux.NewRouter()
	router.PathPrefix(prefix)
	router.HandleFunc("/"+prefix+"/player/register/", registerPlayerHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/"+prefix+"/player/unregister/", unregisterPlayerHandler).Methods("POST", "OPTIONS")
	http.Handle("/", router)
}

type _RegisterPlayerBody struct {
	Name string `json:"name"`
}

type _RegisterPlayerResponse struct {
	ID string `json:"id"`
}

func registerPlayerHandler(w http.ResponseWriter, req *http.Request) {
	var body _RegisterPlayerBody

	if err := ReadBytesFromBody(req.Body, &body); err != nil {
		SendErrorResponse(w, err)
		return
	}

	id, err := GameInstance.addPlayerWithName(body.Name)
	if err != nil {
		SendErrorResponse(w, err)
		return
	}

	response := _RegisterPlayerResponse{
		ID: *id,
	}

	SendResponse(w, response)
}

func unregisterPlayerHandler(w http.ResponseWriter, req *http.Request) {
	var body _RegisterPlayerResponse

	if err := ReadBytesFromBody(req.Body, &body); err != nil {
		SendErrorResponse(w, err)
		return
	}

	GameInstance.removePlayerWithID(body.ID)
	SendResponse(w, nil)
}

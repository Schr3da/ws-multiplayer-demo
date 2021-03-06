package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//Response Basic Response type
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

//NewErrorResponse Default error response
func NewErrorResponse() []byte {
	d := Response{
		Status: http.StatusInternalServerError,
		Data:   nil,
	}
	data, _ := json.Marshal(&d)
	return data
}

//NewResponse Create a new response json Object
func NewResponse(status int, data interface{}) []byte {
	d, err := json.Marshal(Response{
		Status: status,
		Data:   data,
	})

	if err != nil {
		CatchError("New Response", err)
		return NewErrorResponse()
	}

	return d
}

//SendErrorResponse Default error response construction
func SendErrorResponse(w http.ResponseWriter, err error) {
	CatchError("SendErrorResponse", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(NewErrorResponse())
}

//SendResponse Default response construction
func SendResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusAccepted)
	w.Write(NewResponse(http.StatusAccepted, data))
}

//ReadBytesFromBody Get Data from body
func ReadBytesFromBody(body io.Reader, dest interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		CatchError("ReadBytesFromBody", err)
		return err
	}

	if err := json.Unmarshal(data, &dest); err != nil {
		CatchError("ReadBytesFromBody", err)
		return err
	}

	return nil
}

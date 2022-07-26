package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"ozonTask/pkg/error"
	"ozonTask/pkg/link"
)

type Handler struct {
	Repo link.LinkStorage
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/createShort", h.GetShortLink)
	mux.HandleFunc("/getOriginal", h.GetOriginalLink)
	return mux
}

func (h *Handler) GetShortLink(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		log.Printf("wrong http method: %s", req.Method)
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: "bad http method, must be POST"})
		return
	}
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Printf("read request body error: %s", err.Error())
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: "no data was received"})
		return
	}

	var originalLink link.Link

	if err = json.Unmarshal(body, &originalLink); err != nil {
		log.Printf("parse json: %s", err.Error())
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: "wrong input data"})
		return
	}

	if _, err = url.ParseRequestURI(originalLink.Data); err != nil {
		log.Printf("parse url: %s", err.Error())
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: "wrong input data"})
		return
	}

	shortLink, err := h.Repo.Add(originalLink.Data)
	if err != nil {
		log.Printf("add new link: %s", err.Error())
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: err.Error()})
		return
	}

	NewLinkResponse(w, http.StatusOK, shortLink)

}

func (h *Handler) GetOriginalLink(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		log.Printf("wrong http method: %s", req.Method)
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: "bad http method, must be GET"})
		return
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Printf("read request body error: %s", err.Error())
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: "no data was received"})
		return
	}

	var shortLink link.Link

	if err = json.Unmarshal(body, &shortLink); err != nil {
		log.Printf("parse json: %s", err.Error())
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: "wrong input data"})
		return
	}

	originalLink, err := h.Repo.Get(shortLink.Data)
	if err != nil {
		log.Printf("get original link: %s", err.Error())
		NewErrorResponse(w, http.StatusBadRequest, error.ResponseError{Message: err.Error()})
		return
	}

	NewLinkResponse(w, http.StatusOK, originalLink)

}

func NewErrorResponse(w http.ResponseWriter, httpCode int, errMsg error.ResponseError) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	_, err := w.Write(errMsg.Error())

	if err != nil {
		log.Printf("response error: %s", err.Error())
	}
}

func NewLinkResponse(w http.ResponseWriter, httpCode int, result string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	jsonResponse, err := json.Marshal(&link.Link{Data: result})

	if err != nil {
		log.Printf("response error: %s", err.Error())
	}

	_, err = w.Write(jsonResponse)

	if err != nil {
		log.Printf("response error: %s", err.Error())
	}
}

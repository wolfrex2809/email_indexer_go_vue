package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/wolfrex2809/email_indexer_go_vue/models"
	"github.com/wolfrex2809/email_indexer_go_vue/service"
)

func IndexEmails(w http.ResponseWriter, r *http.Request) {

	err := service.IndexEmails()

	if err != nil {
		log.Println("There was an error trying to index all emails: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func SearchEmails(w http.ResponseWriter, r *http.Request) {

	// Get "text" param from the Request's query params
	searchType := "matchphrase"
	text := r.URL.Query().Get("text")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	page := 0

	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.CustomError{
			Status: "Failed",
			Msg:    err.Error(),
		})
	}

	if from != "" || to != "" {
		if from != "" {
			text = fmt.Sprintf("%s +from:%s", text, from)
		}

		if to != "" {
			text = fmt.Sprintf("%s +to:%s", text, to)
		}
		searchType = "querystring"
	}

	// Call the search function using the "text"
	emails, total, err := service.SearchEmails(text, searchType, page)

	if err != nil {
		// Prapare the response if search wasn't successfull
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.CustomError{
			Status: "Failed",
			Msg:    err.Error(),
		})
	} else {
		// Prapare the response if search was successfull
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.EmailSearchResponse{
			Status: "Success",
			Data:   emails,
			Total:  total,
		})
	}
}

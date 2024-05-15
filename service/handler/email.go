package handler

import (
	"log"
	"net/http"

	"github.com/wolfrex2809/email_indexer_go_vue/service"
)

type Email struct{}

func (e *Email) IndexEmails(w http.ResponseWriter, r *http.Request) {

	err := service.IndexEmails()

	if err != nil {
		log.Println("There was an error tryiing to index all emails: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

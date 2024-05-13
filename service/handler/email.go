package handler

import (
	"fmt"
	"net/http"
)

type Email struct{}

func (e *Email) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("It is listing the emails...")
}

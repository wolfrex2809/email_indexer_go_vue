package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/wolfrex2809/email_indexer_go_vue/models"
)

func GetAllUsers() ([]string, error) {

	var users []string
	files, err := os.ReadDir("./enron_mail_20110402/maildir")

	if err != nil {
		return nil, err
	}

	for _, item := range files {
		if item.IsDir() {
			users = append(users, item.Name())
		}
	}

	return users, nil
}

func IndexEmails() error {

	users, err := GetAllUsers()

	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, user := range users {
		wg.Add(1)
		go IndexEmailsByUser(user, &wg)
	}

	wg.Wait()

	return nil
}

func IndexEmailsByUser(user string, wg *sync.WaitGroup) {

	defer wg.Done()

	emails, err := GetEmailsByUser(user)

	if err != nil {
		log.Println("Error extracting emails from user: ", user)
		return
	}
}

func GetEmailsByUser(user string) ([]models.Email, error) {

	var emails []models.Email

	userPath := fmt.Sprintf("%s/%s", "./enron_mail_20110402/maildir", user)

	err := filepath.Walk(userPath, CheckEmailsFiles(&emails))

	if err != nil {
		log.Println(fmt.Sprintf("There was an error trying to read emails from '%s' : ", user), err)
	}
	return emails, nil
}

func CheckEmailsFiles(emails *[]models.Email) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("There was an error checking emails files: ", err)
		}

		if !info.IsDir() {
			email, err := ProcessEmailData(path)

			if err != nil || email == nil {
				return nil
			}

			*emails = append(*emails, *email)
		}

		return nil
	}
}

func ProcessEmailData(path string) (*models.Email, error) {

	email := &models.Email{}

	file, err := os.ReadFile(filepath.Clean(path))

	if err != nil {
		log.Println("There was an error reading emails files:")
	}

	arr := strings.SplitN(string(file), "\r\n\r\n", 2)

	if len(arr) != 2 {
		log.Println("There was an error processing this email file: ", path)
		return nil, err
	}

	email.Content = arr[1]

	emailDetails := strings.Split(arr[0], "\r\n")

	for i := 0; i < len(emailDetails); i++ {
		detail := strings.SplitN(emailDetails[i], ": ", 2)
		switch detail[0] {
		case "Message-ID":
			email.MessageID = detail[1]
		case "Date":
			email.Date = detail[1]
		case "From":
			email.From = detail[1]
		case "To":
			email.To = detail[1]
		case "Subject":
			email.Subject = detail[1]
		default:
			continue
		}
	}

	return email, nil

}

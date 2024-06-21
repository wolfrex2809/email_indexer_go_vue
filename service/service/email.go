package service

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/wolfrex2809/email_indexer_go_vue/client"
	"github.com/wolfrex2809/email_indexer_go_vue/config"
	"github.com/wolfrex2809/email_indexer_go_vue/models"
)

var TarSupportedExt = []string{
	"gz",
	"tgz",
	"bz",
	"tbz",
	"bz2",
	"tbz2",
	"lzo",
	"tzo",
	"7z",
	"xz",
}

func GetAllUsers() ([]string, error) {

	var users []string
	files, err := os.ReadDir(config.ConfigEnv.EmailDBPath)

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

	emails, total, err := SearchEmails("a", "matchphrase", 0)
	if emails != nil || total != 0 {
		return errors.New("Email database is already indexed.")
	}

	err = FetchDecompress()
	if err != nil {
		return err
	}
	log.Println("* Starting index process...")
	users, err := GetAllUsers()

	if err != nil {
		return err
	}

	limiter := make(chan int, config.ConfigEnv.EmailIndexerRoutinesNum)
	var wg sync.WaitGroup

	for _, user := range users {
		limiter <- 1
		wg.Add(1)
		go IndexEmailsByUser(user, &wg, limiter)
	}

	wg.Wait()
	log.Println("* Index process has finished successfully.")
	log.Println("* Cleaning...")
	os.RemoveAll(fmt.Sprintf("%s/", config.ConfigEnv.EmailDatabaseDir))
	os.Remove(config.ConfigEnv.EmailDatabaseFile)
	log.Println("* Done.")
	return nil
}

func IndexEmailsByUser(user string, wg *sync.WaitGroup, limiter chan int) {

	defer wg.Done()
	log.Println("+ Start Process for user: ", user)
	emails, err := GetEmailsByUser(user)

	if err != nil {
		log.Println("Error extracting emails from user: ", user)
		return
	}

	var requestBody models.BulkV2Request

	requestBody.Index = config.ConfigEnv.EmailIndexName
	requestBody.Records = emails

	err = client.IndexDataByBulk(requestBody)

	if err != nil {
		log.Println("There was an error indexing data: ", err)
		return
	}
	log.Println("- End Process for user: ", user)

	<-limiter

}

func GetEmailsByUser(user string) ([]models.Email, error) {

	var emails []models.Email

	userPath := fmt.Sprintf("%s/%s", config.ConfigEnv.EmailDBPath, user)

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
				log.Printf("There was an error processing this email file: %s | Error: %s", path, err)
			} else {
				*emails = append(*emails, *email)
			}

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

	arr := strings.SplitN(string(file), config.ConfigEnv.EmailContentDelimiter, 2)

	if len(arr) != 2 {
		log.Println("There was an error processing this email file: ", path)
		return nil, err
	}

	email.Content = arr[1]

	emailDetails := strings.Split(arr[0], config.ConfigEnv.EmailDetailsDelimiter)

	for i := 0; i < len(emailDetails); i++ {
		detail := strings.SplitN(emailDetails[i], ": ", 2)
		switch detail[0] {
		case "Message-ID":
			email.MessageID = detail[1]
		case "Date":
			date, err := ParseDate(detail[1])
			if err != nil {
				return nil, err
			}
			email.Date = date
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

func SearchEmails(text string, searchType string, page int) ([]models.Email, int, error) {
	body := models.SearchRequest{
		SearchType: searchType,
		Query: models.SearchQuery{
			Term: text,
		},
		SortFiels: []string{
			"-date",
		},
		From:       page * 10,
		MaxResults: 10,
	}

	response, err := client.SearchData(body)

	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	var formatResponse models.SearchResponse
	var emails []models.Email
	err = json.Unmarshal(response, &formatResponse)

	if err != nil {
		log.Println("There was an error trying format response (SearchEmails): ", err)
		return nil, 0, err
	}

	// Process response data (Map to Email model)
	for _, hit := range formatResponse.Hits.Hits {
		var email models.Email

		rawEmail, _ := json.Marshal(hit.Source)

		err := json.Unmarshal(rawEmail, &email)

		if err != nil {
			log.Println("There was an error trying to map an email: ", err)
			continue
		}

		emails = append(emails, email)
	}

	//Return formated data to handler
	return emails, formatResponse.Hits.Total.Value, nil
}

func FetchDecompress() error {

	fileExt := strings.Split(config.ConfigEnv.EmailDatabaseUrl, ".")
	if !slices.Contains(TarSupportedExt, fileExt[len(fileExt)-1]) {
		return errors.New("Unsupported file extension")
	}

	err := FetchEmailDatabase()
	if err != nil {
		return err
	}

	database, err := os.Open(config.ConfigEnv.EmailDatabaseFile)
	if err != nil {
		return err
	}

	err = UnTar(database)
	if err != nil {
		return err
	}

	return nil
}

func FetchEmailDatabase() error {

	log.Println("* Fetching email database...")
	resp, err := http.Get(config.ConfigEnv.EmailDatabaseUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(config.ConfigEnv.EmailDatabaseFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func UnTar(gzipStream io.Reader) error {

	log.Println("* Extracting email database...")
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)
	var header *tar.Header
	for header, err = tarReader.Next(); err == nil; header, err = tarReader.Next() {
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				return fmt.Errorf("ExtractTarGz: Mkdir() failed: %w", err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				return fmt.Errorf("ExtractTarGz: Create() failed: %w", err)
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				// outFile.Close error omitted as Copy error is more interesting at this point
				outFile.Close()
				return fmt.Errorf("ExtractTarGz: Copy() failed: %w", err)
			}
			if err := outFile.Close(); err != nil {
				return fmt.Errorf("ExtractTarGz: Close() failed: %w", err)
			}
		default:
			return fmt.Errorf("ExtractTarGz: uknown type: %b in %s", header.Typeflag, header.Name)
		}
	}
	if err != io.EOF {
		return fmt.Errorf("ExtractTarGz: Next() failed: %w", err)
	}
	return nil
}

func ParseDate(text string) (time.Time, error) {
	output, err := time.Parse(config.ConfigEnv.EmailDateFormat, text)
	if err != nil {
		return time.Time{}, fmt.Errorf("There was a error trying to parse a date: %s", err)
	}
	return output, nil
}

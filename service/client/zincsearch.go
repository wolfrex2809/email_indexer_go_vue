package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/wolfrex2809/email_indexer_go_vue/config"
	"github.com/wolfrex2809/email_indexer_go_vue/models"
)

func IndexDataByBulk(body models.BulkV2Request) error {

	client := &http.Client{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/_bulkv2", config.ConfigEnv.ZincsearchHost), bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.SetBasicAuth(config.ConfigEnv.ZincsearchUsername, config.ConfigEnv.ZincsearchPassword)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		responseError := fmt.Sprintf("There was an error trying to Bulk data in Zincsearch: %s", string(respBody))
		log.Println(responseError)
		return errors.New(responseError)
	}

	return nil
}

func SearchData(body models.SearchRequest) ([]byte, error) {

	client := &http.Client{}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/%s/_search", config.ConfigEnv.ZincsearchHost, config.ConfigEnv.EmailIndexName), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(config.ConfigEnv.ZincsearchUsername, config.ConfigEnv.ZincsearchPassword)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		responseError := fmt.Sprintf("There was an error trying to search data in Zincsearch: %s", string(respBody))
		return nil, errors.New(responseError)
	}

	return respBody, nil
}

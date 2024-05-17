package config

import (
	"log"
	"os"
	"strconv"
)

var ConfigEnv Config

type Config struct {
	Port                    string `mapstructure:"port" json:"port"`
	EmailDBPath             string `mapstructure:"email_db_path" json:"email_db_path"`
	EmailIndexName          string `mapstructure:"email_index_name" json:"email_index_name"`
	EmailContentDelimiter   string `mapstructure:"email_content_delimiter" json:"email_content_delimiter"`
	EmailDetailsDelimiter   string `mapstructure:"email_details_delimiter" json:"email_details_delimiter"`
	EmailIndexerRoutinesNum int    `mapstructure:"email_indexer_routines_num" json:"email_indexer_routines_num"`
	ZincsearchHost          string `mapstructure:"zincsearch_host" json:"zincsearch_host"`
	ZincsearchUsername      string `mapstructure:"zincsearch_username" json:"zincsearch_username"`
	ZincsearchPassword      string `mapstructure:"zincsearch_password" json:"zincsearch_password"`
}

func LoadEnvs() {

	//Port
	if os.Getenv("PORT") != "" {
		ConfigEnv.Port = os.Getenv("PORT")
	} else {
		ConfigEnv.Port = "3000"
	}

	//EmailDBPath
	if os.Getenv("EMAIL_DB_PATH") != "" {
		ConfigEnv.EmailDBPath = os.Getenv("EMAIL_DB_PATH")
	} else {
		ConfigEnv.EmailDBPath = "./enron_mail_20110402/maildir"
	}

	//EmailIndexName
	if os.Getenv("EMAIL_INDEX_NAME") != "" {
		ConfigEnv.EmailIndexName = os.Getenv("EMAIL_INDEX_NAME")
	} else {
		ConfigEnv.EmailIndexName = "emails"
	}

	//EmailContentDelimiter
	if os.Getenv("EMAIL_CONTENT_DELIMITER") != "" {
		ConfigEnv.EmailContentDelimiter = os.Getenv("EMAIL_CONTENT_DELIMITER")
	} else {
		ConfigEnv.EmailContentDelimiter = "\r\n\r\n"
	}

	//EmailDetailsDelimiter
	if os.Getenv("EMAIL_DETAILS_DELIMITER") != "" {
		ConfigEnv.EmailDetailsDelimiter = os.Getenv("EMAIL_DETAILS_DELIMITER")
	} else {
		ConfigEnv.EmailDetailsDelimiter = "\r\n"
	}

	//EmailIndexerRoutinesNum
	if os.Getenv("EMAIL_INDEXER_ROUTINES_NUM") != "" {
		num, err := strconv.Atoi(os.Getenv("EMAIL_INDEXER_ROUTINES_NUM"))
		if err != nil {
			log.Println("There was an error loading 'EMAIL_INDEXER_ROUTINES_NUM' env: ")
		}
		ConfigEnv.EmailIndexerRoutinesNum = num
	} else {
		ConfigEnv.EmailIndexerRoutinesNum = 30
	}

	//ZincsearchHost
	if os.Getenv("ZINCSEARCH_HOST") != "" {
		ConfigEnv.ZincsearchHost = os.Getenv("ZINCSEARCH_HOST")
	} else {
		ConfigEnv.ZincsearchHost = "http://localhost:4080"
	}

	//ZincsearchUsername
	if os.Getenv("ZINCSEARCH_USERNAME") != "" {
		ConfigEnv.ZincsearchUsername = os.Getenv("ZINCSEARCH_USERNAME")
	} else {
		ConfigEnv.ZincsearchUsername = "admin"
	}

	//ZincsearchPassword
	if os.Getenv("ZINCSEARCH_PASSWORD") != "" {
		ConfigEnv.ZincsearchPassword = os.Getenv("ZINCSEARCH_PASSWORD")
	} else {
		ConfigEnv.ZincsearchPassword = "admin"
	}
}

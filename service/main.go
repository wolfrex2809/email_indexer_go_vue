package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/wolfrex2809/email_indexer_go_vue/app"
	"github.com/wolfrex2809/email_indexer_go_vue/config"
	"github.com/wolfrex2809/email_indexer_go_vue/service"
)

func main() {

	f, err := os.Create("profile.pb.gz")
	if err != nil {
		log.Fatal(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	config.LoadEnvs()
	if config.ConfigEnv.ExecutionMode == "indexer" {
		err := service.IndexEmails()
		if err != nil {
			log.Println("There was an error trying to index all emails: ", err)
		}
	}
	if config.ConfigEnv.ExecutionMode == "service" {
		app := app.New()
		err := app.Start(context.TODO())
		if err != nil {
			fmt.Println("Failed to start app: ", err)
		}
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/wolfrex2809/email_indexer_go_vue/app"
	"github.com/wolfrex2809/email_indexer_go_vue/config"
)

func main() {

	config.LoadEnvs()
	app := app.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}
}

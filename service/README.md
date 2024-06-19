<h1 align="center">
  <img alt="logo-zincsearch" src="https://zincsearch-docs.zinc.dev/images/zinc.png" width="100px"/>
  <img alt="logo-golang" src="https://go.dev/images/go-logo-blue.svg" width="150px"/>
  <br/>
  Email Indexer Service
</h1>

Service developed to download, extract and index an email database. This service also works as API that expose an endpoint to make search for the indexed emails.

## 📄 Requirements
* Golang 1.22 [→link](https://go.dev/dl/)
* Zincsearch Service [→link](https://zincsearch-docs.zinc.dev/installation/)

## 🌐 Environment Variables
| Variable                   | Description                                                                                                                                      | Default Value                                                  |
|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------|
| PORT                       | Defines what port the service will start listening to                                                                                            | 3000                                                           |
| EMAIL_DATABASE_URL         | Defines the URL used to download the email database (Only tar files are supported)                                                               | http://download.srv.cs.cmu.edu//~enron/enron_mail_20110402.tgz |
| EMAIL_INDEX_NAME           | Defines the index name where the data will be stored                                                                                             | emails                                                         |
| EMAIL_CONTENT_DELIMITER    | Defines what element is between the details and the content of an email file                                                                     | \\r\\n\\r\\n                                                   |
| EMAIL_DETAILS_DELIMITER    | Defines what element is between every detail of an email                                                                                         | \\r\\n                                                         |
| EMAIL_INDEXER_ROUTINES_NUM | Defines the max number of routines will run at same time during the index process (used to reduce consumed resources or improve the index speed) | 40                                                             |
| EMAIL_DATE_FORMAT          | Defines the format used to parse the emails date                                                                                                 | Mon, 2 Jan 2006 15:04:05 -0700 (MST)                           |
| EXECUTION_MODE             | Defines the way the program will be executed (there are only two modes `indexer` and `service`)                                                  | service                                                        |
| ZINCSEARCH_HOST            | Defines the URL where the Zincsearch service is listening to                                                                                     | http://localhost:4080                                          |
| ZINCSEARCH_USERNAME        | Defines the admin username used to authenticate in Zincsearch service                                                                            | admin                                                          |
| ZINCSEARCH_PASSWORD        | Defines the admin password used to authenticate in Zincsearch service                                                                            | admin                                                          |

>**ⓘ NOTE**: If you set the `EXECUTION_MODE` as `indexer`, the program will execute the proccess to download, extract and index the emails database, Otherwise, if you dont set it, the program will start in `service` mode.

## 🧩 Development
Firstly, you must fetch the `dependencies` by using the following command:
```bash
go mod tidy
```

Later, you run the program use the following command:
```bash
go run main.go
```
>**ⓘ NOTE**: If you want to especify the `EXECUTION_MODE`, you can pass the execution mode by this way.
```bash
EXECUTION_MODE=indexer go run main.go
```


## 🏗️ Build
There are 2 way to build this application.

### Go Cli
Execute the following command (Changing the "{binary_name} "):
```bash
go build -o {binary_name} main.go
```

### Docker
You can build a docker image by using the following command (Changing the "{image_name}"):
```bash
docker build -t {image_name} .
```
>**ⓘ NOTE**: You must have docker installed in your environment.


 😊 Made with ❤️!
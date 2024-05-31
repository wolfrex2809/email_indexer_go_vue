<h1 align="center">
  <img alt="logo-zincsearch" src="https://zincsearch-docs.zinc.dev/images/zinc.png" width="100px"/>
  <img alt="logo-golang" src="https://go.dev/images/go-logo-blue.svg" width="150px"/>
  <img alt="logo-vue" src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/Vue.js_Logo_2.svg/512px-Vue.js_Logo_2.svg.png?20170919082558" width="100px"/>
  <img alt="logo-typescript" src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4c/Typescript_logo_2020.svg/512px-Typescript_logo_2020.svg.png" width="100px"/>
  <br/>
  Email Indexer Go Vue
</h1>
<p align="center">Monorepo application developed to download, extract, index and show an email database. Using <a href="https://zincsearch-docs.zinc.dev/">Zincsearch</a> as search engine and data storage, <a href="https://go.dev/">Go</a> as Backend language and <a href="https://vuejs.org/">Vue</a>+<a href="https://www.typescriptlang.org/">TS</a> as Frontend technology.</p>


## 📜 Table of Projects

* [Client](client/)
* [Service](service/)

## 🖥️ Usage
To run this monorepo there was performed a `Docker Compose` file to run it by using a single command.
```bash
docker compose up
```
This command will build all 3 images and run their very containers with a basic configuration by default.

## ⚙️ Configuration
There are several configurations can be set to customize this application, this configurations can be set by using the following environment variables:

| Variable                   | Description                                                                                                                                      | Default Value                                                                                                                    |
| -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------- |
| ZINC_FIRST_ADMIN_USER      | Defines Zincsearch admin username (used by the Service and Zincsearch UI)                                                                        | admin                                                                                                                            |
| ZINC_FIRST_ADMIN_PASSWORD  | Defines Zincsearch admin password (used by the Service and Zincsearch UI)                                                                        | admin                                                                                                                            |
| EMAIL_INDEX_NAME           | Defines the index name where the data will be stored                                                                                             | emails                                                                                                                           |
| EMAIL_CONTENT_DELIMITER    | Defines what element is between the details and the content of an email file                                                                     | \\r\\n\\r\\n                                                                                                                     |
| EMAIL_DETAILS_DELIMITER    | Defines what element is between every detail of an email file                                                                                    | \\r\\n                                                                                                                           |
| EMAIL_INDEXER_ROUTINES_NUM | Defines the max number of routines will run at same time during the index process (used to reduce consumed resources of improve the index speed) | 40                                                                                                                               |
| EMAIL_DATABASE_URL         | Defines the URL where the database is going to be downloaded (Only tar files are supported)                                                      | [http://download.srv.cs.cmu.edu//~enron/enron_mail_20110402.tgz](http://download.srv.cs.cmu.edu//~enron/enron_mail_20110402.tgz) |


## 📱 Contact

* 📧 **Email**: [Adrian V.](mailto:theblacksstrike@gmail.com)
* 🧳 **Linkedin**: [Adrian V.](https://www.linkedin.com/in/adrian-velasquez-lopez-261570135/)

 😊 Made with ❤️!
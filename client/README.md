<h1 align="center">
  <img alt="logo-vue" src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/95/Vue.js_Logo_2.svg/512px-Vue.js_Logo_2.svg.png?20170919082558" width="100px"/>
  <img alt="logo-typescript" src="https://upload.wikimedia.org/wikipedia/commons/thumb/4/4c/Typescript_logo_2020.svg/512px-Typescript_logo_2020.svg.png" width="100px"/>
  <br/>
  Email Searcher WebApp
</h1>

Web application developed to search and show emails fetched from a Backend application powered by a search engine.

<h1 align="center">
  <img alt="logo-vue" src="https://github.com/wolfrex2809/email_indexer_go_vue/assets/82114869/8a7fc9cf-3323-45ff-ad56-2a8c383c849b" width="500px"/>
  <img alt="logo-vue" src="https://github.com/wolfrex2809/email_indexer_go_vue/assets/82114869/e406598f-3e61-4fd5-a280-b4530bcdb981" width="150px"/>
  <img alt="logo-vue" src="https://github.com/wolfrex2809/email_indexer_go_vue/assets/82114869/729553b9-468f-47ca-abfe-2bfc69f938a5" width="150px"/>
</h1>

## 📄 Requirements
* Node 22 [→link](https://nodejs.org/en/)

## 🌐 Environment Variables
| Variable     | Description                                              | Default Value         |
|--------------|----------------------------------------------------------|-----------------------|
| SERVICE_HOST | Defines the URL of the API used to fetch the emails data | http://localhost:3000 |


## 🧩 Development
Firstly, you must fetch the `dependencies` by using the following command:
```bash
npm install
```

Later, you must serve locally by using the following command:
```bash
npm run dev
```

## 🏗️ Build
There are 2 way to build this application.

### Node Cli
Execute the following command:
```bash
npm run build
```
>**ⓘ NOTE**: The output of this command can be found at `/dist` directory. 

### Docker
You can build a docker image by using the following command (Changing the "{service_url}" and the "{image_name}"):
```bash
docker build --build-arg="SERVICE_HOST={service_url}" -t {image_name} .
```
>**ⓘ NOTE**: You must have docker installed in your environment. 


## ➕ Additional Information
 To avoid "CORS Errors" during the `Preflight Request`, there was configured a `proxy` across "Vite Configuration", but this proxy only works when you run the development server. To keep this configuration, there was  also configured a `proxy` across "Nginx Configuration". In both cases you must to pass the `SERVICE_HOST` environment variable.


 😊 Made with ❤️!

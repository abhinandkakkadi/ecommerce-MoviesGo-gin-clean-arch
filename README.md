
# MoviesGo E-Commerce App in Golang with Gin framework

This is an E-commerce web API which is built using Go programming language and gin as fraemwork. I have followed clean architecture design pattern to build this project in order to implements decoupling and seperation of concerns.






## Project Overview

It's an E commerce website which sell Movie CD of different genres and formats( 4k, BLURAY ). It have all basic functionalites of an E-commerce website along with some advanced features like multiple offer management. 


## Tech Stack

**Server:** Go, Gin


## Run Locally

Clone the project

```bash
  git clone git@github.com:abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch.git
```

Go to the project directory

```bash
  cd ecommerce-MoviesGo-gin-clean-arch
```

Install dependencies

```bash
  make deps
  
  go mod tidy
```

Start the server

```bash
  make run
```


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_HOST`=`your database host name`

`DB_NAME`=`your database name`

`DB_USER`=`your database username`

`DB_PORT`=`your database port number`

`DB_PASSWORD`=`your database owner password`

`DB_AUTHTOKEN`=`Twilio Authentication token`

`DB_ACCOUNTSID`=`Twilio account SID`

`DB_SERVICESID`=`Twilio message service SID`
## 🔗 Reach Me

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/abhinand-k-r-300036129/)


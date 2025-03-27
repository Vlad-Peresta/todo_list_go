# TODO List API

API for organizing and managing tasks written on Gin

## Installation
Golang should be installed

#### Download the code
```angular2html
git clone https://github.com/Vlad-Peresta/todo_list_go.git
cd todo_list_go
```

#### Set Up for Unix, macOS
```angular2html
go mod download
```

#### Set Up environment variables and Database for Unix, macOS
```angular2html
export POSTGRES_USER=<your db username>
export POSTGRES_PASSWORD=<your db user password>
export DB_HOST=<your db hostname>
export POSTGRES_DB=<your db name>
export DB_PORT_INTERNAL=<your internal port>
export DB_PORT=<your port>
export JWT_SECRET_KEY=<your secret key>
```

#### Start the app
```angular2html
go run main.go
```

#### Run with Docker
Docker should be installed
```angular2html
docker-compose build
docker-compose up
```

## Features

* JWT authenticated
* Documentation is located at /api/v1/docs/swagger/
* Managing Todos
* Filtering Todos

## Getting access
* create user via /api/v1/auth/signup/
* get access token /api/v1/auth/login/

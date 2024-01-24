# go-backend
 This is a backend server built with go, it is a webforum that supports basic CRUD, authentication, and tagging(WIP).

## Note:
This is a sample project!
The frontend of the project can be found at : [https://github.com/leroytan/react-frontend](https://github.com/leroytan/react-frontend).\
The deployment of the project can be found at [https://www.forunme.leroymx.com/](https://www.forunme.leroymx.com/).

## Requirements
Download go from https://go.dev/dl/


## To run the server:
### Step 1:
In your terminal, run the following commands:
```
go build
./go-backend.exe
```

### Step 2:
Create a sql database.
E.g. postgres/sqlite/mysql

### Step 3:
Create a .env file \
This file will handle sensitive information \
Since this is a sample project, I have included a .env.template file \
Change the .env.template file to a .env file \
Replace the fields with your own database, port, frontend and secret key (for jwt authorization).







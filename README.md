# Book Management System API

A simple CRUD API for managing books in a bookstore built using:

- [Go](https://go.dev/)
- [GorillaMux](https://github.com/gorilla/mux)
- [GORM](https://gorm.io/)
- [MySQL](https://www.mysql.com/)

## Environment Variables

This project makes use of the [godotenv](github.com/joho/godotenv) package to store environment variables. To run this project on your local machine, you will need to add your MySQL database details as environment variables. This [article](https://dev.to/schadokar/use-environment-variable-in-your-next-golang-project-2o6c) explains how to use the package.
> Your env variables should have the names `DBUSERNAME`, `DBPASSWORD` and `DBNAME` respectively.

## Run API locally

- Clone Repo

    ```bash
    git clone https://github.com/Marvellous-Chimaraoke/book-management-system.git
    ```

- Make sure to have at least Go version 1.18 installed on your device
- Open the code base directory in terminal
- Run program:

    ```go
    go build -o bms cmd/main.go
    ./bms
    ```

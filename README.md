# Bookstore REST API using Gin and Gorm

Based on the [article](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/) by Rahman Fadhil.

- install dependencies

  ```go
  $ go mod tidy
  ```

- Create a **`.env`** file with vars from **`.env.example`**

- Run the server:

  ```go
  $ go run main.go
  ```

- Run the server with a clean DB

  ```go
  $ go run main.go --clean
  ```

## API docs

Add GOPATH to PATH

- Add GOPATH to PATH

  ```shell
  $ export GOPATH=$(go env GOPATH)/bin
  ```

- Generate docs

  ```shell
  $ swag init --parseDependency --parseInternal
  ```

- Run the server and visit

  ```shell
  http://localhost:8080/swagger/index.html
  ```

## ToDo

- ✅ CRUD
- ✅ Logging
- ✅ Auth middleware JWT
- ✅ Validation
- ✅ API docs (swagger)
- ⬜ Migrations
- ✅ pagination
- ⬜ filter
- ⬜ rate limiting
- ⬜ handle timeouts
- ⬜ background task: email
- ⬜ password reset
- ⬜ file upload
- ⬜ requests.Session: avoid excessive connections, avoid slowdown
- ⬜ dockerize

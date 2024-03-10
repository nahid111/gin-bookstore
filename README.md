# Bookstore REST API using Gin and Gorm

Based on the [article](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/) by Rahman Fadhil.

- install dependencies

  ```go
  $ go mod download
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

## ToDo

- ✅ Auth middleware JWT
- ✅ Validation
- ⬜ Migrations
- ⬜ logging
- ⬜ pagination
- ⬜ filter
- ⬜ rate limiting
- ⬜ handle timeouts
- ⬜ background task: email
- ⬜ password reset
- ⬜ file upload
- ⬜ requests.Session: avoid excessive connections, avoid slowdown
- ⬜ dockerize

# GoToStudy
Golang learning labs

Credits: [@wesleywillians](https://github.com/wesleywillians)

---
###Dependencies:
<ol>
<li>Go: [https://go.dev/dl](https://go.dev/dl)</li>
</ol>

---
####Go Mods:
<ol>
    <li>github.com/mattn/go-sqlite3</li>
    <li>github.com/stretchr/testify/assert</li>
    <li>github.com/golang/mock/gomock</li>
</ol>

To install Go mods, run the following command in the project directory:

`$  go get <<mod_path>>`

Example:

`$ go get github.com/maragudk/migrate`

###Testing

To run the implemented tests, run the following command in the project directory:

`$ go test ./...`


###Run the App

In terminal, run the command from the app root directory:

`$ go run cmd/main.go`

or just go run main.go pointing it directory whereby located.
# making a basic api in golang

initialize the project

`go mod init first-api`

install dependency:

`go mod tidy`

curl:
get books:

`curl localhost:8080/books`

`curl localhost:8080/books/2`

insert books (i created a file for keep easy update with curl)

```
curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
```

patch 

`curl localhost:8080/checkout?id=3 --request "PATCH"`

`curl localhost:8080/return?id=3 --request "PATCH"`
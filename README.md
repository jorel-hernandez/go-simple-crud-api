# Init go module at folder go-simple-crud-api
```go
go mod init go-simple-crud-api
```

# Run main.go 
```go
go run main.go
```
# Test Resources
GET homepage: 
```curl
curl --location --request GET 'http://localhost:10000'

# Respose:
Welcome to the HomePage!
```

GET all articles
```curl
curl --location --request GET 'http://localhost:10000/articles'

# Respose:
[
    {"Id":"1","title":"Hello","desc":"Article Description","content":"Article Content"},
    {"Id":"2","title":"Hello 2","desc":"Article Description","content":"Article Content"}
]
```

POST new article
```curl
curl --location --request POST 'http://localhost:10000/article' \
--header 'Content-Type: application/json' \
--data '{
    "Id": "3", 
    "Title": "Newly Created Post", 
    "desc": "The description for my new post", 
    "content": "my articles content" 
}'

# Respose:
{"Id":"3","title":"Newly Created Post","desc":"The description for my new post","content":"my articles content"}
```

PUT article
```curl
curl --location --request PUT 'http://localhost:10000/article/2' \
--header 'Content-Type: application/json' \
--data '{
    "Id": "2", 
    "Title": "Updated PUT", 
    "desc": "The description for my new put", 
    "content": "my articles content" 
}'

# Respose:
{"Id":"2","title":"Updated PUT","desc":"The description for my new put","content":"my articles content"}
```

GET article
```curl
curl --location --request GET 'http://localhost:10000/article/2'

# Respose:
{"Id":"2","title":"Updated PUT","desc":"The description for my new put","content":"my articles content"}
```

DELETE article
```curl
curl --location --request DELETE 'http://localhost:10000/article/3'

# Respose:
None
```

# Source
- https://tutorialedge.net/golang/creating-restful-api-with-golang/
- https://tutorialedge.net/software-eng/what-is-a-rest-api/
- https://golang.org/pkg/encoding/json/
- https://github.com/gorilla/mux

# go-sample-api
# version 1.0.0

### set up
```
export GO111MODULE=on 
cp .env.dev .env
go run main.go
```

### get request
```
curl http://localhost:8000/
```

### post request
```
curl -X POST -H "Content-Type: application/json" -d '{"name":"takahashi", "email":"sample@test.com"}' localhost:8000/v1/users
```
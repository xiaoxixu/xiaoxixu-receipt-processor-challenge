# Receipt Processor

A webservice that fulfils the API specified in [receipt-processor-challenge](https://github.com/fetch-rewards/receipt-processor-challenge/tree/main).

---
## Run Application
Make sure you have go installed. To check:
```
go version
```
Navigate to the project and run:
```
go run main.go
```
---
## Testing
### Testing with an API platfrom
Download a platform such as [Postman](https://www.postman.com/).
Send an example Post request with URL and any [example](./examples/) as the payload:
```
http://localhost:8080/receipts/process
```

Send an example Get request with URL and an ID you get from a Post request:
```
http://localhost:8080/receipts/b5ab6339-1cb2-4047-9566-848e12901f2f/points
```

### Unit tests
Navigate to the project and run:
```
go test ./...
```

Or navigate to the `svc` folder from project root by:
```
cd internal/svc
```
and run:
```
go test
```
---

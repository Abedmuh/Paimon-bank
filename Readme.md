### What this app do?

This project was a backend restfull api with golang it receive request and respons it with JSON. the app was replica for banking apps.

### command apps

yout can use makefile as well but here the manual, all of these was on run on dev mode. for prod watch your .env

```
//build
go build -o main

//migrate up
migrate -path db/migrations -database "postgresql://abdillah:pass@localhost:5432/paimonbank?sslmode=disable" up

//migrate down
migrate -path db/migrations -database "postgresql://abdillah:pass@localhost:5432/paimonbank?sslmode=disable" down

testing on k6 with windows on powershell
$env:BASE_URL = 'http://localhost:8080'
k6 run --vus 1 --iterations 1 script.js
```

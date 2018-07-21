## Build the app
```
go build main.go
```
## Run the app
Fill in the Environment Variables with Postgres Connection details
```
export AWS_RDS_HOSTNAME=***
export AWS_RDS_USERNAME=***
export AWS_RDS_PASSWORD=****
export DATABASE_NAME=****
./main
```

or

```
go run main.go
```

* TODO's
  * Using unicreds or credstash for credentials to be encrypted
  * Add unit testing

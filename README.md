# How to run

## Clone this repository
```
use this following command
git clone https://github.com/Amzar123/lms-golang.git
```

## Run apps using docker 
```
docker-compose -f docker-compose.yml up
```

## Standard running apps
```
npm run main.go
```

## Run unit testing
```
go test -v -coverprofile=cover.out && go tool cover -html=cover.out
```
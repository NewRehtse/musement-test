# Musement | Backend tech homework

This is a test application which gets the list of the cities from Musement's API for each city gets the forecast for 
the next 2 days using http://api.weatherapi.com and print to STDOUT 
"Processed city [city name] | [weather today] - [weather tomorrow]"

*Example:*
> Processed city Milan | Heavy rain - Partly cloudy
>
> Processed city Rome | Sunny - Sunny

## Author

- Esther Ibáñez González <esther.ibezgonzalez@gmail.com> 

## Installation

* Clone project

### Run application with Docker

```bash
docker build --tag musement .
docker run --rm musement:latest
```

###  Run application in local machine

* Install golang v1.13 or above (I've tried it with v1.13, v1.14 and v1.17, it has anything too new in recent versions).
Instructions: https://golang.org/doc/install
* Install vendors
```bash
go mod download
```
* Run application without building it
```bash
go run main.go
```
* Or build executable
```bash
go build -o ./mus
``` 
* Run it!
```bash
./mus
``` 

### Run tests

```bash
go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
```

### Run Cover test

1. Generate cover file:
```bash
go test ./.. -coverprofile=coverage.out
```
2. Inspect cover file with go tool cover
```bash
go tool cover -html=coverage.out
```

### Run clean Code goimport

**Goimports** updates Go import lines, adding missing ones and removing unreferences ones, 
as well as formats the code to follow one style code.
**Golint** as the name says is a linter for Go.

1. Install tools:
```
go get -u golang.org/x/lint/golint
go get golang.org/x/tools/cmd/goimports
```
2. And run tools:
```
goimports -w -l *.go && golint *.go
```
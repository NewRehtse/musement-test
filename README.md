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

## Running tests

```bash
$ go test -v ./...
```

### Clean Code goimport

``
goimports -w -l *.go && golint *.go
``
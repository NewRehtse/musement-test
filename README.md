# Musement | Backend tech homework

## Author

- Esther Ibáñez González <esther.ibezgonzalez@gmail.com>

## Index

* [Step 1](#step-1)
* [Step 2](#step-2)

## Step 1

This is a test application which gets the list of the cities from Musement's API for each city gets the forecast for 
the next 2 days using http://api.weatherapi.com and print to STDOUT 
"Processed city [city name] | [weather today] - [weather tomorrow]"

*Example:*
> Processed city Milan | Heavy rain - Partly cloudy
>
> Processed city Rome | Sunny - Sunny


### Installation

* Clone project

#### Run application with Docker

```bash
docker build --tag musement .
docker run --rm musement:latest
```

####  Run application in local machine

* Install golang v1.13 or above (I've tried it with v1.13, v1.14 and v1.17, it has anything too new in recent versions).
Instructions: https://golang.org/doc/install
* Install vendors
```bash
make deps
```
* Run application without building it
```bash
make build
./bin/app
``` 

#### Run tests

```bash
make tests
```

You can find coverage file in build/coverage.html

#### Run benchmarks

```bash
make bench
```

#### Run clean Code goimport

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

## Step 2

I propose three endpoints:

* GET /api/v3/cities/{cityId}/forecasts to get all forecasts for a city
* GET /api/v3/cities/{cityId}/forecasts/{forecast_date} to get forecast for a specific date as today, tomorrow or 2021-09-22
* PUT /api/v3/cities/{cityId}/forecasts to set forecasts for a specific city

I have update OpenApi specs for Musement's API here: https://app.swaggerhub.com/apis/n2628/MusmentApiBackendTechHomework/1.0.0

You can get it as well at [api_specs.yaml](api_specs.yaml)

This is exactly what I added:

```yaml
 /cities/{cityId}/forecasts:
    get:
      tags:
        - City
      summary: Get city forecasts by unique identifier
      operationId: GetCitiesCityIdForecasts
      parameters:
        - $ref: '#/components/parameters/X-Musement-Version'
        - $ref: '#/components/parameters/Accept-Language'
        - $ref: '#/components/parameters/cityId'
      responses:
        '200':
          description: Returned when successful
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Forecast'
        '404':
          description: Returned when resource is not found
        '503':
          description: Returned when the service is unavailable
    put:
      tags:
        - city
      summary: Appends or updates forecast for a city at a specific date
      operationId: PutForecastsCity
      parameters:
        - $ref: '#/components/parameters/X-Musement-Version'
        - $ref: '#/components/parameters/Accept-Language'
        - $ref: '#/components/parameters/cityId'
      requestBody:
        description: Forecast put request
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Forecast'
      responses:
        '200':
          description: Returned when successful
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Forecast'
        '400':
          description: Returned if sent data contains errors
        '404':
          description: Returned when resource is not found
        '503':
          description: Returned when the service is unavailable
  /cities/{cityId}/forecasts/{forecast_date}:
    get:
      tags:
        - City
      summary: Get city forecasts by unique identifier and specific date
      operationId: GetCitiesCityIdForecastsByDate
      parameters:
        - $ref: '#/components/parameters/X-Musement-Version'
        - $ref: '#/components/parameters/Accept-Language'
        - $ref: '#/components/parameters/cityId'
        - name: forecast_date
          in: path
          description: 'Forecast date | If not specified set to today | Use format: YYYY-MM-DD'
          required: true
          schema:
            type: string
            enum:
              - today
              - tomorrow
              - date
      responses:
        '200':
          description: Returned when successful
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Forecast'
        '404':
          description: Returned when resource is not found
        '405':
          description: Returned when there is no forecast for that date
        '503':
          description: Returned when the service is unavailable

    Forecast:
      properties:
        date:
          type: string
          format: date-time
        condition:
          type: string
```
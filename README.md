
# Kosha Duo Connector

Securing your entire organization has never been easier. Duo’s multi-factor authentication (MFA), single sign-on (SSO), remote access and access control products deploy fast in any environment. 
Duo security help keep companies safer than ever before with minimal downtime and optimized productivity.

The connector APIs allow you to perform 'RESTful' operations such as reading, modifying, adding or deleting data. The APIs also support Cross-Origin Resource Sharing (CORS).

![Duo Security](images/duosecurity.jpg)

## Build

To build the project binary, run
```
    go build -o main .

```

## Run locally

To run the project, simply provide env variables to supply the API key and Duo Security domain name.


```bash
go build -o main .
API_KEY=<API_KEY> DOMAIN_NAME=<DOMAIN_NAME> ./main
```

This will start a worker and expose the API on port `8012` on the host machine

Swagger docs is available at `https://localhost:8012/docs`

## Generating Swagger Documentation

To generate `swagger.json` and `swagger.yaml` files based on the API documentation, simple run -

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g main.go --parseDependency --parseInternal
```

To generate OpenAPISpec version 3 from Swagger 2.0 specification, run -

```bash
npm i api-spec-converter
npx api-spec-converter --from=swagger_2 --to=openapi_3 --syntax=json ./docs/swagger.json > openapi.json
```

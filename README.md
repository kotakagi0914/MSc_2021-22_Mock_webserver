# MSc_2021-22_Mock_webserver
## Abstruct
This is a repository for MSc Cyber Security individual project at City, University of London.
The code in the repository is used for a mock webserver that has login form and is protected by reCAPTCHA v3.

## How to run the server
```
## Create `.secret` from `.secret.template`
$ cp .secret.template .secret

## Set `site-key` and secret-key to `.secret`
$ vim .secret

## Run HTTP server
$ go run cmd/mock-server/main.go

## Access the localhost:8000
```

## How to run the server through docker image
```
## Build docker image
$ docker build -t <image name> .

## Run the docker container with the image
$ docker run -d -p <local port>:8000 <image name>

## Access the localhost:`<local port>`
```

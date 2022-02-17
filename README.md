## go-url-shortener
Simple URL shortener service that will accept a URL as an argument over a REST API and return a shortened URL as a result

## Installation 
#### 1. Build from source.
Prerequisite: Latest go, git installed
```
export GOPATH=<dir_to_set_as_GOPATH>
mkdir -p $GOPATH/src/github.com/vitt-bagal
cd $GOPATH/src/github.com/vitt-bagal
git clone https://github.com/vitt-bagal/go-url-shortener.git
cd go-url-shortener.git
go mod download && go build -o /url-shortner-service
```
#### 2. docker
To use shortener-url service, run below docker command to run docker container:
```
docker run -d -p 9090:9090 bagalvitthal/go-url-shortener:latest
``` 
## Usage
Start service with below command
```
/url-shortner-service &
```
shortener-url service will run on `http://localhost:9090/`

To create short-url for long-url, do post API request - `http://localhost:9090/create-short-url` using below json body
```
{
"long_url": "<LONG_URL>"
}
```
Redirect to long URL using `http://localhost:9090/<Created_Short_URL_ID>`


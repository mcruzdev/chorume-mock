# ChorumeMock

ChorumeMock converts [OpenAPI Specification V3](https://swagger.io/specification/) to Wiremock [Bulk Stubs](https://wiremock.org/docs/stubbing/#bulk-importing-stubs).

> ⚠️ **NOTE** Project status: `development`

<p align="center">
  <img src="img.png" width="200" height="200" style="display: block; margin: 0 auto;" />
</p>


## Getting Started

### Installing

Clone this repo:

```shell
git clone git@github.com:mcruzdev/chorume-mock.git
```

Build the `chorume-cli` project:

```shell
go build 
```


## How-to guides

### How-to generate a Wiremock definition from OpenAPI specification

```
chorumemock generate -oapi=openapi.yaml
```

After, you can execute the wiremock using docker:

```shell
docker run -it --rm -p 8080:8080 -v "$PWD:/home/wiremock" --name wiremock wiremock/wiremock --verbose
```

Make a HTTP request

```shell
curl http://localhost:8080/ping
```

The output should looks something like it:

```shell
2023-09-07 18:41:35.197 Request received:
172.17.0.1 - GET /ping

Host: [localhost:8080]
Connection: [keep-alive]
Cache-Control: [max-age=0]
sec-ch-ua: ["Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"]
sec-ch-ua-mobile: [?0]
sec-ch-ua-platform: ["macOS"]
Upgrade-Insecure-Requests: [1]
User-Agent: [Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36]
Accept: [text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7]
Sec-Fetch-Site: [none]
Sec-Fetch-Mode: [navigate]
Sec-Fetch-User: [?1]
Sec-Fetch-Dest: [document]
Accept-Encoding: [gzip, deflate, br]
Accept-Language: [pt-BR,pt;q=0.9,en-US;q=0.8,en;q=0.7]



Matched response definition:
{
  "status" : 200
}

Response:
HTTP/1.1 200
Matched-Stub-Id: [f15588fb-cfa9-4a3b-a689-8f8199f9b0ff]
```

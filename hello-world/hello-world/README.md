# Simple example


## How to run?

```go
go run some.go
```

## Compile and create executables

```go
go build some.go
```

Run it!

```
./main
```

## here is an example curl

```bash
curl -X POST localhost:5555 -d "http://onet.pl"
```

## here is and example request for postman

```json
{
	"variables": [],
	"info": {
		"name": "GoLang",
		"_postman_id": "aceaf0a4-1bb1-692c-e7cd-6165e0020ef3",
		"description": "",
		"schema": "https://schema.getpostman.com/json/collection/v2.0.0/collection.json"
	},
	"item": [
		{
			"name": "localhost:5555",
			"request": {
				"url": "localhost:5555",
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "http://google.pl\nhttp://onet.pl\nhttp://facebook.com\nhttp://fakt.pl"
				},
				"description": "Call urls and save the content to file"
			},
			"response": []
		}
	]
}
```

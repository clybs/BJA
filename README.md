# BJA
Sugar high simple API

### Installation

BJA requires latest [Golang](https://golang.org/doc/install) to run.

### Run the app
Login using default username and password to get token.
Token expires after 10 mins.

```sh
$ curl -X POST \
    http://localhost:3000/login \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/x-www-form-urlencoded' \
    -d 'username=clybs&password=123456'
```

Create an ice cream detail. Use token from login.
```sh
$ curl -X POST \
    http://localhost:3000/icecreams \
    -H 'Authorization: Bearer <token>' \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json' 
```

Read details of an ice cream. Use token from login.
```sh
$ curl -X GET \
    http://localhost:3000/icecreams/<ice cream id> \
    -H 'Authorization: Bearer <token>' \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json'
```

Update details of an ice cream. Use token from login.
```sh
$ curl -X PUT \
    http://localhost:3000/icecreams/<ice cream id> \
    -H 'Authorization: Bearer <token>' \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json' \
    -d '{
    "name": "New Name",
    "description": "New description"
  }'
```

List all ice creams. Use token from login.
```sh
$ curl -X GET \
    http://localhost:3000/icecreams \
    -H 'Authorization: Bearer <token>' \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json'
```

Delete an ice cream. Use token from login.
```sh
$ curl -X DELETE \
    http://localhost:3000/icecreams/<ice cream id> \
    -H 'Authorization: Bearer <token>' \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json' 
```

### Tests
Run the tests:
```sh
$ go test -cover ./... -v
```
### Documentation
Run the docs:
```sh
$ godoc -http=":6060"
```
Then visit: [http://localhost:6060/pkg/github.com/clybs/](http://localhost:6060/pkg/github.com/clybs/)
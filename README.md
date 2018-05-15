# BJA
Sugar high simple API

### Installation

BJA requires latest [Golang](https://golang.org/doc/install) to run.

### Build and run the app (Optional)
Go to project folder and type:

```sh
$ cd BJA
$ go build
$ ./BJA
```

### Run the app (Not built)
Go to project folder and type:

```sh
$ cd BJA
$ go run main.go
```

### Interact with the app
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
      -H 'Content-Type: application/json' \
      -d '{
      "name": "Everything But The...",
      "image_closed": "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing.png",
      "image_open": "/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/vanilla-toffee-landing-open.png",
      "description": "Vanilla Ice Cream with Fudge-Covered Toffee Pieces",
      "story": "Vanilla What Bar Crunch? We gave this flavor a new name to go with the new toffee bars we’re using as part of our commitment to source Fairtrade Certified and non-GMO ingredients. We love it and know you will too!",
      "sourcing_values": [
        "Non-GMO",
        "Cage-Free Eggs",
        "Fairtrade",
        "Responsibly Sourced Packaging",
        "Caring Dairy"
      ],
      "ingredients": [
        "cream",
        "skim milk",
        "liquid sugar",
        "water",
        "sugar",
        "coconut oil",
        "egg yolks",
        "butter",
        "vanilla extract",
        "almonds",
        "cocoa (processed with alkali)",
        "milk",
        "soy lecithin",
        "cocoa",
        "natural flavor",
        "salt",
        "vegetable oil",
        "guar gum",
        "carrageenan"
      ],
      "allergy_info": "may contain wheat, peanuts and other tree nuts",
      "dietary_certifications": "Kosher",
      "productId": "646"
    }'
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

### Dev Notes
- Decided to go with GoCraft for framework as it is super fast and minimal. It also 
has nested routers, contexts, and middleware.
- Decided to use JWT as it stateless self contained token which has authetication 
information, expire time information, and other user defined claims 
digitally signed. Portable: A single token can be used with multiple backends. Good Performance: It reduces the network round trip time 
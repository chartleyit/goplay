# PART-3 HTTP

this section shows how to make simple HTTP endpoint for handling GET and POST requests

## How-To dev

```bash
docker build --ta4rget dev . t go
docker run -it -p 8080:8080 -v ${PWD}:/work go sh
go version
```

This will put you inside the dev container for working

## How-To run

While inside the dev container execute

```bash
cd videos
go build
./videos
```

You will then be able to access http://localhost:8080 and get back the contents of videos.json

### Update

you make also submit an update to the endpoint by providing JSON contents via POST to http://localhost:8080/update
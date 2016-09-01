# talk-vote

talk-vote is a simple service using the Go-Kit library for microsservices. 

## Endpoints

| Verb | Path | Description |
| ---- | ---- | ----------- |
| GET | /api/v1/talks | provides a list of talks in JSON |
| POST | /api/v1/vote | receives POST'd JSON and increments a talk vote count | 

## Testing

Testing for this service utilizes Go's built in testing facilities as well as a secondary tool called [GoConvey](http://goconvey.co/).  With GoConvey you can write out your tests in a more human readable manner and get instant feeback from it's UI.

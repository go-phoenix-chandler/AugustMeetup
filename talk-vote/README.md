# talk-vote

talk-vote is a simple service using the Go-Kit library for microsservices. 

## Endpoints

Verb | Path | Description
GET | /api/v1/talks | provides a list of talks in JSON
POST | /api/v1/vote | receives POST'd JSON and increments a talk vote count 
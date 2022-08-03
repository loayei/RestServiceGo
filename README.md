### RestServiceGo
## A basic REST API for AutoParts.

# Features implemented:
- GET all parts
- GET a single part by ID
- POST a new part
- PUT an existing part
- DELETE an existing part

POSTMAN used for testing using 
PORT: 6969

## Instructions:
Clone the repo and run `go build` to build the binary.
Run `./AUTO_PARTS` to start the server.
Use POSTMAN to test the API.

## Format:
`
{
    "manufacturer": "BMW",
    "name": "Spark Plug",
    "price": 13.99,
    "part_number": "A344BC"
}
`

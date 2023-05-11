# APIs
## Generate Token
- http://localhost:5000/api/auth for generating token
```json
{
  "email": "jinheehanaaa@gmail.com",
  "password": "supersecurepassword"
}
```
- X-Api-Token

## List a Hotel
- http://localhost:5000/api/v1/hotel for listing hotel

## Book a room
- http://localhost:5000/api/v1/room/645bff84fe030d781f0060df/book
```json
{
  "fromDate": "2023-06-01T00:00:00Z",
  "tillDate": "2023-06-05T00:00:00Z",
  "numPersons": 2
}
```

## List Rooms
- http://localhost:5000/api/v1/room/
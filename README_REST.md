# APIs
## User/Admin Auth (POST)
- http://localhost:5000/api/auth
- X-Api-Token

```json
{
  "email": "jinheehanaaa@gmail.com",
  "password": "supersecurepassword"
}
```
```json
{
  "email": "admin@admin.com",
  "password": "adminpassword123"
}
```

## List a hotel (GET)
- http://localhost:5000/api/v1/hotel

## Book a room (POST)
- http://localhost:5000/api/v1/room/645bff84fe030d781f0060df/book
```json
{
  "fromDate": "2024-06-01T00:00:00Z",
  "tillDate": "2024-06-05T00:00:00Z",
  "numPersons": 2
}
```

## List Rooms (GET)
- http://localhost:5000/api/v1/room/

## Admin booking list (GET)
- http://localhost:5000/api/v1/admin/booking

## User booking list (GET)
- http://localhost:5000/api/v1/booking/64625d36532b9268e41c6637

## Cancel booking (GET)
- http://localhost:5000/api/v1/booking/64625d36532b9268e41c6637/cancel
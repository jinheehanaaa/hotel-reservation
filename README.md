
# CMD
- ./bin/api --listenAddr :7000
- docker run --name mongodb -d mongo:latest -p 27017:27017

# APIs
- http://localhost:5000/api/auth for generating token
```json
{
  "email": "jinheehanaaa@gmail.com",
  "password": "supersecurepassword"
  }
```
- http://localhost:5000/api/v1/hotel for listing hotel
- http://localhost:5000/api/v1/hotel


{
  "numPersons": "2",
  "fromDate": "2023-01-01",
  "tillDate": "2023-01-02"
}



# Features
- Store User in DB
- Store Hotel & Room in DB
- User Authentication using JWT


# Go Packages
- Fiber
- mongodb
- golang-jwt

# Tools
- MongoDB & MongoDB Compass
- Docker

# Review
## MongoDB Code
- bson (D,E,M)


# Resources
- [Official GitHub](https://github.com/fulltimegodev/hotel-reservation)
- [MongoDB Doc](https://mongodb.com/docs/drivers/go/current/quick-start/)
- [GoFiber](https://gofiber.io)

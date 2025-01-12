# DoHabitsAPI

## Description
DoHabitsAPI controls the backend operations for the DoHabitsApp. 

MVP:
- User can create, update and delete up to 5 habits to track daily.

## Features

## Architecture / Design Pattern
MVC Architecture, CRUD Operations and Go Middleware.

## Session
JWT Tokens:
- Access Token - Short-Lived JWT Token
- Refresh Token - Long-Lived JWT Token stored in DB to refresh the user's Access Token. When expires the user must login.

## Database
No-SQL MongoDB

# Collection(s)
habits:
Grows linearly.
```
{
  "_id": {
    "$oid": "677ac8294620315e952dabd7"
  },
  "UserID": {
    "$oid": "677ac7224620315e952dabd6"
  },
  "CreatedAt": {
    "$date": {
      "$numberLong": "1726914600000"
    }
  },
  "Name": "Code everyday",
  "Days": {
    "$numberInt": "30"
  },
  "DaysTarget": {
    "$numberInt": "66"
  },
  "CompletionDates": [
    "2024-12-02",
    "2024-12-02",
    "2024-12-11"
  ]
}
```

user_session:
Stores the user session. Is deleted when the user logs out.
```
{
  "_id": {
    "$oid": "677ac7224620315e952dabd6"
  },
  "RefreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyMUBleGFtcGxlLmNvbSIsImV4cCI6MTczNjc3MjQ3MX0.vfVn044yRxx5kiLfo_VYxzqturusyN2gZofoGPK5hVg",
  "Device": "LSB",
  "IpAddress": "169.254.159.127",
  "CreatedAt": {
    "$date": {
      "$numberLong": "1736686071678"
    }
  }
}
```

users:
Grows linearly.
```
{
  "_id": {
    "$oid": "677ac7224620315e952dabd6"
  },
  "RefreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyMUBleGFtcGxlLmNvbSIsImV4cCI6MTczNjc3MjQ3MX0.vfVn044yRxx5kiLfo_VYxzqturusyN2gZofoGPK5hVg",
  "Device": "LSB",
  "IpAddress": "169.254.159.127",
  "CreatedAt": {
    "$date": {
      "$numberLong": "1736686071678"
    }
  }
}
```

## API Documentation
Link to api.md

## Tech Stack
Golang
MongoDB

## Installation
MongoDB...

Locally...
env vars
go run .

Docker...
https://github.com/LukeASB/HabitsApp/blob/main/Dockerfile

## Dependencies
Required Go Modulesin `go.mod`
MongoDB
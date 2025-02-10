# DoHabitsAPI

## Description
DoHabitsAPI controls the backend operations for the DoHabitsApp. 

## Endpoints
See: https://github.com/LukeASB/HabitsApp/blob/documentation/backend/api.md
### User Endpoints
1. **Register**: `POST /dohabitsapp/v1/register`
2. **Login**: `POST /dohabitsapp/v1/login`
3. **Logout**: `POST /dohabitsapp/v1/logout`
4. **Refresh**: `POST /dohabitsapp/v1/refresh`

### Habit Endpoints
1. **Create Habit**: `POST /dohabitsapp/v1/createhabit`
2. **Retrieve Habit**: `GET /dohabitsapp/v1/retrievehabit`
3. **Retrieve All Habits**: `GET /dohabitsapp/v1/retrievehabits`
4. **Update Habit**: `PUT /dohabitsapp/v1/updatehabit`
5. **Update All Habits**: `PUT /dohabitsapp/v1/updatehabits`
6. **Delete Habit**: `DELETE /dohabitsapp/v1/deletehabit`

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

Has an index TTL on "CreatedAt" to delete the document in MongoDB after 24 Hours.
```
{
  "_id": {
    "$oid": "677ac7224620315e952dabd6"
  },
  "RefreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyMUBleGFtcGxlLmNvbSIsImV4cCI6MTczNjc3MjQ3MX0.vfVn044yRxx5kiLfo_VYxzqturusyN2gZofoGPK5hVg",
  "Device": "Device",
  "IpAddress": "10.10.10.10",
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
    "_id": "679a81a7f0881bc3a7e6aff8",
    "Password": "$2a$10$3vkUrYqkXQoPtvB0NTOkHum7/L8JghFYuqOaNnvTDR3B6UvLXWkqW",
    "FirstName": "TestUser123",
    "LastName": "TestUser123",
    "EmailAddress": "test334@example.com",
    "CreatedAt": "2025-01-29T19:29:43.793+00:00",
    "LastLogin": "2025-01-29T19:30:21.653+00:00"
}
```

## Tech Stack
Golang, MongoDB

# Dependencies
See go.mod

# Build
See: https://github.com/LukeASB/HabitsApp/blob/documentation/README.md

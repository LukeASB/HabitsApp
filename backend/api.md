# DoHabitsApp API Documentation

This document outlines the endpoints for the DoHabitsApp API. 

Each endpoint is prefixed with the API name and version, e.g., `/dohabitsapp/v1`.

## User Endpoints
The following endpoints create or control the user's session state.

### 1. Register
**Endpoint** `POST /dohabitsapp/v1/register`

Registers the user. Validates user data, if successful adds their credentials to the DB Layer.

**Request:**

Request Headers:
| Key            | Value            |
|----------------|------------------|
| Content-Type   | application/json |

Request Body:
| Field         | Type   | Description            | Example              |
|---------------|--------|------------------------|----------------------|
| FirstName     | string | User's first name      | TestUser123          |
| LastName      | string | User's last name       | TestUser123          |
| EmailAddress  | string | User's email address   | test222@example.com  |
| Password      | string | User's password        | !secretPASSWORD123   |

Request Body Example:
```json
{
    "FirstName": "TestUser123",
    "LastName": "TestUser123",
    "EmailAddress": "test222@example.com",
    "Password": "!secretPASSWORD123"
}
```

**Response:**

Response Body:
**Response Body**:
| Field        | Type    | Description                | Example                      |
|--------------|---------|----------------------------|------------------------------|
| Success      | boolean | Indicates request success  | true                         |
| User         | object  | Details of the registered user | See nested fields below    |
| └ FirstName  | string  | User's first name          | TestUser123                  |
| └ LastName   | string  | User's last name           | TestUser123                  |
| └ EmailAddress | string | User's email address       | test333@example.com          |
| └ CreatedAt  | string  | Timestamp of account creation | 2025-01-18T17:00:13.9474518Z |


Response Body Example:
```json
{
    "Success": true,
    "User": {
        "FirstName": "TestUser123",
        "LastName": "TestUser123",
        "EmailAddress": "test333@example.com",
        "CreatedAt": "2025-01-18T17:00:13.9474518Z"
    }
}
```

**Example cURL**
```bash
curl -X POST http://localhost/dohabitsapp/v1/register \
-H "Content-Type: application/json" \
-d '{
    "FirstName": "TestUser123",
    "LastName": "TestUser123",
    "EmailAddress": "test222@example.com",
    "Password": "!secretPASSWORD123"
}'
```

### 2. Login
**Endpoint** `POST /dohabitsapp/v1/login`

**Request**

Request Headers:
Headers:
| Key            | Value            |
|----------------|------------------|
| Content-Type   | application/json |

**Request Body**:
| Field         | Type   | Description            | Example                     |
|---------------|--------|------------------------|-----------------------------|
| EmailAddress  | string | User's email address   | test333@example.com         |
| Password      | string | User's password        | secretPassword012!          |

**Example Request Body**:
```json
{
    "EmailAddress": "test333@example.com", /*"johndoe1@example.com"*/
    "Password": "secretPassword012!"
}
```

**Response:**

Response Headers:
| Key            | Value            |
|----------------|------------------|
| Authorization  | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDY3OTN9.5-pP_mSsdVUhVy6i7fSwLYZzi4ZDIKxGxEyyERlIQRQ |
| X-Csrf-Token   | OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs |
| Set-Cookie     | csrf_token=OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs; Path=/; HttpOnly; SameSite=Strict |

**Response Body**:
| Field         | Type    | Description                     | Example                      |
|---------------|---------|---------------------------------|------------------------------|
| Success       | boolean | Indicates request success       | true                         |
| User          | object  | Details of the logged-in user   | See nested fields below      |
| └ FirstName   | string  | User's first name               | TestUser123                  |
| └ LastName    | string  | User's last name                | TestUser123                  |
| └ EmailAddress | string | User's email address            | test333@example.com          |
| └ CreatedAt   | string  | Timestamp of account creation   | 2025-01-18T17:00:13.947Z     |
| LoggedInAt    | string  | Timestamp of login              | 2025-01-18T17:13:13.0799828Z|

**Example Response**:
```json
{
    "Success": true,
    "User": {
        "FirstName": "TestUser123",
        "LastName": "TestUser123",
        "EmailAddress": "test333@example.com",
        "CreatedAt": "2025-01-18T17:00:13.947Z"
    },
    "LoggedInAt": "2025-01-18T17:13:13.0799828Z"
}
```

**Example cURL**
```bash
curl -X POST http://localhost/dohabitsapp/v1/login \
-H "Content-Type: application/json" \
-d '{
    "EmailAddress": "test333@example.com",
    "Password": "secretPassword012!"
}'
```

### 3. Logout
**Endpoint** `POST /dohabitsapp/v1/logout`

**Request**

Request Headers:
| Key              | Value                                                                                                           |
|-------------------|-----------------------------------------------------------------------------------------------------------------|
| Content-Type      | application/json                                                                                              |
| X-CSRF-Token      | Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM                                                                    |
| Authorization     | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDY3OTN9.5-pP_mSsdVUhVy6i7fSwLYZzi4ZDIKxGxEyyERlIQRQ |

**Example cURL**:
```bash
curl -X POST http://localhost/dohabitsapp/v1/logout \
-H "Content-Type: application/json" \
-H "X-CSRF-Token: Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDY3OTN9.5-pP_mSsdVUhVy6i7fSwLYZzi4ZDIKxGxEyyERlIQRQ"
```

### 4. Refresh
**Endpoint** `POST /dohabitsapp/v1/refresh`

**Request**

Request Headers:
| Key              | Value                                                                                                           |
|-------------------|-----------------------------------------------------------------------------------------------------------------|
| Content-Type      | application/json                                                                                              |
| X-CSRF-Token      | Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM                                                                    |
| Authorization     | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDczNDR9.nX7pZO9q6otlF9Y_s_3-ktgjuOF0Zm6z6KFB0KCqL3A |

Request Body:
| Field         | Type   | Description            | Example                     |
|---------------|--------|------------------------|-----------------------------|
| EmailAddress  | string | User's email address   | test333@example.com         |

Request Body Example:

**Response:**
```json
{
    "EmailAddress": "test333@example.com"
}
```

Response Headers:

Response Body:
| Field         | Type    | Description                      | Example                     |
|---------------|---------|----------------------------------|-----------------------------|
| Success       | boolean | Indicates request success        | true                        |
| EmailAddress  | string  | Email address of the user        | test333@example.com         |

Response Body Example:
```json
{
    "Success": true,
    "EmailAddress": "test333@example.com"
}
```

**Example cURL**
curl -X POST http://localhost/dohabitsapp/v1/refresh \
-H "Content-Type: application/json" \
-H "X-CSRF-Token: Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDczNDR9.nX7pZO9q6otlF9Y_s_3-ktgjuOF0Zm6z6KFB0KCqL3A" \
-d '{
    "EmailAddress": "test333@example.com"
}'

## Habit Endpoints
The following endpoints are used to retrieve/manipulate the user's habits data. Each endpoint is protected via a short-lived JWT Access Token and Cross-Site Request Forgery Token (CSRF) which are required in the Request Header.

### 1. Create Habit
**Endpoint** `POST /dohabitsapp/v1/createhabit`

**Request**

Request Headers:
| Key            | Value                                                                                                           |
|----------------|---------------------------------------------------------------------------------------------------------------|
| Authorization  | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDczNzh9.AcTzMY18Lmc1uWCN90Bz_krGtXsaAkB4g_Z2UGrnlow |
| X-CSRF-Token   | X-CSRF-Token                                                                                                   |

Request Body:
| Field       | Type    | Description                       | Example           |
|-------------|---------|-----------------------------------|-------------------|
| Name        | string  | Name of the habit                | Meditate Daily    |
| Days        | integer | Current number of days completed | 30                |
| DaysTarget  | integer | Target number of days for habit  | 30                |


Request Body Example:
```json
{
    "Name": "Meditate Daily",
    "Days": 30,
    "DaysTarget": 30
}
```

**Response:**

Response Headers:
| Key            | Value            |
|----------------|------------------|
| X-Csrf-Token   | OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs |
| Set-Cookie     | csrf_token=OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs; Path=/; HttpOnly; SameSite=Strict |

Response Body:
| Field       | Type    | Description                       | Example           |
|-------------|---------|-----------------------------------|-------------------|
| Name        | string  | Name of the habit                | Meditate Daily    |
| Days        | integer | Current number of days completed | 30                |
| DaysTarget  | integer | Target number of days for habit  | 30                |

Response Body Example:
```json
{
    "Name": "Meditate Daily",
    "Days": 30,
    "DaysTarget": 30
}
```
```

**Example cURL**
```bash
curl -X POST http://localhost/dohabitsapp/v1/createhabit \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDczNzh9.AcTzMY18Lmc1uWCN90Bz_krGtXsaAkB4g_Z2UGrnlow" \
-H "X-CSRF-Token: X-CSRF-Token" \
-H "Content-Type: application/json" \
-d '{
    "Name": "Meditate Daily",
    "Days": 30,
    "DaysTarget": 30
}'
```

### 2. Retrieve Habit
**Endpoint** `GET /dohabitsapp/v1/retrievehabit?habitId={habitId}`

**Request**

Request Headers:
| Key              | Value                                                                                                           |
|-------------------|-----------------------------------------------------------------------------------------------------------------|
| Content-Type      | application/json                                                                                              |
| X-CSRF-Token      | Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM                                                                    |
| Authorization     | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDY3OTN9.5-pP_mSsdVUhVy6i7fSwLYZzi4ZDIKxGxEyyERlIQRQ |

**Response:**

Response Headers:
| Key            | Value            |
|----------------|------------------|
| X-Csrf-Token   | OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs |
| Set-Cookie     | csrf_token=OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs; Path=/; HttpOnly; SameSite=Strict |

Response Body:
| Field            | Type     | Description                               | Example                           |
|------------------|----------|-------------------------------------------|-----------------------------------|
| habitId          | string   | Unique identifier for the habit           | 67828988bfd0d3825fa10ee4          |
| userId           | string   | Unique identifier for the user            | 677ac7224620315e952dabd6          |
| createdAt        | datetime | Timestamp when the habit was created      | 2025-01-11T15:08:56.342Z          |
| name             | string   | Name of the habit                        | No Fap Updated                    |
| days             | integer  | Current number of days completed         | 30                                |
| daysTarget       | integer  | Target number of days for habit          | 60                                |
| completionDates  | array    | List of dates the habit was completed on | ["2025-01-01"]                    |

Response Body Example:
```json
{
    "habitId": "67828988bfd0d3825fa10ee4",
    "userId": "677ac7224620315e952dabd6",
    "createdAt": "2025-01-11T15:08:56.342Z",
    "name": "No Fap Updated",
    "days": 30,
    "daysTarget": 60,
    "completionDates": [
        "2025-01-01"
    ]
}
```

**Example cURL**
```bash
curl -X GET "http://localhost/dohabitsapp/v1/retrievehabit?habitId=67828988bfd0d3825fa10ee4" \
-H "Content-Type: application/json" \
-H "X-CSRF-Token: sXNia5iMYKiDUFyY0YcJLXJnh0GHeEwqpA4cXZQMiN0" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDczNDR9.nX7pZO9q6otlF9Y_s_3-ktgjuOF0Zm6z6KFB0KCqL3A"
```

### 3. Retrieve All Habits
**Endpoint** `GET /dohabitsapp/v1/retrievehabits`

**Request**

Request Headers:
| Key              | Value                                                                                                           |
|-------------------|-----------------------------------------------------------------------------------------------------------------|
| Content-Type      | application/json                                                                                              |
| X-CSRF-Token      | Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM                                                                    |
| Authorization     | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDY3OTN9.5-pP_mSsdVUhVy6i7fSwLYZzi4ZDIKxGxEyyERlIQRQ |

**Response:**

Response Headers:
| Key            | Value            |
|----------------|------------------|
| X-Csrf-Token   | OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs |
| Set-Cookie     | csrf_token=OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs; Path=/; HttpOnly; SameSite=Strict |

Response Body:
| Field            | Type     | Description                               | Example                           |
|------------------|----------|-------------------------------------------|-----------------------------------|
| habitId          | string   | Unique identifier for the habit           | 67828988bfd0d3825fa10ee4          |
| userId           | string   | Unique identifier for the user            | 677ac7224620315e952dabd6          |
| createdAt        | datetime | Timestamp when the habit was created      | 2025-01-11T15:08:56.342Z          |
| name             | string   | Name of the habit                        | No Fap Updated                    |
| days             | integer  | Current number of days completed         | 30                                |
| daysTarget       | integer  | Target number of days for habit          | 60                                |
| completionDates  | array    | List of dates the habit was completed on | ["2025-01-01"]                    |

Response Body Example:
```json
[
    {
        "habitId": "677ac8294620315e952dabd7",
        "userId": "677ac7224620315e952dabd6",
        "createdAt": "2024-09-21T11:30:00+01:00",
        "name": "Code everyday",
        "days": 30,
        "daysTarget": 66,
        "completionDates": [
            "2024-12-02",
            "2024-12-02",
            "2024-12-11"
        ]
    },
    {
        "habitId": "678be5466b92995d30e58dad",
        "userId": "678bde1d6b92995d30e58dac",
        "createdAt": "2025-01-18T17:30:46.674Z",
        "name": "Meditate Daily",
        "days": 30,
        "daysTarget": 30,
        "completionDates": []
    }
]
```

**Example cURL**
```bash
curl -X GET "http://localhost/dohabitsapp/v1/retrievehabits" \
-H "Content-Type: application/json" \
-H "X-CSRF-Token: sXNia5iMYKiDUFyY0YcJLXJnh0GHeEwqpA4cXZQMiN0" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDczNDR9.nX7pZO9q6otlF9Y_s_3-ktgjuOF0Zm6z6KFB0KCqL3A"
```

### 4. Update Habit
**Endpoint** `PUT /dohabitsapp/v1/updatehabit?habitId={habitId}`

**Request**

Request Headers:
| Key              | Value                                                                                                           |
|-------------------|-----------------------------------------------------------------------------------------------------------------|
| Content-Type      | application/json                                                                                              |
| X-CSRF-Token      | Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM                                                                    |
| Authorization     | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDY3OTN9.5-pP_mSsdVUhVy6i7fSwLYZzi4ZDIKxGxEyyERlIQRQ |

Request Body:
| Field            | Type     | Description                               | Example                           |
|------------------|----------|-------------------------------------------|-----------------------------------|
| name             | string   | Name of the habit                        | No Fap Updated                    |
| days             | integer  | Current number of days completed         | 30                                |
| daysTarget       | integer  | Target number of days for habit          | 60                                |
| completionDates  | array    | List of dates the habit was completed on | ["2025-01-01"]                    |


Request Body Example:
```json
{
    "Name": "Delete Me Again",
	"Days": 1,
	"DaysTarget": 60,
	"CompletionDates": ["2025-01-01"]
}
```

**Response:**

Response Headers:
| Key            | Value            |
|----------------|------------------|
| X-Csrf-Token   | OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs |
| Set-Cookie     | csrf_token=OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs; Path=/; HttpOnly; SameSite=Strict |

Response Body:
| Field            | Type     | Description                               | Example                           |
|------------------|----------|-------------------------------------------|-----------------------------------|
| name             | string   | Name of the habit                        | No Fap Updated                    |
| days             | integer  | Current number of days completed         | 30                                |
| daysTarget       | integer  | Target number of days for habit          | 60                                |
| completionDates  | array    | List of dates the habit was completed on | ["2025-01-01"]                    |

Response Body Example:
```json
{
    "Name": "Delete Me Again",
	"Days": 1,
	"DaysTarget": 60,
	"CompletionDates": ["2025-01-01"]
}
```

**Example cURL**
```bash
curl -X PUT "http://localhost/dohabitsapp/v1/updatehabit?habitId=678be5466b92995d30e58dad" \
  -H "Content-Type: application/json" \
  -H "X-CSRF-Token: sXNia5iMYKiDUFyY0YcJLXJnh0GHeEwqpA4cXZQMiN0" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3M373MDczNDR9.nX7pZO9q6otlF9Y_s_3-ktgjuOF0Zm6z6KFB0KCqL3A" \
  -d '{
    "Name": "Delete Me Again",
    "Days": 1,
    "DaysTarget": 60,
    "CompletionDates": ["2025-01-01"]
}'
```

### 5. Delete Habit
**Endpoint** `DELETE /dohabitsapp/v1/deletehabit?habitId={habitId}`

**Request**

Request Headers:
| Key              | Value                                                                                                           |
|-------------------|-----------------------------------------------------------------------------------------------------------------|
| Content-Type      | application/json                                                                                              |
| X-CSRF-Token      | Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM                                                                    |
| Authorization     | Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDY3OTN9.5-pP_mSsdVUhVy6i7fSwLYZzi4ZDIKxGxEyyERlIQRQ |

Request Body:

Request Body Example:

**Response:**

Response Headers:
| Key            | Value            |
|----------------|------------------|
| X-Csrf-Token   | OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs |
| Set-Cookie     | csrf_token=OdPd7MYHVUlwjPxpTuF_D4IohzmUsmZOzJLOQYz7Vhs; Path=/; HttpOnly; SameSite=Strict |

Response Body:
| Field            | Type     | Description                               | Example                           |
|------------------|----------|-------------------------------------------|-----------------------------------|
| success          | boolean  | Indicates success                         |                                   |


Response Body Example:
```json
{
    "success": true
}
```

**Example cURL**
```bash
curl -X DELETE "http://localhost/dohabitsapp/v1/deletehabit?habitId=67828a01bfd0d3825fa10ee5" \
  -H "Content-Type: application/json" \
  -H "X-CSRF-Token: Xjn3I8OU_kgBLRfa1DlKOX-Zk9JuArNiE47gqLNPHCM" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QzMzNAZXhhbXBsZS5jb20iLCJleHAiOjE3MzczMDY3OTN9.5-pP_mSsdVUhVy6i7fSwLYZzi4ZDIKxGxEyyERlIQRQ"
```


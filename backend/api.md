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
**Endpoint** `GET /dohabitsapp/v1/retrievehabit`

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

Response Body Example:

**Example cURL**

### 3. Retrieve All Habits
**Endpoint** `GET /dohabitsapp/v1/retrievehabits`

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

Response Body Example:

**Example cURL**

### 4. Update Habit
**Endpoint** `PUT /dohabitsapp/v1/updatehabit`

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

Response Body Example:

**Example cURL**

### 5. Delete Habit
**Endpoint** `DELETE /dohabitsapp/v1/deletehabit`

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

Response Body Example:

**Example cURL**

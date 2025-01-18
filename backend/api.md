# DoHabitsApp API Documentation

This document outlines the endpoints for the DoHabitsApp API. 

Each endpoint is prefixed with the API name and version, e.g., `/dohabitsapp/v1`.

## User Endpoints
The following endpoints create or control the user's session state.

### 1. Register
**Endpoint** `POST /dohabitsapp/v1/register`

Registers the user. Validates user data, if successful adds their credentials to the DB Layer.

**Request:**

Headers:
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

Headers:

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







Explanation of the endpoint.
Request Headers, Query Params, Body.
Response Headers, Body, Status Code(s).
Example data.
Example cURL for Postman.



#### Response


#### Example cURL
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

#### Request


#### Response


#### Example cURL

### 3. Logout
**Endpoint** `POST /dohabitsapp/v1/logout`

#### Request


#### Response


#### Example cURL

### 4. Refresh
**Endpoint** `POST /dohabitsapp/v1/refresh`

#### Request


#### Response


#### Example cURL

## Habit Endpoints
The following endpoints are used to retrieve/manipulate the user's habits data. Each endpoint is protected via a short-lived JWT Access Token and Cross-Site Request Forgery Token (CSRF) which are required in the Request Header.

### 1. Create Habit
**Endpoint** `POST /dohabitsapp/v1/createhabit`

#### Request


#### Response


#### Example cURL

### 2. Retrieve Habit
**Endpoint** `GET /dohabitsapp/v1/retrievehabit`

#### Request


#### Response


#### Example cURL

### 3. Retrieve All Habits
**Endpoint** `GET /dohabitsapp/v1/retrievehabits`

#### Request


#### Response


#### Example cURL

### 4. Update Habit
**Endpoint** `PUT /dohabitsapp/v1/updatehabit`

#### Request


#### Response


#### Example cURL

### 5. Delete Habit
**Endpoint** `DELETE /dohabitsapp/v1/deletehabit`

#### Request


#### Response


#### Example cURL

# DoHabitsApp API Documentation

This document outlines the endpoints for the DoHabitsApp API. 

Each endpoint is prefixed with the API name and version, e.g., `/dohabitsapp/v1`.

## User Endpoints
The following endpoints create or control the user's session state.

### 1. Register
**Endpoint** `POST /dohabitsapp/v1/register`

Registers the user. Validates user data, if successful adds their credentials to the DB Layer.

**Headers**:
| Key            | Value            |
|----------------|------------------|
| Content-Type   | application/json |

**Request Body**:
Example Request Body Data:
```json
{
    "FirstName": "TestUser123",
    "LastName": "TestUser123",
    "EmailAddress": "test222@example.com",
    "Password": "!secretPASSWORD123"
}
```





Explanation of the endpoint.
Request Headers, Query Params, Body.
Response Headers, Body, Status Code(s).
Example data.
Example cURL for Postman.

#### Request



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

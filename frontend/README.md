# DoHabitsApp

## Description
DoHabitsApp frontend - controls the UI layer.

## Endpoints ( see /backend/api.md for more details)
### User Endpoints
1. **Register**: `POST /dohabitsapp/v1/register` - /register page
2. **Login**: `POST /dohabitsapp/v1/login` - /login page
3. **Logout**: `POST /dohabitsapp/v1/logout` - Logout button
4. **Refresh**: `POST /dohabitsapp/v1/refresh` - Sent on protected endpoints like the Habit Endpoints if the DoHabitsAPI returns a 401 Unauthorized Response to get a new JWT access-token.

### Habit Endpoints
1. **Create Habit**: `POST /dohabitsapp/v1/createhabit` - Insert Button
2. **Retrieve Habit**: `GET /dohabitsapp/v1/retrievehabit`
3. **Retrieve All Habits**: `GET /dohabitsapp/v1/retrievehabits` - Called on the / directory if the user's logged in. Called each time a new habit is created or deleted.
4. **Update Habit**: `PUT /dohabitsapp/v1/updatehabit` - Update Button
5. **Delete Habit**: `DELETE /dohabitsapp/v1/deletehabit` - Delete Button

## Architecture / Design Pattern
ReactJS, Model Controller Architecture for habits/user services.

## Session
JWT Tokens:
- Access Token - Short-Lived JWT Token - stored in session as access-token - sent in the Request Header for protected endpoints that required authentication..
- Cross-Site Request Forgery Token (CSRF) - stored in session as csrf - send in the Request Header for protected endpoints.

## Tech Stack
ReactJS, NextJS, TypeScript, Bootstrap CSS

# Dependencies
If the .env ENVIRONMENT var is not DEV then requires the DoHabitsAPI

# Build See base README.md

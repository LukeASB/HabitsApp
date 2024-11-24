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
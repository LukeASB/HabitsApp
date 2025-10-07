# DoHabitsApp
The DoHabitsApp is a web app that allows the user to add a number of habits to track daily on a calendar.

Minimum Viable Product:
Web App that allows individual users to track their habits by ticking the day of on the calendar.

Users should be able to register/login to view all their habits on the home page and individual habits by clicking on them in a navigation menu.

Users should be able to create, update and delete their habits.

# View
**Login:**
![image](https://github.com/user-attachments/assets/8fcb6fdd-981a-4f17-9494-0613d755784d)


**Register:**
![image](https://github.com/user-attachments/assets/7d62ca42-a277-4be3-8af1-1dc86a24273f)

**Home:**
![image](https://github.com/user-attachments/assets/50716ec9-df0a-41e5-8ed7-5b5fb5efa755)

**Create Habit:**
![image](https://github.com/user-attachments/assets/91ed0cbc-0789-4462-805b-fcce2bf6078f)

**Update Habit:**
![image](https://github.com/user-attachments/assets/5b55869b-75d2-42f5-b976-c1a7c7c85326)

**Delete Habit:**
![image](https://github.com/user-attachments/assets/eaf23368-3b58-4982-813e-58a4bae19519)

# Demo Video
<div align="center">
https://www.youtube.com/watch?v=PHqA0KKy7u0
</div>
<div align="center">
  <a href="https://www.youtube.com/watch?v=PHqA0KKy7u0">
    <img src="https://img.youtube.com/vi/PHqA0KKy7u0/0.jpg" alt="Demo Video">
  </a>
</div>

# Diagram - High Level Overview
![HabitsAppWorkflow](https://github.com/user-attachments/assets/cc51ee7b-7bbf-4523-91bf-76830247a3d1)

![HabitsApp](https://github.com/user-attachments/assets/b8c84b40-50b2-4406-83b7-d80cf43f8764)

Credit: https://sequencediagram.org/

# Tech Stack
- Backend: Golang, MongoDB
- Frontend: React, Next.js, TypeScript, Bootstrap CSS

# Running The Application
## Setting Up Dependencies:
### MongoDB
Create a habitsapp database in MongoDB with the following collections: 
- habits
- user_session
- users

Example documents data can be seen in the /backend/README.md

### Backend
Navigate to the `backend` directory:
```sh
cd backend
```
Install dependencies:
```sh
go mod tidy
```
Create a `.env` file based on `.example_env` (For example, refer to [backend/.example_env](backend/.example_env).
```
DB_TYPE=mockdb (Note: this dictates the DB you want to use via Strategy Design Pattern).
DB_URL=connectionstring@example:username/password
DB_NAME=habitsapp
USERS_COLLECTION=users
USER_SESSION_COLLECTION=user_session
HABITS_COLLECTION=habits
PORT=80
SITE_URL=http://localhost
API_NAME=dohabitsapp
APP_VERSION=1.0
API_VERSION=v1
LOG_VERBOSITY=2
JWT_SECRET=your_secret_key
```
If you want to run locally - Run the application:
```sh
go run main.go
```

### Frontend
Navigate to the `frontend` directory:
```sh
cd frontend
```
Install dependencies:
```sh
npm install
```
Create a `.env` file based on `.example_env` (For example, refer to [frontend/.example_env](frontend/.example_env).
```
API_URL=dohabitsapp/v1 
ENVIRONMENT=DEV (runs mock data) | PROD (runs data from the backend)
```
If you want to run locally - Run the application:
```sh
npm run dev
```

next.config.mjs should be set up appropriately depending on how you want to run the application.

If running the application with Docker:
Ensure `frontend/next.config.mjs` is the following so it redirects to the habitsappbackend:
```js
  async rewrites() {
    return [
      {
        source: "/api/:path*", // Match API requests
        destination: "http://habitsappbackend/:path*", // Forward to backend
      },
    ];
  },
```

If running the application without Docker:
Ensure `frontend/next.config.mjs` is the following so it redirects to the localhost:
```js
  async rewrites() {
    return [
      {
        source: "/api/:path*", // Match API requests
        destination: "http://localhost/:path*", // Forward to backend
      },
    ];
  },
```

## With Docker Compose
1. Create Images: lukesbdev/backend:habitsappbackend2025, lukesbdev/frontend:habitsappfrontend2025
```sh
    # Step 1: Build the Docker image
    docker image build -t lukesbdev/backend:habitsappbackend2025 .

    # Step 2: Push the Docker image to a registry
    docker login
    docker image push lukesbdev/backend:habitsappbackend2025
```
```sh
    # Step 1: Build the Docker image
    docker image build -t lukesbdev/frontend:habitsappfrontend2025 .

    # Step 2: Push the Docker image to a registry
    docker login
    docker image push lukesbdev/frontend:habitsappfrontend2025
```

2. Build and run the Docker containers:
```sh
docker-compose up --build
```
3. Access the application at `http://localhost:3000` for the frontend and `http://localhost:80` for the backend.

## Without Docker Compose
1. Create the network so containerised apps can communicate.
```sh
    docker network create habitsapp-network # Create the habitsapp-network network

    docker network ls # Check habitsapp-network is present in the list.
```

2. Frontend - create network, image and container
```sh
    # Step 1: Build the Docker image
    docker image build -t lukesbdev/frontend:habitsappfrontend2025 .

    # Step 2: Push the Docker image to a registry
    docker login
    docker image push lukesbdev/frontend:habitsappfrontend2025

    # Step 3: Run the Docker container
    docker run -d -p 3000:3000 --name habitsappfrontend lukesbdev/frontend:habitsappfrontend2025

    # Step 4: Verify the running container
    docker ps

    # Step 5: Checks Logs
    docker logs habitsfrontend

    # Optional how to stop and remove
    # Step 6: Stop the running container
    docker stop habitsappfrontend

    # Step 7: Remove the stopped container
    docker rm habitsappfrontend

    docker rmi lukesbdev/frontend:habitsappfrontend2025
```

3. Backend - create network, image and container
```sh
    # Step 1: Build the Docker image
    docker image build -t lukesbdev/backend:habitsappbackend2025 .

    # Step 2: Push the Docker image to a registry
    docker login
    docker image push lukesbdev/backend:habitsappbackend2025

    # Step 3: Run the Docker container
    docker run -d -p 8000:80 --name habitsappbackend lukesbdev/backend:habitsappbackend2025

    # Step 4: Verify the running container
    docker ps

    # Optional how to stop and remove
    # Step 5: Stop the running container
    docker stop habitsappbackend

    # Step 6: Remove the stopped container
    docker rm habitsappbackend

    # Step 7: Remove the Docker image (optional)
    docker rmi lukesbdev/backend:habitsappbackend2025
```

## Without Docker
1. **Backend**:
Follow the "Setting Up Dependencies" section above.

Run the application:
```sh
  go run main.go
```

2. **Frontend**:
Follow the "Setting Up Dependencies" section above.

Run the development server:
```sh
  npm run dev
```

## Futher Information
See the Frontend/Backend README.md on how to execute.

## Found this useful?
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/lukesb)

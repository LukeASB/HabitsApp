# DoHabitsApp
The DoHabitsApp is a web app that allows the user to add a number of habits to track daily on a calendar.

Minimum Viable Product:
Web App that allows individual users to track their habits by ticking the day of on the calendar.

Users should be able to register/login to view all their habits on the home page and individual habits by clicking on them in a navigation menu.

Users should be able to create, update and delete their habits.

# View
**Login:**

**Register:**

**Home:**
![title](https://raw.githubusercontent.com/LukeASB/BookStoreApp/master/View-home.png)

**Create Habit:**
![Add Book Screenshot](https://raw.githubusercontent.com/LukeASB/BookStoreApp/master/View-addbook.png)

**Update Habit:**
![Add Book Screenshot](https://raw.githubusercontent.com/LukeASB/BookStoreApp/master/View-addbook.png)

**Delete Habit:**
![Add Book Screenshot](https://raw.githubusercontent.com/LukeASB/BookStoreApp/master/View-addbook.png)

# Demo Video
To do

# Tech Stack
- Backend: Golang, MongoDB
- Frontend: React, Next.js, TypeScript, Bootstrap CSS

# Running The Application
## With Docker
1. Build and run the Docker containers:
    ```sh
    docker-compose up --build
    ```
2. Access the application at `http://localhost:3000` for the frontend and `http://localhost:8080` for the backend.
   
## Without Docker
1. **Backend**:
    1. Navigate to the `backend` directory:
        ```sh
        cd backend
        ```
    2. Install dependencies:
        ```sh
        go mod tidy
        ```
    3. Create a `.env` file based on `.example_env` (For example, refer to [backend/.example_env](backend/.example_env).
       ```
        DB_URL=connectionstring@example:username/password
        ENV=dev (runs mock data in mock_db.go) | PROD (runs data from the Database)
        PORT=8080
        SITE_URL=http://localhost
        API_NAME=dohabitsapp
        APP_VERSION=1.0
        API_VERSION=v1
        LOG_VERBOSITY=2
        JWT_SECRET=your_secret_key
       ```
    5. Run the application:
        ```sh
        go run main.go
        ```

2. **Frontend**:
    1. Navigate to the `frontend` directory:
        ```sh
        cd frontend
        ```
    2. Install dependencies:
        ```sh
        npm install
        ```
    3. 3. Create a `.env` file based on `.example_env` (For example, refer to [frontend/.example_env](frontend/.example_env).
       ```
       API_URL=dohabitsapp/v1 
       ENVIRONMENT=DEV (runs mock data) | PROD (runs data from the backend)
       ```
       
    5. Run the development server:
        ```sh
        npm run dev
        ```
    6. Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.   

## Futher Information
See the Frontend/Backend README.md on how to execute.

## Found this useful?
[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/lukesb)

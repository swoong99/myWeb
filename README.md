# myWeb

## Project Abstract
This project consists Backend written in **Go** using the **Gin** framework and Frontend is written in **Vue.js**. The backend reads data from json files and provide it through API endpoints. The frontend receives data and display it on the web.

## How to run the project
1. **make sure Docker is installed**
```
    docker --version
```

2. **Clone the Repository**
```
    git clone https://github.com/swoong99/myWeb.git myWeb
    cd myWeb
```
3. **Run with Docker Compose**
```
    docker-compose up --build
```
4. **Access the Web**
- Open your browser and go to 'http://localhost' to view the app.

## Backend
- The backend provides an endpoint like `GET /api/instruction` that returns data in JSON format.

## Frontend
- The frontend uses axios library to send requests to the backend for data.
- The fronend uses router library to navigate between three sections, **introduction**, **projects**, and **research**.

## Tasks to do
- The json files are now in the same directory. We can make DB to manage them.
- We can handle other data types, for example images.
- The website works only in local network. Find the way to visit the website with devices in other network using specific url?
- Not only read data, I want to write data from website, like comments. (After constructing DB)

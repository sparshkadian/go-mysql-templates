# Go MySQL Templates Project

## Overview
This project is a basic web application developed as part of a Go learning series. It allows users to manage tasks and includes features like homepage for adding tasks and task listing.

## Technologies Used
- Go (Golang)
- MySQL (as the database)
- HTML/Template (for rendering tasks)
- [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) (MySQL driver)
- [github.com/joho/godotenv](https://github.com/joho/godotenv) (for loading environment variables)

## Installation
1. Clone this repository: `git clone https://github.com/yourusername/golang-mysql-templates.git`
2. Navigate to the project directory: `cd golang-mysql-templates`
3. Create a `.env` file in the root directory and add the following environment variables:
    ```
    DB_USER=your_mysql_username
    DB_PASSWORD=your_mysql_password
    DB_HOST=your_mysql_host
    DB_PORT=your_mysql_port
    DB_NAME=your_database_name
    ```
4. Build the project: `go build`
5. Run the executable: `./golang-mysql-templates`

## Usage
Make requests to the API endpoints using tools like cURL, Postman, or your preferred HTTP client.
Example:
- To add a new task:
    ```
    curl -X POST -H "Content-Type: application/json" -d '{"task":"Finish project report","username":"john_doe"}' http://localhost:8080/
    ```
- To retrieve all tasks:
    ```
    curl http://localhost:8080/tasks
    ```

## Learning Journey
This project has been an invaluable learning experience in mastering Go. Here are some key aspects of the learning process:
- **Motivation:** Inspired by Go's simplicity, performance, and industry popularity.
- **Challenges:** Overcoming hurdles in learning a new language, mitigated by Go's clear syntax and extensive standard library.
- **Progress:** Deeper understanding of Go's features like database interaction, templating, and HTTP request handling through project development.
- **Community:** Leveraging support and resources from the vibrant Go community through forums, documentation, and open-source projects.


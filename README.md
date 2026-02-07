# Student API with Gin (Homework 4)

## How to Run
1. Navigate to the `go-api-gin` folder.
2. Run the command: `go mod tidy`
3. Run the command: `go run main.go`

## API Endpoints
- **GET** `/students` : List all students
- **POST** `/students` : Create a new student (Required: id, name, major, gpa)
- **PUT** `/students/:id` : Update student details by ID
- **DELETE** `/students/:id` : Delete a student by ID

## Features
- Layered Architecture (Models, Handlers)
- Input Validation (GPA must be 0.00 - 4.00)
- SQLite Database integration
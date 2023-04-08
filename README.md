# School Management Backend Server

This is a backend server for a school management system. It provides APIs for managing students, teachers, courses, and classrooms.

## Installation

To install this project, you will need to have [Golang](https://go.dev/) and [MySql](https://www.mysql.com/downloads/) installed on your system. Once you have these dependencies installed, you can follow these steps:

1. Clone this repository:

git clone https://github.com/manlikeNacho/School-Management-Backend-Server.git


2. Install the dependencies:

go get i

3. Start Server

go run . Or go run main.go


This will start the server on port 8080 by default. You can configure the port by setting the `PORT` environment variable.

## Usage

This server provides the following APIs:

- `GET /students`: Get a list of all students
- `GET /students/:id`: Get a single student by ID
- `POST /students`: Create a new student
- `PUT /students/:id`: Update an existing student
- `DELETE /students/:id`: Delete a student by ID
- `GET /teachers`: Get a list of all teachers
- `GET /teachers/:id`: Get a single teacher by ID
- `POST /teachers`: Create a new teacher
- `PUT /teachers/:id`: Update an existing teacher
- `DELETE /teachers/:id`: Delete a teacher by ID
- `GET /courses`: Get a list of all courses
- `GET /courses/:id`: Get a single course by ID
- `POST /courses`: Create a new course
- `PUT /courses/:id`: Update an existing course
- `DELETE /courses/:id`: Delete a course by ID
- `GET /classrooms`: Get a list of all classrooms
- `GET /classrooms/:id`: Get a single classroom by ID
- `POST /classrooms`: Create a new classroom
- `PUT /classrooms/:id`: Update an existing classroom
- `DELETE /classrooms/:id`: Delete a classroom by ID

You can use these APIs to manage the school's data. You can send requests to these APIs using a tool like [Postman](https://www.postman.com/) or by writing your own client.

## Contributing

If you would like to contribute to this project, please feel free to submit a pull request. Before submitting a pull request, please make sure that your changes are tested and follow the [code of conduct](CODE_OF_CONDUCT.md).

## License

This project is licensed under the [MIT License](LICENSE).

## Contact

If you have any questions or issues, please feel free to [open an issue](https://github.com/manlikeNacho/School-Management-Backend-Server/issues/new) or contact me directly at [eiheanacho52@gmail.com](mailto:eiheanacho52@gamil.com).

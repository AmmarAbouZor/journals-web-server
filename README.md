# Journals-Web-Server

Journals-Web-Server is a backend server written in Go that provides CRUD (Create, Read, Update, Delete) operations for managing journals. It serves as the backend for the [TUI-Journal](https://github.com/AmmarAbouZor/tui-journal) application.

## Features

- Create new journals with a title, date, and content.
- Retrieve a list of all journals.
- Update the title, date, or content of a journal.
- Delete a journal.

## Prerequisites

- Go 1.16 or above
- SQLite 3

## Getting Started

1. Clone the repository:

```shell
git clone https://github.com/AmmarAbouZor/journals-web-server.git
```

2. Change to the project directory:

```shell
cd journals-web-server
```

3. Install the dependencies:

```shell
go mod download
```

4. Run the server:

```shell
go run main.go
```

By default, the server will start on port 8080. If you want to use a different port, you can set the `PORT` environment variable. For example:

```shell
export PORT=8000
go run main.go
```

The server will now start on port 8000.

By default, the server will automatically creat the SQLite database file if it doesn't exist. The database file is named `journals.db` and is created in the same directory as the server executable. You don't need to create the database file yourself. You can set the path of the sqlite database file via the `DB_PATH` environment variable

```shell
export DB_PATH=<Path_to_database>
go run main.go
```

## Configuration

The server uses environment variables for configuration. The following variables can be set:

- `PORT` - Specifies the port on which the server should listen. If not set, it defaults to 8080.
- `DB_PATH` - Specifies the path to the SQLite database file. The default is `journals.db` in the server's directory.

## Endpoints

The following endpoints are available:

- `GET /journals` - Retrieves a list of all journals.
- `POST /journals` - Creates a new journal.
- `PUT /journals` - Updates an existing journal.
- `DELETE /journals/?id` - Deletes a journal with the specified ID.

For detailed information about each endpoint and the expected request/response formats, refer to the API documentation or the code comments.
Certainly! Here's a simplified roadmap with checkmarks for adding authentication and creating a Docker image:

## Roadmap

- [ ] Authentication
  - [ ] Choose an authentication mechanism
  - [ ] Implement authentication in the server code
  - [ ] Add endpoints for user registration, login, and logout
  - [ ] Implement middleware for request authentication
  - [ ] Optional: Add additional features like password reset or RBAC

- [ ] Docker Image

Feel free to update and expand the roadmap based on your specific implementation and requirements.

## Contributing

Contributions to the Journals Web Server are welcome! If you encounter any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

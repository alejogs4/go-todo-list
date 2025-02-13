# Go Todo List Application

This is a simple Todo List application built with Golang for the backend and served by an Nginx server. The application uses PostgreSQL as the database and Docker Compose to manage the services.

## Prerequisites

- Docker
- Docker Compose
- Make

### Build and Run the Application

To build and run the application, use the following command:

```sh
make run
```

This command will start the Golang application, Nginx server, and PostgreSQL database using Docker Compose.

This will run client application in port 80

### Running Tests

To run the tests for the Golang application, use the following command:


```sh
make test
```

### Running Integration Tests

To run the integration tests for the Golang application, use the following command:

```sh
make test-integration
```

### Cleaning Up

To stop and remove the Docker containers, networks, and volumes, use the following command:

```sh
make clean
```
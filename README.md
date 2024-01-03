# Snippetbox

## Project Overview

Snippetbox is a simple but robust application for sharing code snippets.
It's built with Go and uses PostgreSQL for data storage.
The application allows users to create, view, and manage their code snippets.

## Features

- User authentication: Sign up, log in, and log out.
- Snippet management: Create, view, and delete snippets.
- Secure: All data is transmitted over HTTPS.

## Technologies Used

- Go: The main programming language used.
- Gorilla/mux: The HTTP router and URL matcher for building Go web servers.
- PostgreSQL: The database used for storing data.
- Docker: Used for creating and managing the application's environment.
- Docker Compose: Used for defining and running multi-container Docker applications.

## Project Structure

The project is organized into several directories:

- `cmd/web`: Contains the application entry point and web server logic, as well as the handlers, middleware, and templates for the web interface.
- `internal/models`: Contains the data models used in the application.
- `ui`: Contains the HTML templates and static files for the web interface.
- `db/migrations`: Contains SQL migration scripts for the database.
- `validator`: Contains validation logic for the application.
- `tls`: Contains TLS configuration for the application.

## Getting Started

### Prerequisites

- Go 1.16 or later
- Docker and Docker Compose for running the PostgreSQL database

### Running the Application

    1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/snippetbox.git
    cd snippetbox
    ```

    2. Start the PostgreSQL database with Docker Compose:

    ```sh
    make docker_compose_up
    ```

    3. Run the database migrations:

    ```sh
    make migrateup
    ```

    4. Start the application:

    ```sh
    go run ./cmd/web
    ```

The application will be available at <https://localhost:4000>.

### Features

- View the latest code snippets on the home page
- View individual snippets
- Create new snippets
- Delete snippets
- Sign up for an account
- Sign in and out of your account

### Demo

<https://github.com/denim-bluu/go-snippetbox-app/assets/66572804/f70ef134-f58a-49dc-a733-616714f99051>

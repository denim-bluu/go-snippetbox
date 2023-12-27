# Snippetbox

Snippetbox is a simple web application for sharing code snippets, built with Go.

## Project Structure

The project is organized into several directories:

- `cmd/web`: Contains the application entry point and web server logic.
- `internal/models`: Contains the data models used in the application.
- `ui`: Contains the HTML templates and static files for the web interface.
- `db/migrations`: Contains SQL migration scripts for the database.

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

The application will be available at http://localhost:4000.

### Features
View the latest code snippets on the home page
View individual snippets
Create new snippets
Delete snippets

# tiny-route

A scalable URL Shortener service built in Go, designed to efficiently handle high traffic with a custom load balancer and an autoscaling system. Initially, MongoDB is used as the primary database for storing shortened URLs, with plans to add Redis for caching and further optimization.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Setup and Installation](#setup-and-installation)
- [Usage](#usage)
- [Future Improvements](#future-improvements)
- [Contributing](#contributing)

## Features

- **URL Shortening**: Generate short URLs that redirect to longer URLs.
- **CRUD for Users**: Manage and view URLs created by individual users.
- **Scalability**: Custom load balancing and autoscaling strategies to handle increased traffic.
- **Monitoring and Logging**: Integrated metrics collection and centralized logging.

## Tech Stack

- **Backend**: Golang (for high performance and efficient concurrency)
- **Database**: MongoDB (initial primary database)
- **Containerization**: Docker (to package the application in containers)
- **Orchestration**: Kubernetes (planned for scaling the service)
- **Load Balancing**: Custom load balancer in Go with strategies like Round Robin, Least Connections, and Health Checks
- **Monitoring**: Prometheus for metrics, Grafana for visualization

## Architecture

The architecture is designed for scalability and ease of maintenance:

1. **API Layer**: Handles incoming requests and interacts with the core logic of URL shortening.
2. **Core Service Layer**: Manages core features such as URL creation, expansion, validation, and analytics.
3. **Persistence Layer**: Interfaces with MongoDB for storage.
4. **Load Balancer**: Custom-built in Go to manage traffic distribution across instances.
5. **Auto-Scaling**: Using Kubernetes Horizontal Pod Autoscaler (HPA) for dynamic scaling based on metrics like CPU and memory usage.

## Setup and Installation

### Prerequisites

- **Docker**: Ensure Docker is installed and running.
- **Docker Compose**: Required for orchestrating multi-container setups.

### Installation Steps

1. Clone the repository:
    ```bash
    git clone https://github.com/pabloespinosa12/tiny-route.git
    cd tiny-route
    ```

2. Run the application with Docker Compose:
    ```bash
    docker-compose up --build
    ```

3. The API will be available at `http://localhost:8080` (configurable in Docker Compose).

### Environment Variables

The application uses the following environment variables (configure as needed in `docker-compose.yml`):

- `MONGO_URI`: MongoDB connection string (e.g., `mongodb://mongo:27017/mydatabase`)
- `MONGO_INITDB_ROOT_USERNAME` and `MONGO_INITDB_ROOT_PASSWORD`: Credentials for MongoDB

## Usage

To shorten a URL or manage user data, you can make HTTP requests to the API.

### Endpoints (Examples)

- **POST** `/shorten` – Shorten a new URL
- **GET** `/expand/:shortURL` – Expand a shortened URL
- **GET** `/users/:userId/urls` – List URLs for a specific user
- **POST** `/users` – Create a new user
- **PUT** `/users/:userId` – Update user details

### Example Request

```bash
curl -X POST http://localhost:8080/shorten \
    -H "Content-Type: application/json" \
    -d '{"longUrl": "https://example.com/long-url"}'

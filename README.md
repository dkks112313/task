# User Activity Tracking Service

A Go REST API service for tracking user activity events and generating 4-hours aggregated statistics.

## Features

- **Event Recording**: Capture user activity events with flexible metadata
- **Event Retrieval**: Query events by user and date range
- **Daily Aggregation**: Background job that calculates event counts per user every 4 hours
- **React Client**: Minimal web interface to view user events
- **Monitoring**: Optional Grafana integration for metrics and logs

## Tech Stack

- **Backend**: Go REST API
- **Database**: PostgreSQL
- **Frontend**: React
- **Containerization**: Docker & Docker Compose
- **Monitoring**: Grafana (optional)

## Ports
- **Backend**: app:8080
- **Database**: postgres:5432
- **Frontend**: interface:80
- **Grafana**: grafana:3000
- **Prometheus**: prometheus:9090

## Quick Start

### Prerequisites

- Docker
- Docker Compose

### Example .env file
Create .env file into root backend folder
``` bash
DB_TYPE=postgres
DB_USER=user
DB_PASSWORD=admin
DB_HOST=postgres
DB_PORT=5432
DB_NAME=basedb
```

### Local Development

```bash
# Clone the repository
git clone https://github.com/dkks112313/task.git
cd task

# Copy environment file(Oprional, can create)
cp .env.example .env

# Start services
docker-compose up -d
```

## How work with back-end

### How to send event to back-end
You can send to route POST /events, like this json
``` bash
{
    "user_id": 1,
    "action": "click",
    "metadata": {"page": "/page"}
}
```

### How to get event from back-end
You can get from route GET /events?query=, like this
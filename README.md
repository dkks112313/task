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

##Frontend

Access at http://localhost

Use filters panel to query events by user_id, from, to, action, metadata.

Events table displays: User ID, Action, Metadata, ISO-8601 Timestamp.

##Back-end API
###Create Event

POST /events
Request body example:
``` bash
{
    "user_id": 1,
    "action": "click",
    "metadata": {"page": "/page"}
}
```

###Retrieve Events
``` bash
GET /events with optional query parameters:
```
user_id — optional, filter by user ID

from — optional, ISO-8601 date string start of range

to — optional, ISO-8601 date string end of range

action — optional, filter by action

metadata — optional, filter by metadata

Example:

GET /events?user_id=1&from=2025-10-15T10:00:00&to=2025-10-15T22:00:00&action=click

##Aggregation Job

Runs every 4 hours.

Selects events from the last aggregation run.

Counts events per user.

Inserts aggregated data into user_event_stats table with columns:
user_id, start_time, end_time, event_count.

Updates the last aggregation timestamp for next run.

##Optional Monitoring

Prometheus collects metrics from the Go service.

Grafana dashboards can visualize event counts and system metrics.

Ports: prometheus:9090, grafana:3000.

##Notes

All timestamps are stored in UTC (ISO-8601) in the database.

Frontend datetime-local filters are converted to ISO-8601 before sending to backend.

Aggregation job ensures lightweight queries by storing precomputed counts.
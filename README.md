# User Activity Tracking Service

A Go REST API service for tracking user activity events and generating daily aggregated statistics.

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

## Quick Start

### Prerequisites

- Docker
- Docker Compose

### Local Development

```bash
# Clone the repository
git clone <repository-url>
cd user-activity-service

# Copy environment file
cp .env.example .env

# Start services
docker-compose up -d

# Run migrations
docker-compose exec api go run cmd/migrate/main.go

# GoExpert Weather API Lab

This project is a sample Go application developed for the Full Cycle GoExpert course. It demonstrates a simple weather API with Docker and Google Cloud Run deployment support.

## Features
- RESTful API to fetch weather information
- Modular Go project structure
- Unit tests for API endpoints
- Docker and Docker Compose support
- Ready for deployment on Google Cloud Run


## Getting Started

### Prerequisites
- Go 1.20 or higher
- Docker (for containerization)

### Running Locally

#### 1. Build and Run
```bash
go build -o weather-api ./cmd/server
./weather-api
```

#### 2. Run with Docker
```bash
docker build -t weather-api .
docker run -p 8080:8080 weather-api
```

#### 3. Run with Docker Compose
```bash
docker-compose up --build
```

### API Usage
- Base URL: `http://localhost:8080`
- Example request: `http://localhost:8080?cep=04649000`

### Running Tests
```bash
go test ./...
```


### App URL on Google Cloud 
https://deploy-gcloudrun-fullcycle-lab-01-159735979117.us-central1.run.app


# Alarm Microservice (GoLang + Gin)

A lightweight alarm and notification microservice with in-memory store, REST APIs, and periodic notification support.

## 🛠 Setup Instructions

1. **Clone the repo**
```bash
git clone <your-repo-url>
cd alarmservice
```

2. **Install dependencies**
```bash
go mod tidy
```

3. **Run the service**
```bash
go run main.go
```

Service runs at `http://localhost:8080`

## 📮 API Endpoints

### 🔸 Create Alarm
```bash
curl -X POST http://localhost:8080/alarm \
  -H "Content-Type: application/json" \
  -d '{
    "name": "High Memory",
    "condition": "mem > 90"
}'
```

### 🔸 Get All Alarms
```bash
curl http://localhost:8080/alarm
```

### 🔸 Get Alarm by ID
```bash
curl http://localhost:8080/alarm/<id>
```

### 🔸 Update Alarm
```bash
curl -X PUT http://localhost:8080/alarm/<id> \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Alarm",
    "condition": "cpu > 95",
    "state": "ACKED"
}'
```

### 🔸 Delete Alarm
```bash
curl -X DELETE http://localhost:8080/alarm/<id>
```

## 🧪 Run Unit Tests
```bash
go test ./tests -v
```

## ✨ Features
- RESTful CRUD APIs
- Periodic Notification Checker
- ACK support to reduce notification frequency
- In-memory store
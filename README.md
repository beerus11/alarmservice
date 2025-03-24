# Alarm Microservice (Go + Gin)

This is a simple microservice built using Go and the Gin framework. It manages alarms with basic CRUD operations and notifies if certain conditions are met. The data is stored in memory (no database).

## How to Run

1. Make sure you have Go installed (version 1.18 or above).
2. Clone the repo and navigate to the folder.
3. Run the following commands:

```bash
go mod tidy
go run main.go
```

The server will start at `http://localhost:8080`.

## Available APIs

### Create an Alarm
```bash
curl -X POST http://localhost:8080/alarm \
  -H "Content-Type: application/json" \
  -d '{"name": "High CPU", "condition": "cpu > 90"}'
```

### Get All Alarms
```bash
curl http://localhost:8080/alarm
```

### Get Alarm by ID
```bash
curl http://localhost:8080/alarm/<alarm-id>
```

### Update an Alarm
```bash
curl -X PUT http://localhost:8080/alarm/<alarm-id> \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Alarm", "condition": "mem > 95", "state": "ACKED"}'
```

### Delete an Alarm
```bash
curl -X DELETE http://localhost:8080/alarm/<alarm-id>
```

## Running Tests

Run the following command to execute unit tests:

```bash
go test ./tests -v
```

---

No database, no UI. Just a clean API to manage alarms. Built for simplicity and easy testing. 😊
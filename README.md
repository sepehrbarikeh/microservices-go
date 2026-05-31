🚀 Microservices Backend System (Go + gRPC + RabbitMQ)
📌 Overview

This project is a microservices-based backend system built with Go.
It demonstrates real-world backend architecture using:

gRPC for service-to-service communication
RabbitMQ for asynchronous messaging
PostgreSQL for data persistence
JWT authentication
Clean separation of microservices

🏗️ Architecture
                +------------------+
                |   Auth Service   |
                | (gRPC + JWT)     |
                +--------+---------+
                         |
                         | gRPC
                         ↓
+------------------+   RabbitMQ   +----------------------+
|  Order Service   | -----------> | Notification Service |
| (REST API)       |              | (Consumer / Logger)  |
+------------------+              +----------------------+
         |
         |
         ↓
   PostgreSQL

⚙️ Services
🔐 Auth Service
User registration & login
JWT token generation & validation
gRPC server for user verification
📦 Order Service
Create & manage orders (REST API using Fiber)
Calls Auth Service via gRPC to validate users
Publishes events to RabbitMQ
📢 Notification Service
RabbitMQ consumer
Receives order events
Logs messages (can be extended to email/SMS system)
Implements retry mechanism + DLQ concept
📨 Messaging System (RabbitMQ)

This project uses RabbitMQ for async communication:

Features:
Queue-based messaging
Retry mechanism (basic backoff)
Dead Letter Queue (DLQ)
Event-driven architecture

Flow:
Order Created
     ↓
RabbitMQ Queue
     ↓
Notification Service
     ↓
Log / Process Event

🔁 Retry & Failure Handling
Failed messages are retried up to 3 times
After max retries → message sent to Dead Letter Queue
Prevents infinite retry loops
🧠 Key Concepts Implemented
Microservices architecture
gRPC communication
Event-driven design
Message broker (RabbitMQ)
Retry mechanism
Dead Letter Queue (DLQ)
JWT authentication middleware
🛠️ Tech Stack
Go (Golang)
Fiber (HTTP framework)
gRPC
PostgreSQL
RabbitMQ
Docker (optional future improvement)
📂 Project Structure
auth-service/
order-service/
notification-service/

internal/
  ├── grpc
  ├── handler
  ├── service
  ├── repository
  ├── rabbitmq
  └── middleware
🚀 How to Run
1. Start RabbitMQ
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
2. Start Auth Service
go run cmd/main.go
go run cmd/grpc/main.go
3. Start Order Service
go run cmd/main.go
4. Start Notification Service
go run cmd/main.go
📌 Future Improvements
Redis-based idempotency
Distributed tracing (OpenTelemetry)
Advanced retry strategy (exponential backoff)
Docker Compose setup
Kubernetes deployment
Structured logging (Zap / Logrus)
🎯 What This Project Shows

This project demonstrates:

Real-world backend architecture design
Microservices communication patterns
Event-driven systems
Fault tolerance basics
Production-like backend structure
👨‍💻 Author

Built for learning and showcasing backend engineering skills using Go.
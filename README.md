# 🚀 Microservices Backend System (Go + gRPC + RabbitMQ)

## 📌 Overview

This project is a microservices-based backend system built with Go.  
It demonstrates real-world backend architecture using modern distributed systems concepts.

The system includes:

- Authentication Service (gRPC + JWT)
- Order Service (REST API)
- Notification Service (RabbitMQ Consumer)
- PostgreSQL Database
- RabbitMQ Message Broker
- Retry mechanism + Dead Letter Queue (DLQ)
- Idempotent consumer (duplicate protection)

---

## 🏗️ Architecture

Auth Service
    |
    v (gRPC)
Order Service -----> RabbitMQ -----> Notification Service
    |
    v
PostgreSQL
---

## ⚙️ Services

### 🔐 Auth Service
- User registration & login
- JWT authentication
- gRPC API for user validation

---

### 📦 Order Service
- Create and manage orders (REST API)
- Validates users via Auth Service (gRPC)
- Publishes events to RabbitMQ

---

### 📢 Notification Service
- Consumes messages from RabbitMQ
- Logs incoming events
- Implements:
  - Retry mechanism
  - Dead Letter Queue (DLQ)
  - Idempotency (duplicate message protection)

---

## 📨 Messaging System (RabbitMQ)

### Flow

Order Created
↓
orders_queue
↓
Notification Service
↓
Process Event
↓
Retry (if fail)
↓
DLQ (if max retries exceeded)



---

## 🔁 Retry & Failure Handling

- Failed messages are retried up to **3 times**
- Retry queue adds delay between attempts
- After max retries → message is sent to **Dead Letter Queue**
- Prevents infinite retry loops

---

## 🧠 Key Concepts Implemented

- Microservices architecture
- gRPC communication
- Event-driven design
- Message broker (RabbitMQ)
- At-least-once delivery model
- Retry mechanism
- Dead Letter Queue (DLQ)
- Idempotent consumer design

---

## 🛠️ Tech Stack

- Go (Golang)
- Fiber (HTTP framework)
- gRPC
- PostgreSQL
- RabbitMQ

---

## 📂 Project Structure
auth-service/
order-service/
notification-service/

internal/
├── handler
├── service
├── repository
├── grpc
├── rabbitmq
└── middleware


---

## 🚀 How to Run

### 1. Start RabbitMQdocker run -d --name rabbitmq
-p 5672:5672 -p 15672:15672
rabbitmq:management

### 2. Run Auth Service
go run cmd/main.go
go run cmd/grpc/main.go

### 3. Run Order Service
go run cmd/main.go

### 4. Run Notification Service
go run cmd/main.go

---

## 📌 Features Highlights

- Clean microservices separation
- Secure authentication using JWT
- gRPC internal communication
- Asynchronous event processing
- Retry + DLQ handling
- Duplicate-safe message processing

---

## 🎯 What This Project Demonstrates

This project shows understanding of:

- Distributed systems fundamentals
- Event-driven architecture
- Fault tolerance patterns
- Backend system design
- Real-world microservices communication

---

## 👨‍💻 Author

Built for learning and backend engineering portfolio.

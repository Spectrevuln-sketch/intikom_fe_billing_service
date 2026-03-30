# Billing Service

## Overview

This project is a simple event-driven billing system using Go, Kafka, and PostgreSQL.

Transactions are processed asynchronously using Kafka to improve scalability and reliability.

---

## Architecture

Client → API (Gin) → Kafka → Worker → PostgreSQL

---

## How to Run

### 1. Run Kafka

```bash
docker-compose up -d
```

### 2. Run Backend

```bash
go run server/main.go
```

### 3. Run Frontend

```bash
cd frontend
npm install
npm run dev
```

---

## API

### Create Transaction

POST /transactions

### Get Transactions

GET /accounts/:id/transactions?page=1&limit=5

---

## Notes

* Uses Kafka for asynchronous processing
* Uses database transaction with row-level locking
* Simple polling is used for frontend updates

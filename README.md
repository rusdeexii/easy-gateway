# 🚀 Easy-GateWay - Payment Gateway Platform

A modern **microservices-based payment gateway** system with full-stack security, real-time communication, role-based access control, and admin dashboard. Built using **Go (Gin + GORM)**, **PostgreSQL**, **RabbitMQ**, and **Vue 3 + TypeScript + TailwindCSS**.

---

## 📦 Features

### ✅ Authentication & Authorization
- JWT-based login (Access & Refresh tokens)
- HttpOnly cookies for enhanced frontend security
- Role-based access control (`admin`, `superadmin`)
- Token rotation & session logout

### 🧱 Microservices Architecture
- **`auth-service`** – Handles login, registration, user management  
- **`callback-service`** – Manages webhook callbacks (SCB, KBANK, TTB)
- **`agent-service`** – Dispatches messages to message queue  
- **`statement-service`** – Aggregates and formats transaction history  
- **`merchant-service`** – Manages merchant registration & webhooks  
- **`notification-service`** – Sends emails/messages notifications  
- Uses **RabbitMQ** for service communication  
- Each service uses **PostgreSQL** for data persistence  

### 🌐 Frontend (Vue 3 + Vite)
- Secure login and dashboard with **Pinia** state management
- Route guard and token validation per page
- Admin UI for managing users and merchants
- Styled with **TailwindCSS** for a modern responsive look

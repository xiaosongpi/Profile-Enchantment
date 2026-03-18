# Profile-Enchantment

## 🎯 Project Purpose
A robust backend service built with **Golang**, designed as a personal laboratory for mastering high-performance web development and clean architectural patterns.

## 🛠 Tech Stack
* **Language:** Go (Golang)
* **Web Framework:** [Fiber v2](https://gofiber.io/)
* **Database:** Integrated via `GORM`
* **Configuration:** Environment-based via `.env`

## 🚀 Repo Architecture
This project follows a modular structure to ensure separation of concerns and maintainability:
```text
.
├── app/
│   ├── cmd/
│   │   └── main.go
│   ├── config/
│   └── internal/
│       └── auth/
│           ├── domain/
│           │   └── user_domain.go
│           ├── dto/
│           │   └── user_dto.go
│           ├── handler/
│           ├── repository/
│           ├── usecase/
│           └── notes/
│   └── pkg/
├── .env
├── .env.example
├── .gitignore
├── docker-compose.yaml
├── go.mod
├── go.sum
└── README.md
```

## 📋 Prerequisites
Before you begin, ensure you have the following installed:
* [Go](https://go.dev/dl/) (version 1.18 or higher recommended)
* An active Database instance (PostgreSQL/MySQL)

## ⚙️ Installation & Setup

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/your-username/profile-enchantment.git](https://github.com/your-username/profile-enchantment.git)
   cd profile-enchantment
# ☕ Caffinity

**Caffinity** is a cafe discovery web application that allows users to explore cafes by city, view detailed information, menus, and photo galleries.

This project originated from an idea by the team at **Tryspace.id**, who envisioned a platform dedicated to helping people discover great places to enjoy coffee.
It is built as a fullstack portfolio to demonstrate API design, clean architecture, and integration between Go (REST/GraphQL) backend and React frontend.

---

## 🚀 Features

* 📍 Discover cafes by city
* 📖 View detailed cafe information
* 🍽️ Browse menus
* 🖼️ Photo gallery for each cafe
* ⚡ REST API and GraphQL support
* 🧩 Clean and modular architecture
* 🔧 Environment-based configuration

---

## 🛠️ Tech Stack

* **Backend:** Golang
* **Database:** PostgreSQL
* **Query Builder:** SQLC
* **API:** REST & GraphQL
* **Frontend:** React (planned / integrated)
* **Tools:** Makefile, Docker (optional)

---

## 💡 Project Background

Caffinity was inspired by a simple idea: making it easier for people to find the perfect place to enjoy coffee.

The concept was initiated by the **Tryspace.id** team, aiming to create a platform where users can:

* Discover cafes in different cities
* Explore menus and ambiance
* Find the best spots for working, meeting, or relaxing

This project brings that idea into a technical implementation using modern backend architecture and scalable design.

---

## 📂 Project Structure

```id="4t4cfx"
.
├── cmd/                # Application entry point
├── internal/
│   ├── handler/       # HTTP / GraphQL handlers
│   ├── service/       # Business logic
│   ├── repository/    # Database layer
│   └── model/         # Data structures
├── pkg/               # Utilities and helpers
├── db/                # SQL queries / migrations
├── .env               # Environment variables
├── Makefile           # Command automation
└── README.md
```

---

## ⚙️ Prerequisites

Make sure you have installed:

* Go >= 1.20
* PostgreSQL
* Make

---

## 🔧 Setup Environment

Create a `.env` file (or copy from `.env.example` if available):

```bash id="m6z0xv"
cp .env.example .env
```

Example configuration:

```env id="m6m2dl"
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=caffinity
APP_PORT=8080
```

---

## 🏃 Running the Project

This project uses a **Makefile** for easier command execution.

### ▶️ Run the application

```bash id="jqlb3g"
make run
```

### 🔨 Build the application

```bash id="zsz3j2"
make build
```

### 🧹 Clean build files

```bash id="j0l7js"
make clean
```

---

## 🗄️ Database

### Run migrations (if configured)

```bash id="3sjjfn"
make migrate-up
make migrate-down
```

### Generate SQL queries (SQLC)

```bash id="s0q5i6"
make sqlc
```

---

## 🧪 Testing

```bash id="l8t2az"
make test
```

---

## 📡 API Endpoints

Example endpoints:

```id="l6lcmq"
GET /cafes?city=jakarta
GET /cafes/:id
```

---

## 📦 Deployment

Build the binary:

```bash id="mwk0q3"
make build
```

Run the application:

```bash id="e2o4o9"
./bin/app
```

---

## 🤝 Contributing

Contributions are welcome!

1. Fork the repository
2. Create a new branch
3. Commit your changes
4. Push and open a Pull Request

---

## 📄 License

This project is licensed under the MIT License.

---

## 👨‍💻 Author

Developed by **Teguh Setiawan** 🚀

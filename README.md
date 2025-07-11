# 99answer

📝 A simple project with an API Gateway, Listing service, and User service written in Go and Python, using SQLite as storage.

## 📦 Services

This project consists of 3 services:

- **gateway** — API Gateway (Go + Echo)
- **listingService** — Service to manage listings (Python + Tornado + SQLite)
- **userService** — Service to manage users (Go + Echo + SQLite)

## 🚀 How to Run

You can run the project in two ways depending on your environment.

### 🐳 Using Docker Compose (recommended)

If you prefer not to install Go or Python locally, you can use Docker Compose:

```bash
docker-compose up -d
```

This will build and start all services in detached mode.  

You can check the logs with:

```bash
docker-compose logs -f
```

And shut down the services with:

```bash
docker-compose down
```

### 🧪 Running services manually
If you want to run the services one by one manually, you can:

#### Run the **listingService** 
```bash
cd listingService
python3 listing_service.py
```

#### Run the **userService**
```bash
cd userService
go run main.go
```

#### Run the **gateway**
```bash
cd apiGateway
go run main.go
```

⚠️ **Make sure all dependencies are installed before running!**

### 🔄 Restarting services
If you make changes to the code, restart the affected service manually by stopping and starting it again, or if using Docker Compose:

```bash
docker-compose restart
```

## 📂 Notes

- SQLite database files (`listings.db`, `user.db`) are mounted as volumes and shared between services.
- Make sure ports `9001`, `6000`, and `9002` are free before starting.
- You can use `docker-compose ps` to check service status.

Enjoy! 🚀

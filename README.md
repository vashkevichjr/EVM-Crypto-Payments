# 🚀 Crypto Payment Gateway

B2B microservice in Go for accepting cryptocurrency payments (USDT/USDC on EVM networks).

## 📋 Requirements

- Go 1.22+
- Docker and Docker Compose
- PostgreSQL 15+ (runs via Docker)

## 🚀 Quick Start

1. **Clone the repository** (if you haven't already)

2. **Copy the environment variables file:**
   ```bash
   cp .env.example .env
   ```

3. **Generate an encryption key:**
   ```bash
   openssl rand -hex 32
   ```
   Add the result to the `.env` file as the value of `ENCRYPTION_KEY`

4. **Start PostgreSQL:**
   ```bash
   make docker-up
   ```

5. **Apply migrations** (after Phase 1 implementation):
   ```bash
   make migrate-up
   ```

6. **Run the application:**
   ```bash
   make run
   ```

## 📁 Project Structure

```
.
├── cmd/gateway/          # Application entry point
├── internal/             # Internal packages
│   ├── app/             # Application assembly (DI)
│   ├── config/          # Configuration
│   ├── handlers/        # HTTP controllers
│   ├── models/          # Domain models
│   ├── services/        # Business logic
│   ├── repositories/    # Database layer
│   ├── blockchain/      # Blockchain integration
│   └── workers/         # Background processes
├── pkg/                  # Public packages
│   ├── logger/          # Logging
│   └── ecdsa/           # Cryptography
├── migrations/           # SQL migrations
└── docker-compose.yaml   # Docker configuration
```

## 🛠 Available Commands

See `Makefile` for all available commands:

```bash
make help          # Show help
make run           # Run application
make build         # Build binary
make docker-up     # Start Docker containers
make docker-down   # Stop Docker containers
make migrate-up    # Apply migrations
make migrate-down  # Rollback migrations
```

## 📖 Implementation Plan

Detailed implementation plan is in [MASTER_PLAN.md](./MASTER_PLAN.md).

### Current Status

- ✅ Phase 1: Infrastructure and Database (in progress)
  - ✅ Project structure created
  - ✅ go.mod initialized
  - ✅ docker-compose.yaml created
  - ⏳ PostgreSQL migrations (next step)
  - ⏳ Repositories layer

## 🔒 Security

- Private keys are encrypted before saving to database
- No keys in logs
- Strict transaction amount validation with token decimals consideration

## 📝 License

MIT

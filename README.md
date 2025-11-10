# Workshop 4 - User Management & Points Transfer API

A Go backend API implementation using Clean Architecture principles with Fiber framework for user management and points transfer functionality.

## 🏗️ Architecture

This project follows **Clean Architecture** principles with clear separation of concerns:

```
workshop4/
├── domain/           # Business Logic Layer
│   ├── entities/     # Business entities
│   └── repositories/ # Repository interfaces
├── usecases/         # Application Business Rules
├── interface/        # Interface Adapters
│   └── controllers/  # HTTP Controllers
├── infrastructure/   # Frameworks & Drivers
│   ├── models/       # Database models
│   └── repositories/ # Repository implementations
└── tests/           # Test files
    ├── controllers/
    ├── mocks/
    ├── repositories/
    └── usecases/
```

## 🚀 Features

- ✅ **User Management**: Full CRUD operations for users
- ✅ **Points Transfer**: Transfer points between users with transaction safety
- ✅ **Clean Architecture**: Modular and testable codebase
- ✅ **Database Integration**: SQLite with GORM ORM
- ✅ **Unit Testing**: Comprehensive test suite with mocks
- ✅ **API Documentation**: RESTful API endpoints
- ✅ **Transaction Safety**: ACID compliance for transfer operations
- ✅ **Decimal Precision**: Accurate monetary calculations

## 📋 Requirements

- Go 1.21+
- SQLite3

## 🛠️ Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd workshop4
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:3000`

## 🧪 Testing

### Run all tests:
```bash
make test
```

### Run with coverage:
```bash
make test-coverage
```

### Run specific test suites:
```bash
# Unit tests only
make test-unit

# Integration tests only  
make test-integration
```

## 📚 API Endpoints

### Users
- `GET /users` - Get all users
- `GET /users/{id}` - Get user by ID
- `POST /users` - Create new user
- `PUT /users/{id}` - Update user
- `DELETE /users/{id}` - Delete user

### Transfers
- `POST /users/{id}/transfer` - Transfer points from user
- `GET /users/{id}/transfer` - Get user's transfer history
- `GET /transfer/{id}` - Get transfer by ID

## 📖 API Usage Examples

### Create User
```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{
    "firstName": "John",
    "lastName": "Doe", 
    "email": "john@example.com",
    "phone": "1234567890",
    "dateOfBirth": "1990-01-01",
    "address": "123 Main St",
    "city": "Bangkok",
    "country": "Thailand",
    "postalCode": "10110"
  }'
```

### Transfer Points
```bash
curl -X POST http://localhost:3000/users/1/transfer \
  -H "Content-Type: application/json" \
  -d '{
    "toUserId": 2,
    "amount": 150.25,
    "description": "Payment for service"
  }'
```

### Get Transfer History
```bash
curl "http://localhost:3000/users/1/transfer?limit=10&offset=0"
```

## 🗄️ Database Schema

See [database.md](./database.md) for detailed database schema documentation including ER diagrams.

### Key Tables:
- **users**: User information and points balance
- **transfers**: Points transfer transactions with audit trail

## 🧪 Test Coverage

The project maintains high test coverage across all layers:

- **Domain Layer**: Entity and repository interface tests
- **Use Case Layer**: Business logic tests with mocked dependencies  
- **Interface Layer**: HTTP controller integration tests
- **Infrastructure Layer**: Database repository tests with in-memory SQLite

Current coverage: **90%+**

## 🏃‍♂️ Development

### Project Structure
Following Clean Architecture principles:

1. **Domain Layer** (`domain/`): Contains business entities and repository interfaces
2. **Use Case Layer** (`usecases/`): Application-specific business rules
3. **Interface Layer** (`interface/`): Controllers and external interfaces
4. **Infrastructure Layer** (`infrastructure/`): Database models and implementations

### Code Style Guidelines

- Use `gofmt` for code formatting
- Follow Go naming conventions
- Write meaningful commit messages
- Maintain test coverage above 80%
- Use dependency injection for better testability

### Available Make Commands

```bash
make test              # Run all tests
make test-verbose      # Run tests with verbose output
make test-coverage     # Run tests with coverage report
make test-unit         # Run unit tests only
make test-integration  # Run integration tests only
make clean-test        # Clean test artifacts
```

## 🚀 Deployment

### Build for production:
```bash
go build -o server main.go
```

### Run production server:
```bash
./server
```

## 🔄 Migration

Database migrations are handled automatically by GORM's AutoMigrate feature. The application will create and update database schema on startup.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is part of KBTG AI Workshop training materials.

## 📞 Support

For questions and support, please refer to the workshop materials or contact the instructors.
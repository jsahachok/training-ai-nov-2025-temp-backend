# Workshop 4 - Copilot Instructions

## Project Overview
This is a Go backend API for user management and points transfer functionality built with Clean Architecture principles. The project follows KBTG workshop standards and implements a complete REST API with comprehensive testing.

**Key Features:**
- User management (CRUD operations)
- Points transfer system with transaction safety
- Clean Architecture implementation
- Comprehensive test coverage (90%+)
- SQLite database with GORM ORM
- Fiber web framework

## Architecture Guidelines

### Clean Architecture Layers
Follow the established Clean Architecture pattern with these layers:

1. **Domain Layer** (`domain/`):
   - `entities/`: Business entities (User, Transfer, Response structs)
   - `repositories/`: Repository interfaces only
   
2. **Use Cases Layer** (`usecases/`):
   - Business logic and application rules
   - Orchestrates data flow between layers
   
3. **Interface Layer** (`interface/controllers/`):
   - HTTP controllers handling requests/responses
   - Input validation and response formatting
   
4. **Infrastructure Layer** (`infrastructure/`):
   - Database models and repository implementations
   - External service integrations

### Dependency Direction
Always follow the dependency rule: outer layers depend on inner layers, never the reverse.

## Code Generation Guidelines

### When creating Go structs:
- Use proper JSON tags with camelCase for API responses
- Include validation tags using `github.com/go-playground/validator/v10`
- Follow existing naming patterns (User, Transfer entities)
- Include pointer fields for optional updates (`*string` for UserUpdateRequest)

### Error Handling:
- Use standardized error codes from `internal/constants/errors.go`
- Follow the pattern: `ErrCode[Entity][Condition]` (e.g., `ErrCodeUserNotFound`)
- Always return structured errors with code and message
- Use the response utility from `internal/utils/response.go`

### Database Operations:
- Use GORM patterns consistent with existing models
- Implement soft deletes with `DeletedAt *time.Time`
- Include audit fields: `CreatedAt`, `UpdatedAt`, `DeletedAt`
- Use transactions for operations affecting multiple entities
- Follow snake_case for database columns, camelCase for Go structs

### API Controllers:
- Use Fiber's `c *fiber.Ctx` parameter
- Implement proper HTTP status codes
- Use structured response format with `Success`, `Error`, and `Data` fields
- Include request validation using validator package
- Handle pagination with `limit` and `offset` parameters (defaults: 10, 0)

### Repository Pattern:
- Define interfaces in `domain/repositories/`
- Implement concrete repositories in `infrastructure/repositories/`
- Use dependency injection in constructors
- Return domain entities, not database models
- Handle database errors and convert to domain errors

### Use Cases:
- Contain business logic only
- Depend on repository interfaces, not implementations
- Return domain entities and standardized errors
- Implement transaction handling for complex operations
- Validate business rules (e.g., self-transfer prevention, balance checks)

## Testing Requirements

### Unit Testing:
- Test all use cases with mocked dependencies
- Use `testify/mock` for repository mocks
- Test both success and error scenarios
- Achieve minimum 90% code coverage
- Place tests in `tests/` directory organized by layer

### Integration Testing:
- Test controllers with in-memory database
- Use `testify/suite` for test organization
- Test complete request/response cycles
- Include edge cases and error conditions

### Mock Generation:
- Create mocks in `tests/mocks/` directory
- Follow naming pattern: `[Interface]Mock` (e.g., `UserRepositoryMock`)
- Mock all external dependencies (database, HTTP clients)

## Validation Standards

### Input Validation:
- Use struct tags for basic validation
- Implement custom validators for business rules
- Validate email formats, required fields, and data types
- Return meaningful error messages with field names

### Business Validation:
- Validate points transfer amounts (positive, sufficient balance)
- Prevent self-transfers
- Check user existence before operations
- Validate date formats and ranges

## API Design Patterns

### RESTful Endpoints:
- Use standard HTTP methods (GET, POST, PUT, DELETE)
- Follow resource-based URL patterns
- Return appropriate HTTP status codes
- Include proper error responses

### Request/Response Format:
```go
// Success Response
{
  "success": true,
  "message": "Success",
  "data": {...}
}

// Error Response
{
  "success": false,
  "error": {
    "code": "USER_NOT_FOUND",
    "message": "User not found"
  }
}
```

### Pagination:
- Support `limit` and `offset` query parameters
- Default limit: 10, max limit: 100
- Include total count in responses when applicable

## Database Conventions

### Naming:
- Use snake_case for database columns
- Use camelCase for Go struct fields
- Use PascalCase for Go struct names

### Relationships:
- Use foreign keys with proper constraints
- Implement soft deletes for audit trail
- Include created_at and updated_at timestamps

### Performance:
- Add indexes for frequently queried fields
- Use database transactions for consistency
- Implement connection pooling

## Security Guidelines

### Input Sanitization:
- Validate all user inputs
- Sanitize data before database operations
- Use parameterized queries (GORM handles this)

### Data Protection:
- Never expose sensitive data in logs
- Use environment variables for configuration
- Implement proper error logging without exposing internal details

## Performance Considerations

### Database:
- Use appropriate GORM preloading for relationships
- Implement pagination for large datasets
- Add database indexes for search fields

### API:
- Use compression middleware
- Implement request timeout handling
- Cache frequently accessed data when appropriate

## Development Workflow

### Code Organization:
- One entity per file in appropriate layer
- Use meaningful package and file names
- Group related functionality together

### Testing Workflow:
- Run tests with `make test`
- Generate coverage reports with `make test-coverage`
- Maintain test-driven development approach

### Build Process:
- Use `go mod` for dependency management
- Follow semantic versioning for releases
- Use Makefile for common tasks

## Specific Patterns to Follow

### User Management:
- Support full CRUD operations
- Include comprehensive user profile fields
- Implement soft delete functionality
- Track points balance with decimal precision

### Points Transfer:
- Ensure transaction atomicity
- Validate business rules (balance, self-transfer)
- Maintain audit trail of all transfers
- Support transfer history queries

### Error Handling:
- Use consistent error codes from constants
- Provide meaningful error messages
- Log errors appropriately without exposing sensitive data
- Return structured error responses
# Database Schema Documentation

## Overview
This document describes the database schema for the Workshop4 Backend API, which implements a user management system with points transfer functionality using Clean Architecture principles.

## Entity Relationship Diagram

```mermaid
erDiagram
    USERS {
        uint id PK "Primary Key, Auto Increment"
        varchar first_name "User's first name"
        varchar last_name "User's last name"
        varchar email UK "Unique email address"
        varchar phone "Phone number"
        date date_of_birth "Date of birth"
        text address "Full address"
        varchar city "City name"
        varchar country "Country name"
        varchar postal_code "Postal/ZIP code"
        varchar avatar "Avatar image URL"
        decimal points "User points balance"
        timestamp created_at "Record creation time"
        timestamp updated_at "Last update time"
        timestamp deleted_at "Soft delete timestamp"
    }

    TRANSFERS {
        uint id PK "Primary Key, Auto Increment"
        uint from_user_id FK "Source user ID"
        uint to_user_id FK "Destination user ID"
        decimal amount "Transfer amount"
        varchar description "Transfer description"
        varchar status "Transfer status"
        timestamp transferred_at "Transfer execution time"
        timestamp created_at "Record creation time"
        timestamp updated_at "Last update time"
        timestamp deleted_at "Soft delete timestamp"
    }

    USERS ||--o{ TRANSFERS : "from_user_id"
    USERS ||--o{ TRANSFERS : "to_user_id"
```

## Table Details

### Users Table
The `users` table stores all user information and their current points balance.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique identifier for each user |
| first_name | varchar(100) | NOT NULL | User's first name |
| last_name | varchar(100) | NOT NULL | User's last name |
| email | varchar(255) | NOT NULL, UNIQUE | User's email address (used for login) |
| phone | varchar(20) | NOT NULL | User's phone number |
| date_of_birth | date | NOT NULL | User's date of birth |
| address | text | NOT NULL | User's full address |
| city | varchar(100) | NOT NULL | User's city |
| country | varchar(100) | NOT NULL | User's country |
| postal_code | varchar(20) | NOT NULL | User's postal/ZIP code |
| avatar | varchar(500) | DEFAULT '' | URL to user's avatar image |
| points | decimal(10,2) | NOT NULL, DEFAULT 0.00 | User's current points balance |
| created_at | timestamp | NOT NULL | When the record was created |
| updated_at | timestamp | NOT NULL | When the record was last updated |
| deleted_at | timestamp | INDEX | Soft delete timestamp (NULL if not deleted) |

#### Indexes
- PRIMARY KEY on `id`
- UNIQUE INDEX on `email`
- INDEX on `deleted_at` (for soft delete queries)

### Transfers Table
The `transfers` table records all point transfer transactions between users.

| Column | Type | Constraints | Description |
|--------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique identifier for each transfer |
| from_user_id | uint | NOT NULL, FOREIGN KEY, INDEX | ID of the user sending points |
| to_user_id | uint | NOT NULL, FOREIGN KEY, INDEX | ID of the user receiving points |
| amount | decimal(10,2) | NOT NULL | Amount of points transferred |
| description | varchar(255) | DEFAULT '' | Optional description of the transfer |
| status | varchar(50) | NOT NULL, DEFAULT 'pending' | Transfer status (pending, completed, failed) |
| transferred_at | timestamp | NOT NULL | When the transfer was executed |
| created_at | timestamp | NOT NULL | When the record was created |
| updated_at | timestamp | NOT NULL | When the record was last updated |
| deleted_at | timestamp | INDEX | Soft delete timestamp (NULL if not deleted) |

#### Indexes
- PRIMARY KEY on `id`
- INDEX on `from_user_id`
- INDEX on `to_user_id`
- INDEX on `deleted_at` (for soft delete queries)

#### Foreign Key Constraints
- `from_user_id` REFERENCES `users(id)` ON UPDATE CASCADE ON DELETE RESTRICT
- `to_user_id` REFERENCES `users(id)` ON UPDATE CASCADE ON DELETE RESTRICT

## SQL Schema

### Create Users Table
```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL,
    date_of_birth DATE NOT NULL,
    address TEXT NOT NULL,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    avatar VARCHAR(500) DEFAULT '',
    points DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Indexes
CREATE UNIQUE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_deleted_at ON users(deleted_at);
```

### Create Transfers Table
```sql
CREATE TABLE transfers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    from_user_id INTEGER NOT NULL,
    to_user_id INTEGER NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    description VARCHAR(255) DEFAULT '',
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    transferred_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (from_user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    FOREIGN KEY (to_user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

-- Indexes
CREATE INDEX idx_transfers_from_user_id ON transfers(from_user_id);
CREATE INDEX idx_transfers_to_user_id ON transfers(to_user_id);
CREATE INDEX idx_transfers_deleted_at ON transfers(deleted_at);
```

## Business Rules

### Transfer Rules
1. **Self-Transfer Prevention**: Users cannot transfer points to themselves
2. **Sufficient Balance**: Sender must have enough points for the transfer
3. **User Existence**: Both sender and receiver must exist in the system
4. **Transaction Integrity**: All transfers use database transactions (ACID compliance)
5. **Minimum Amount**: Transfer amount must be greater than 0.01
6. **Decimal Precision**: Supports up to 2 decimal places for precise monetary values

### Data Integrity
1. **Soft Delete**: Records are never physically deleted, only marked with `deleted_at`
2. **Audit Trail**: All records maintain creation and update timestamps
3. **Email Uniqueness**: Each email can only be associated with one active user
4. **Foreign Key Constraints**: Ensure referential integrity between users and transfers

## Status Values

### Transfer Status
- `pending`: Transfer has been initiated but not yet processed
- `completed`: Transfer has been successfully processed
- `failed`: Transfer failed due to validation or system errors

## Database Engine
- **SQLite**: Used for development and testing
- **GORM**: ORM for Go with automatic migrations
- **Transaction Support**: Full ACID transaction support for transfer operations

## Migration Notes
The database schema is managed by GORM's AutoMigrate feature, which automatically creates and updates tables based on the model definitions.

Initial migration creates:
1. `users` table with all fields and constraints
2. `transfers` table with foreign key relationships
3. Appropriate indexes for performance optimization

## Performance Considerations
- Indexes on foreign keys (`from_user_id`, `to_user_id`) for fast transfer lookups
- Unique index on email for fast user authentication
- Soft delete indexes for efficient filtering of active records

## Example Queries

### Get User with Points Balance
```sql
SELECT id, first_name, last_name, email, points 
FROM users 
WHERE deleted_at IS NULL 
AND id = ?;
```

### Get Transfer History for User
```sql
SELECT t.*, 
       u1.first_name as sender_first_name, 
       u1.last_name as sender_last_name,
       u2.first_name as receiver_first_name, 
       u2.last_name as receiver_last_name
FROM transfers t
JOIN users u1 ON t.from_user_id = u1.id
JOIN users u2 ON t.to_user_id = u2.id
WHERE (t.from_user_id = ? OR t.to_user_id = ?)
AND t.deleted_at IS NULL
ORDER BY t.transferred_at DESC
LIMIT ? OFFSET ?;
```

### Update User Points (Transaction)
```sql
BEGIN TRANSACTION;

-- Deduct points from sender
UPDATE users 
SET points = points - ?, updated_at = CURRENT_TIMESTAMP 
WHERE id = ? AND points >= ?;

-- Add points to receiver
UPDATE users 
SET points = points + ?, updated_at = CURRENT_TIMESTAMP 
WHERE id = ?;

-- Create transfer record
INSERT INTO transfers (from_user_id, to_user_id, amount, description, status, transferred_at)
VALUES (?, ?, ?, ?, 'completed', CURRENT_TIMESTAMP);

COMMIT;
```
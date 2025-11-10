# Database Schema Documentation

## Overview
This document describes the database schema for the Workshop4 Backend API, which implements a user management system with points transfer functionality using Clean Architecture principles.

## Entity Relationship Diagram

```mermaid
erDiagram
    USERS {
        uint id PK "Primary Key, Auto Increment"
        varchar(100) first_name "User's first name"
        varchar(100) last_name "User's last name"
        varchar(255) email UK "Unique email address"
        varchar(20) phone "Phone number"
        date date_of_birth "Date of birth"
        text address "Full address"
        varchar(100) city "City name"
        varchar(100) country "Country name"
        varchar(20) postal_code "Postal/ZIP code"
        varchar(500) avatar "Avatar image URL"
        decimal(10,2) points "User points balance"
        timestamp created_at "Record creation time"
        timestamp updated_at "Last update time"
        timestamp deleted_at "Soft delete timestamp"
    }

    TRANSFERS {
        uint id PK "Primary Key, Auto Increment"
        uint from_user_id FK "Source user ID"
        uint to_user_id FK "Destination user ID"
        decimal(10,2) amount "Transfer amount"
        varchar(255) description "Transfer description"
        varchar(50) status "Transfer status"
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
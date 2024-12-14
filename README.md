# DBMiner Multi

A post-exploitation tool designed to scan local database instances for sensitive data patterns including credit card details, SSNs, and other compliance-critical PII data.

## Features

- Interface-driven design for consistent behavior across different database backends
- Currently supports:
  - MongoDB
  - MySQL
- Pattern matching for sensitive data fields including:
  - Social Security Numbers
  - Credit Card Numbers
  - Passwords
  - Security/Authentication data
  - Address information
  - PII data fields

## Installation

Requires Go 1.23.0 or higher.

## Usage

### MongoDB Scanner

Example:
```bash
go run ./mongo/main.go localhost
```

### MySQL Scanner

Example:
```bash
go run ./mysql/main.go localhost
```

## Architecture

The project uses a modular architecture based on Go interfaces to ensure consistent behavior across different database implementations:

1. Core Interface (`dbminer.DatabaseMiner`):
   - Defines common schema inspection methods
   - Provides unified search functionality
   - Implements regex-based pattern matching

2. Database-Specific Implementations:
   - Each database type implements the DatabaseMiner interface
   - Handles connection management
   - Provides schema extraction logic specific to the database


## Pattern Detection

The scanner looks for common field names that might contain sensitive data:

- Social Security related fields
- Password/security fields
- Credit card related data (numbers, CVV, expiration)
- Address information (city, state, zip)

## Development

To add support for a new database type:

1. Create a new package for the database
2. Implement the DatabaseMiner interface
3. Provide database-specific connection and schema extraction logic

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## Security Notice

This tool is intended for authorized security testing and educational purposes only. Unauthorized database scanning or data extraction may violate applicable laws and regulations.

## Disclaimer

This tool is provided for educational and authorized testing purposes only. Users are responsible for ensuring compliance with applicable laws and regulations when using this tool.

## License

This project is licensed under the GNU General Public License v3.0 - see the LICENSE file for details.

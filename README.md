# Go - The Complete Guide

Welcome to the Go programming course repository! This project contains all the code examples and exercises from the Udemy course [Go - The Complete Guide](https://www.udemy.com/course/go-the-complete-guide/?couponCode=PLOYALTY0923).

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (latest version recommended)
- A code editor or IDE of your choice

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/go-complete-guide.git
   cd go-complete-guide
   ```

2. Initialize the Go module:
   ```bash
   go mod init example.com/first-app
   ```

3. Download dependencies:
   ```bash
   go mod tidy
   ```

## 📁 Project Structure

```
├── Section-1/          # Basic Go concepts and first applications
│   ├── app.go          # Main application file
│   ├── go.mod          # Go module definition
│   └── first-app       # Compiled binary (if built)
└── Section-2/          # Advanced topics and projects
    ├── investment.go   # Investment calculator example
    └── go.mod          # Go module definition
```

## 🏃‍♂️ Running the Code

### Method 1: Direct Execution
Run Go files directly without building:
```bash
go run Section-1/app.go
go run Section-2/investment.go
```

### Method 2: Build and Execute
Build the application and run the compiled binary:
```bash
# Build the application
go build

# Run the compiled binary
./first-app
```

### Method 3: Run Specific Sections
```bash
# Run Section 1 examples
cd Section-1
go run app.go

# Run Section 2 examples
cd Section-2
go run investment.go
```

## 📚 Course Sections

### Section 1: Getting Started with Go
- Basic syntax and structure
- Variables and data types
- Control flow statements
- Functions and packages

### Section 2: Intermediate Go Concepts
- Structs and interfaces
- Error handling
- Working with files
- Investment calculator project

## 🛠️ Development

### Adding New Examples
1. Create a new directory for your section
2. Initialize a Go module: `go mod init your-module-name`
3. Create your `.go` files
4. Add dependencies to `go.mod` as needed

## 📖 Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)

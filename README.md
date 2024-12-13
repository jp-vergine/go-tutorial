
# **Go Tutorial Project**

This project demonstrates building a Go-based web server with a hot-reloading development workflow, code linting, and structured project organization. It also incorporates essential practices like creating reusable modules and maintaining documentation.

---

## **Project Structure**

The project follows a modular structure for better maintainability and scalability:

```
- cmd           # Entry point for the application
- config        # Configuration files and settings
- user          # User-related logic (e.g., models, handlers)
- itinerary     # Itinerary management (e.g., models, routes)
- destination   # Destination-related features
- search        # Search functionalities (e.g., filters, queries)
```

---

## **Installation**

### **Prerequisites**
- [Go](https://golang.org/) 1.20+
- [golangci-lint](https://github.com/golangci/golangci-lint)
- [Air](https://github.com/cosmtrek/air) (hot reload)

### **Clone the Repository**
```bash
git clone https://github.com/jpvergine/go-tutorial.git
cd go-tutorial
```

---

## **Commands**

### **Development**

#### Hot Reload Setup
**Command:**
```bash
air
```
**Why?**
- Starts the development server with automatic reload on file changes for a faster development workflow.

#### Dev Server Command
**Command:**
```bash
make dev
```
**Why?**
- Uses the `Makefile` to standardize the development workflow. Running `make dev` executes the `air` command seamlessly.

---

### **Dependency Management**

#### Clean Dependencies
**Command:**
```bash
go mod tidy
```
**Why?**
- Cleans unused dependencies and ensures the `go.mod` file is accurate.

---

### **Linting**

#### Run Linters
**Command:**
```bash
make lint
```
**Why?**
- Executes `golangci-lint` to ensure code follows best practices and maintains consistency.

#### Fix Formatting Issues
**Command:**
```bash
golangci-lint run --fix
```
**Why?**
- Automatically resolves simple formatting issues flagged by the linter.

---

### **Build**

#### Build the Application
**Command:**
```bash
make build
```
**Why?**
- Compiles the application into a binary executable (`bin/app`).

---

## **Design Patterns**

This project follows key design patterns to ensure modularity and scalability:

1. **Separation of Concerns**:
   - Each package (`user`, `itinerary`, etc.) encapsulates a specific domain or functionality.
   - The `cmd` folder contains the main entry point of the application.

2. **Config Management**:
   - Configuration is centralized in the `config` package to ensure all components share consistent settings.

3. **Reusable Modules**:
   - Logic for features like `user`, `itinerary`, and `search` is encapsulated in individual packages, promoting code reuse.

---

## **How to Use**

1. Start the development server:
   ```bash
   make dev
   ```
   Visit the app at [http://localhost:8080](http://localhost:8080).

2. Make changes to the code and observe the live reload.

3. Clean up the dependencies:
   ```bash
   make tidy
   ```

4. Lint the code:
   ```bash
   make lint
   ```

5. Build the project for production:
   ```bash
   make build
   ```

---

## **Conceptual Documentation**

### **Design Choices**

1. **Hot Reload (`air`)**:
   - Chosen for its ease of use and ability to speed up development by avoiding manual restarts.

2. **Code Linting (`golangci-lint`)**:
   - Enforces consistent formatting and catches potential errors early.

3. **Makefile**:
   - Standardizes project commands, making it easier for team members to follow a consistent workflow.

4. **Project Structure**:
   - Each package corresponds to a specific feature or domain. This structure ensures clean separation of concerns and makes the codebase easier to navigate and scale.

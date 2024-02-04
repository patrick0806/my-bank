# My Bank API

### Project status: In Development

This is a API for stude about golang, and concourrence

## API Development Process

This Api uses golang 

To running this project in dev mode you need follow this steps

### 1 - install dependencies
```bash
    go mod tidy
```

### 2 - Create your dot env file with this structure
```bash
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=my-bank
DB_USER=root
DB_PASSWORD=123
API_PORT=:8080
```

### 3 - Create and running local database
```bash
docker run --name my-bank-postgres \
    -p 5432:5432 \
    -e POSTGRES_DB=my-bank \
    -e POSTGRES_USER=root \
    -e POSTGRES_PASSWORD=123 \
    -d postgres:14.4-alpine
```

### 4 - Running migrations <br>
 - Install migrate cli [see how.](https://github.com/golang-migrate/migrate)
 - Install make command [see how.](https://pt.linux-console.net/?p=14595#:~:text=Como%20instalar%20e%20usar%20o%20Make%20no%20Ubuntu,de%20comando%20make%2C%20execute%20o%20comando%3A%20%24%20make--version)
 - running migrations : 
 ```bash
    make migrate
 ```
 - create migration : 
 ```bash
    make createmigration
 ```
  - drop migrations : 
 ```bash
    make migratedown
 ```

### 5- Running the API 
```bash
    go run main.go -env=DEV
```

## Reflection

I'm creating this project to learn more about go.

There are some requirements for things I want to implement that I consider basic in any API, such as documentation (swagger), tests and logs

I'm also going to do my best to do everything with just the golang standard library, because I want to master the stack before delving into libs that the community uses a lot like GIN.

### Swagger route
http://localhost:8080/docs/index.html
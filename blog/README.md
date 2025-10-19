# Blog

## Run Locally

1. Make sure you have Go installed.
2. Create a `.env` file with your postgres database credentials:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb
GIN_MODE=release
PORT=8080
````

3. Run the application:

```bash
go run main.go
```

4. Open your browser at:

```
http://localhost:8080/
```



## Run with Docker

1. Build Docker image:

```bash
docker build -t blog:latest .
```

2. Run Docker container with `.env` file:

```bash
docker run -d --env-file .env -p 8080:8080 --name blog blog:latest
```

3. Open your browser at:

```
http://localhost:8080/
```


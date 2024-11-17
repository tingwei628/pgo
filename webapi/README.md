### webapi

upgrade to go 1.23
```sql
go mod tidy -go=1.23
```


IDE: GoLand \
Shortcuts: \
```
Alt + Fn + Insert // Generate...

```

Tests:
```
based on "Generate" shortcut to generate tests
```


DB schema:
```sql

CREATE TABLE todo_items (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

```


packages
```
authorization
middleware
jwt

sql fuzz test
```


swagger
```
swag init --parseDependency  -d cmd,internal/transport
```
https://github.com/swaggo/swag?tab=readme-ov-file#swag-cli


Generate a avatar picture(.png)

- Usage
```
go run main.go -inputs=abcde
``` 
"abcde" will be hashed, then generates a png file
```
go run main.go -h // check how to use
```

**main.go**
```go
package main
import (
  "github.com/tingwei628/pgo/avatarme"
)
func main () {
    avatarme := avatarme.Avatarme{}
    avatarme.Generate() // it generates .png file
}
```

- Reference
[Tutorial: Identicon generator in Go](https://bartfokker.com/posts/identicon/)
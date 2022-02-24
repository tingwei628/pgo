Speak out the system information in macOS

- Usage

```
go run main.go
```

**main.go**
```go
package main
import (
  "github.com/tingwei628/pgo/sysspeak"
)
func main () {
    rc := sysspeak.RC{}
    rc.Command()
}
```

- Reference
[gopsutil](https://github.com/shirou/gopsutil)
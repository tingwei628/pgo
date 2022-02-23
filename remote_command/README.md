Speak out the system information `e.g. Current UTC time, CPU usage, Available RAM` in macOS

- Usage

```
go run main.go
```

**main.go**
```go
package main
import (
  "github.com/tingwei628/pgo/remote_command"
)
func main () {
    rc := remote_command.RC{}
    rc.Command()
}
```

- Reference
[gopsutil](https://github.com/shirou/gopsutil)
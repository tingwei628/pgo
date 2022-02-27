The payment test with a Telegram bot

- [Demo](https://youtube.com/shorts/D3oFsNR9u7Y?feature=share)

- Usage

```
go run main.go
```

**main.go**
```go
package main
import (
  "github.com/tingwei628/pgo/tgbotpay"
)
func main () {
    tgbotpay.PayStart()
}
```

- Reference
[Telegram bot](https://core.telegram.org/bots) \
[Telegram bot payments](https://core.telegram.org/bots/api#payments) \
[Stripe Tesing](https://stripe.com/docs/testing#cards)
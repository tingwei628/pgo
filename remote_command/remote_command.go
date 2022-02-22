package remote_command

import (
	"fmt"
)

/*
系統資訊
-t, Current UTC time
-cu, Current CPU usage
-ar, Available RAM
-culh, CPU usage over last hour
-arlh,  Available RAM over last hour

-s=xxxx, Make computer "say" something

用 espeak 功能 -> 把系統資訊 唸出來

*/
type RC struct {
	inputs              string
	hashAlgPtr, sizePtr *int
}

func init() {
	fmt.Println("Hello")
}

func (rc *RC) Command() int32 {
	return 100
}

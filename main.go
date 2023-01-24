package main

import "github.com/tingwei628/pgo/webapi"

//"github.com/tingwei628/pgo/avatarme"
//"github.com/tingwei628/pgo/sysspeak"
//"github.com/tingwei628/pgo/tgbotpay"
//"github.com/tingwei628/pgo/webapi"

func main() {

	//tgbotpay.PayStart()
	// rc := sysspeak.RC{}
	// rc.Command()
	//avatarme := avatarme.Avatarme{}
	//avatarme.Generate() // it generates .png file

	api := webapi.API{}
	api.Run()
	// sigs := make(chan os.Signal, 1)
	// signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// done := make(chan bool, 1)
	// go func() {
	// 	// This goroutine executes a blocking receive for
	// 	// signals. When it gets one it'll print it out
	// 	// and then notify the program that it can finish.
	// 	sig := <-sigs
	// 	fmt.Println()
	// 	fmt.Println(sig)
	// 	done <- true
	// }()
	// fmt.Println("awaiting signal")
	// <-done
	// fmt.Println("exiting")
}

package internal

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type RC struct {
	// inputs              string
	// hashAlgPtr, sizePtr *int
}

func init() {
	//fmt.Println("init RC")
}

func (rc *RC) Command() {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("get total mem info failed, err:%v", err)
	}
	// bytes to gigabytes
	memAvailableStr := fmt.Sprintf("Available RAM is %.3f GB \n", float64(memInfo.Available)/(1<<30))

	// cpuInfos, err := cpu.Info()
	// if err != nil {
	// 	fmt.Printf("get cpu info failed, err:%v", err)
	// }
	// for _, info := range cpuInfos {
	// 	data, _ := json.MarshalIndent(info, "", " ")
	// 	fmt.Print(string(data))
	// }

	// Current UTC time
	// "2006-01-02 15:04:05" (go的誕生時間)
	timeStr := fmt.Sprintf("Current UTC time is %v \n", time.Now().UTC().Format("2006/01/02 15:04:05"))
	// 2022/02/23 06:37:01
	//fmt.Printf("Current UTC time %v\n", time.Now().UTC().Format(time.UnixDate))
	// Wed Feb 23 06:39:28 UTC 2022

	// CPU usage
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatalf("get total cpu percent failed, err:%v", err)
	}
	percentStr := fmt.Sprintf("CPU total percent is %.2f%% \n", percent[0])
	// CPU usage over last hour

	//fmt.Print(timeStr)
	//fmt.Print(percentStr)
	//fmt.Print(memAvailableStr)

	// var strBuilder strings.Builder
	// strBuilder.WriteString(timeStr)
	// strBuilder.WriteString(percentStr)
	// strBuilder.WriteString(memAvailableStr)
	// fmt.Print(strBuilder.String())

	for _, sentence := range []string{timeStr, percentStr, memAvailableStr} {
		cmd := exec.Command("say", "-v", "fred", sentence)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

package main

import (
	"github.com/gookit/color"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// print the logo
	color.Green.Print(" ____                    _             ____                           \r\n" +
		"|  _ \\ _   _ _ __  _ __ (_)_ __   __ _/ ___|  ___ _ ____   _____ _ __ \r\n" +
		"| |_) | | | | '_ \\| '_ \\| | '_ \\ / _` \\___ \\ / _ \\ '__\\ \\ / / _ \\ '__|\r\n" +
		"|  _ <| |_| | | | | | | | | | | | (_| |___) |  __/ |   \\ V /  __/ |   \r\n" +
		"|_| \\_\\\\__,_|_| |_|_| |_|_|_| |_|\\__, |____/ \\___|_|    \\_/ \\___|_|   \r\n" +
		"                                 |___/                                \r\n")

	color.LightCyan.Println("\n[+] Version: 1.0.0 | Branch: develop/v1.0.0")
	color.LightCyan.Println("[+] https://github.com/SaigyoujiYuyuko233/RunningServer")
	color.LightCyan.Println("[+] Copyright © 2019 - " + time.Now().Format("2006") + " SaigyoujiYuyuko[3558168775]. All rights reserved.")

	color.LightCyan.Println("")

	/**
	  *	获取信息
	  */

	// 主机名
	name, err := os.Hostname()
	if err == nil {}

	// cpu信息
	var exec_cmd = exec.Command("/bin/bash","-c","cat /proc/cpuinfo | grep name | cut -f2 -d: | uniq -c")
	var out,erra = exec_cmd.Output()
	if erra != nil {}

	print(strings.Split(byteString(out)," ")[6])

	//var cpu_name = strings.Split(byteString(out)," ")[1]	// cpu型号
	//var cpu_cores = strings.Split(byteString(out)," ")[0]	// cpu核数


	// cpu信息 #2
	//exec_cmd = exec.Command("cat /proc/cpuinfo | grep \"model name\"")
	//var cpu_modul_buf bytes.Buffer
	//exec_cmd.Stdout = &cpu_modul_buf
	//var cpu_modul_name = cpu_modul_buf.String()

	color.White.Println("设备信息 >")
	color.White.Println(" - 主机名: " + name )
	//color.White.Println(" - CPU信息: " + cpu_name + " | " + cpu_cores + "核" )

}


func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

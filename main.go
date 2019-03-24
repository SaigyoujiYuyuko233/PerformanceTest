package main

import (
	"github.com/gookit/color"
	"os"
	"os/exec"
	"strconv"
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
	out,err := exec_cmd.Output()
	if err != nil {}

	var cpu_name = strings.Split(byteString(out),"  ")[3]	// cpu型号
	var cpu_cores = strings.Split(byteString(out),"  ")[2]	// cpu核数

	cpu_name = strings.Replace(cpu_name,"\n","",-1)
	cpu_cores = strings.Replace(cpu_cores," ","",-1)


	// cpu信息 #2
	exec_cmd = exec.Command("/bin/bash","-c","cat /proc/cpuinfo |grep MHz|uniq")
	out,err = exec_cmd.Output()
	if err != nil {}

	var cpu_clock = strings.Split(byteString(out),"cpu MHz		: ")[1]	// cpu频率
	cpu_clock = strings.Replace(cpu_clock,"\n","",-1)

	// 内存信息
	exec_cmd = exec.Command("/bin/bash","-c","cat /proc/meminfo | grep \"MemTotal\"")
	out,err = exec_cmd.Output()
	if err != nil {}

	var memory = strings.Replace(byteString(out),"MemTotal:","",-1)	// 内存大小
	memory = strings.Replace(memory," ","",-1)
	memory = strings.Replace(memory,"\n","",-1)
	memory = strings.Replace(memory,"kB","",-1)
	memory_num,err := strconv.Atoi(memory)
	memory = strconv.Itoa(memory_num/1024)

	color.White.Println("设备信息 >")
	color.White.Println(" - 主机名: " + name )
	color.White.Println(" - CPU信息: " + cpu_name + "@" + cpu_clock + "Mhz | " + cpu_cores + "核" )
	color.White.Println(" - 内存信息: " + memory + "GB")

}


func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

package main

import (
	"crypto/md5"
	"crypto/sha1"
	"flag"
	"github.com/gookit/color"
	"runtime"
	"strconv"
	"time"
)

func main() {

	// arg
	runningCore := flag.Int("core",1,"将要使用的CPU核数")
	fullcore_mod :=  flag.Bool("daze",false,"使用全部核心运行")
	flag.Parse()

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

	if *fullcore_mod == true {
		color.LightBlue.Println("\n[da★ze~] Full core mode! Core: " + strconv.Itoa(runtime.NumCPU()))
		runtime.GOMAXPROCS(runtime.NumCPU())
	}else{
		color.LightBlue.Println("\n[Core] Core: " + strconv.Itoa(*runningCore))
		runtime.GOMAXPROCS(*runningCore)
	}


	/**
	  *	获取信息
	  */

	// 主机名
	//name, err := os.Hostname()
	//if err == nil {}

	/*
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
	memory_num,err := strconv.ParseInt(memory,10,64)
	memory = strconv.FormatFloat(float64(memory_num/1024), 'f', 2, 64)

	color.White.Println("设备信息 >")
	color.White.Println(" - 主机名: " + name )
	color.White.Println(" - CPU信息: " + cpu_name + "@" + cpu_clock + "Mhz | " + cpu_cores + "核" )
	color.White.Println(" - 内存信息: " + memory + "MB")
	*/

	color.White.Print("\n开始第 1 次浮点运算测试...")
	float_usetime_1 := runFloatTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(float_usetime_1/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 2 次浮点运算测试...")
	float_usetime_2 := runFloatTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(float_usetime_2/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 3 次浮点运算测试...")
	float_usetime_3 := runFloatTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(float_usetime_3/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 4 次浮点运算测试...")
	float_usetime_4 := runFloatTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(float_usetime_4/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 5 次浮点运算测试...")
	float_usetime_5 := runFloatTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(float_usetime_5/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 6 次浮点运算测试...")
	float_usetime_6 := runFloatTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(float_usetime_6/1000/1000/1000), 'f', 4, 64) + "s")

	color.LightCyan.Println("")


	float_usetime_all := float_usetime_1 + float_usetime_2 + float_usetime_3 + float_usetime_4 + float_usetime_5 + float_usetime_6
	float_grade := (60 - float64(float_usetime_all/1000/1000/1000/6))*100

	color.LightCyan.Print("\n + 浮点测试完成! 用时: " + strconv.FormatFloat(float64(float_usetime_all/1000/1000/1000), 'f', 4, 64) + "s | 平均用时: ")
	color.Green.Print(strconv.FormatFloat(float64(float_usetime_all/1000/1000/1000/6), 'f', 4, 64) + "s")
	color.LightCyan.Print("  得分: ")
	color.Green.Print(strconv.FormatFloat(float_grade, 'f', 0, 64))

	color.Gray.Println("\n\n--------------------------------------------------------")
	color.Gray.Println("准备开始哈希计算测试")

	color.White.Print("\n开始第 1 次哈希运算测试...")
	hash_usetime_1 := runHashTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(hash_usetime_1/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 2 次哈希运算测试...")
	hash_usetime_2 := runHashTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(hash_usetime_2/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 3 次哈希运算测试...")
	hash_usetime_3 := runHashTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(hash_usetime_3/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 4 次哈希运算测试...")
	hash_usetime_4 := runHashTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(hash_usetime_4/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 5 次哈希运算测试...")
	hash_usetime_5 := runHashTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(hash_usetime_5/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 6 次哈希运算测试...")
	hash_usetime_6 := runHashTest()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(hash_usetime_6/1000/1000/1000), 'f', 4, 64) + "s")

	hash_usetime_all := hash_usetime_1 + hash_usetime_2 + hash_usetime_3 + hash_usetime_4 + hash_usetime_5 + hash_usetime_6
	hash_grade := (60 - float64(hash_usetime_all/1000/1000/1000/6))*200

	color.LightCyan.Print("\n\n + 哈希运算测试完成! 用时: " + strconv.FormatFloat(float64(hash_usetime_all/1000/1000/1000), 'f', 4, 64) + "s | 平均用时: ")
	color.Green.Print(strconv.FormatFloat(float64(hash_usetime_all/1000/1000/1000/6), 'f', 4, 64) + "s")
	color.LightCyan.Print("  得分: ")
	color.Green.Print(strconv.FormatFloat(hash_grade, 'f', 0, 64))


	color.Gray.Println("\n\n--------------------------------------------------------")
	color.Gray.Println("准备开始 MD5 计算测试")


	color.White.Print("\n开始第 1 次 MD5 运算测试...")
	md5_usetime_1 := runMd5Test()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(md5_usetime_1/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 2 次 MD5 运算测试...")
	md5_usetime_2 := runMd5Test()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(md5_usetime_2/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 3 次 MD5 运算测试...")
	md5_usetime_3 := runMd5Test()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(md5_usetime_3/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 4 次 MD5 运算测试...")
	md5_usetime_4 := runMd5Test()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(md5_usetime_4/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 5 次 MD5 运算测试...")
	md5_usetime_5 := runMd5Test()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(md5_usetime_5/1000/1000/1000), 'f', 4, 64) + "s")

	color.White.Print("\n开始第 6 次 MD5 运算测试...")
	md5_usetime_6 := runMd5Test()
	color.LightGreen.Print("  完成! 用时: " + strconv.FormatFloat(float64(md5_usetime_6/1000/1000/1000), 'f', 4, 64) + "s")

	md5_usetime_all := md5_usetime_1 + md5_usetime_2 + md5_usetime_3 + md5_usetime_4 + md5_usetime_5 + md5_usetime_6
	md5_grade := (60 - float64(md5_usetime_all/1000/1000/1000/6))*200

	color.LightCyan.Print("\n\n + MD5 运算测试完成! 用时: " + strconv.FormatFloat(float64(hash_usetime_all/1000/1000/1000), 'f', 4, 64) + "s | 平均用时: ")
	color.Green.Print(strconv.FormatFloat(float64(md5_usetime_all/1000/1000/1000/6), 'f', 4, 64) + "s")
	color.LightCyan.Print("  得分: ")
	color.Green.Print(strconv.FormatFloat(md5_grade, 'f', 0, 64))

}


func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func runFloatTest() float64 {
	var start = time.Now().UnixNano()

	var float_a = 2.33
	var float_b = 0.3

	var float_sum = 0.00

	for i := 0;i < 50000000000 ;i++  {
		float_sum = float_a * float_b * float_sum
	}

	var end  = time.Now().UnixNano()
	return float64(end - start)
}

func runHashTest() float64 {
	var start = time.Now().UnixNano()

	var hash_1 = "SaigyoujiYuyuko"
	hash_sum := []byte(hash_1)

	for i := 0;i < 40000000 ;i++  {
		hash := sha1.New()
		hash.Write([]byte(hash_1 + byteString(hash_sum)))
		hash_sum = hash.Sum(nil)
	}

	var end  = time.Now().UnixNano()
	return float64(end - start)
}

func runMd5Test() float64 {
	var start = time.Now().UnixNano()

	var md_1 = "SaigyoujiYuyuko"
	md_sum := []byte(md_1)

	for i := 0;i < 40000000 ;i++  {
		md5 := md5.New()
		md5.Write([]byte(md_1 + byteString(md_sum)))
		md_sum = md5.Sum(nil)
	}

	var end  = time.Now().UnixNano()
	return float64(end - start)
}
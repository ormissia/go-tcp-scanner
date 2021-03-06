package main

import (
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"tcp-scanner/util"
)

var (
	address = "127.0.0.1" //要扫描的ip,默认为本机ip
)

func main() {
	//获取从命令行输入的参数（从第二个参数开始,到第四个参数结束）
	configs := os.Args[1:4]
	address = configs[0]
	//使用正则匹配参数是否时域名或者IP
	isIp := util.StringRegexp(util.RegexpExpressionIP, address)
	isDomainName := util.StringRegexp(util.RegexpExpressionDomainName, address)
	//当参数既不是IP也不是域名时终止
	if !isIp && !isDomainName {
		panic("请正确输入需要扫描的地址")
	}
	startPort, err := strconv.Atoi(configs[1])
	if err != nil {
		fmt.Println("请输入正确的端口号")
		panic(err)
	}
	endPort, err := strconv.Atoi(configs[2])
	if err != nil {
		fmt.Println("请输入正确的端口号")
		panic(err)
	}
	//创建缓冲区为5000的channel
	ports := make(chan int, 5000)
	defer close(ports)
	result := make(chan int)
	defer close(result)
	var openPorts []int

	//创建worker
	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}

	//另起一个线程分发任务
	go func() {
		for i := startPort; i < endPort; i++ {
			ports <- i
		}
	}()

	//接收返回结果
	for i := startPort; i < endPort; i++ {
		port := <-result
		//port不为0时为打开的端口
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	//排序
	sort.Ints(openPorts)

	fmt.Printf("=================\n以下是开放端口:\n")
	//输出结果
	for _, port := range openPorts {
		fmt.Printf("%d 开放了！！！\n", port)
	}
}

func worker(ports <-chan int, result chan<- int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", address, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%d端口未开放\n", port)
			//端口未则打开传0
			result <- 0
			continue
		}
		conn.Close()
		fmt.Printf("%d端口已开放\n", port)
		//端口打开则传端口值
		result <- port
	}
}

//非并发版扫描器
/*func main() {
	for i := 0; i < 120; i++ {
		address := fmt.Sprintf("127.0.0.1:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s 关闭了\n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s 打开了！！！", address)
	}
}*/

//普通并发版扫描器
/*func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 6000; i++ {
		//每一次循环加以1
		wg.Add(1)
		go func(j int) {
			//每一个goroutine执行完后减一
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s 关闭了\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了！！！\n", address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(startTime) / 1e9
	fmt.Printf("\n\n一共使用了%d秒", elapsed)
}*/

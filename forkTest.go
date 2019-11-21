package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func init() {
	godaeam := flag.Bool("d", false, "app run as daemon")

	if !flag.Parsed() {
		flag.Parse()
	}
	if *godaeam {
		log.Println("尝试启动子进程")
		cmd := exec.Command(os.Args[0])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			log.Print("启动子进程失败")
			return
		}
		log.Println("子进程[PID]:", cmd.Process.Pid)
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT)
		//主进程在这里被阻塞
		exitChan := make(chan bool)
		go func(chan bool) {
			_ = cmd.Wait()
			exitChan <- true
		}(exitChan)

		select {
		case <-ch:
			log.Println("信号量!")
		case <-exitChan:
			log.Println("子进程退出", cmd.ProcessState.ExitCode())
		}
		//for{
		//父进程这里
		//if err:=cmd.Wait();err==nil{
		//	log.Println("子进程执行完毕")
		//}
		//}
		//log.Println("wait 非阻塞的")
		//	os.Exit(0)
	}
}
func main() {
	log.Println("子进程的输出\n")
}

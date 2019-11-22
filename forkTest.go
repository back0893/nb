package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func init() {
	godaeam := flag.Bool("d", false, "已守护进程执行")
	sig := flag.String("s", "start", "动作指令")

	if !flag.Parsed() {
		flag.Parse()
	}
	switch *sig {
	case "start":
		log.Println("启动")
	case "stop":
		fp, err := os.Open("./daemon")
		if err != nil {
			log.Println("进程pid获得失败")
		}
		data, err := ioutil.ReadAll(fp)
		if err != nil {
			log.Println("进程pid获得失败")
		}
		pid := string(data)
		cmd := exec.Command("kill", "-2", pid)
		cmd.Run()
		log.Println("停止完成")
		os.Exit(0)
	}
	log.Println("开始")
	if *godaeam {
		log.Println("尝试启动子进程")
		cmd := exec.Command(os.Args[0])
		defer cmd.Process.Kill()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			log.Print("启动子进程失败")
			return
		}
		log.Println("子进程[PID]:", cmd.Process.Pid)
		//文件保存进程pid,方便重启和停止
		fp, err := os.Create("./daemon")
		if err != nil {
			panic("保存子进程pid失败")
		}
		defer fp.Close()
		fp.WriteString(fmt.Sprint(cmd.Process.Pid))
	}
}
func main() {
	//子进程注册信号
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hi,word!"))
	})
	server := &http.Server{Addr: "0.0.0.0:8002", Handler: nil}
	go server.ListenAndServe()
	//监听信号量
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	select {
	case <-ch:
		log.Println("信号量!")
		server.Close()
		os.Exit(0)
	}
}

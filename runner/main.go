package main

import (
	"log"
	"os"
	"time"

	"github.com/lxxtech/runner"
)
const timeout=3*time.Second
func main(){
	log.Println("Starting work")
	//为本次运行分配超时时间
	r:=runner.New(timeout)
	//加入要执行的任务
	r.Add(createTask(),createTask(),createTask())
	if err:=r.Start();err!=nil{
		switch err{
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout")
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended")
}

func createTask() func(int){
	return func(id int){
		log.Printf("Processor - Task #%d.",id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}
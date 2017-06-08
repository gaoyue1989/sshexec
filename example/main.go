package main

import (
	"time"
	"github.com/gaoyue1989/sshexec"
	"log"
)

func main() {
	sshExecAgent := sshexec.SSHExecAgent{}
	sshExecAgent.Worker = 10;
	sshExecAgent.TimeOut = time.Duration(120) * time.Second
	s, err := sshExecAgent.SftpHostByKey([]string{"193.168.2.1","193.168.2.2"}, 22,"root", "/data/test.log", "/data/test.log")
	log.Println("res:",s)
	log.Println("err:",err)

	s1, err1 := sshExecAgent.SshHostByKey([]string{"193.168.2.1","193.168.2.2"}, 22,"root", "ll")
	log.Println("res:",s1)
	log.Println("err:",err1)
	for {
		time.Sleep(1000)
	}
}

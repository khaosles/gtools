package gcmd

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

/*
   @File: command.go
   @Author: khaosles
   @Time: 2023/4/12 20:38
   @Desc:
*/

func Exec(cmdName string, args ...string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	log.Println("[CMD] 执行命令=> ", cmdName)
	log.Println("[CMD] 参数=> ", args)
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println("[CMD] Error ", err.Error())
		return "", err
	}
	resp := string(bytes)
	fmt.Println(resp)
	log.Println("[CMD] ", "Command finished")
	return resp, nil
}

// Sync 同步执行cmd并打印输出
func Sync(cmdName string, args ...string) error {
	log.Println("[CMD] 执行命令=> ", cmdName)
	log.Println("[CMD] 参数=> ", args)
	cmd := exec.Command(cmdName, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("[CMD] Error:", err)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("[CMD] Error:", err)
		return err
	}
	if err := cmd.Start(); err != nil {
		log.Println("[CMD] Error:", err)
		return err
	}
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			log.Println("[CMD] ", scanner.Text())
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Println("[CMD] ", scanner.Text())
		}
	}()
	if err := cmd.Wait(); err != nil {
		log.Println("[CMD] Error:", err)
		return err
	}

	log.Println("[CMD] ", "Command finished")
	return nil
}

// Asyn 异步执行cmd
func Asyn(cmdName string, args ...string) {
	log.Println("[CMD] 执行命令=> ", cmdName)
	log.Println("[CMD] 参数=> ", args)
	cmd := exec.Command(cmdName, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	if err := cmd.Start(); err != nil {
		log.Println("[CMD] Error:", err)
		return
	}
	log.Println("[CMD] Command start!")
}

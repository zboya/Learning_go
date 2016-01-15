package main

import (
//	"fmt"
	"os/exec"
	"os"
)

func main () {
	env:=os.Environ()
	procAttr:=&os.ProcAttr{
		Env:env,
		Files:[]*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	pid,_:=os.StartProcess("/bin/ls",[]string{"ls","-l"},procAttr)
	println("pid",pid)

	//cmd:=exec.Command("ls") //why can't run
	cmd:=exec.Command("gedit")
	cmd.Run()
}

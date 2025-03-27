package main

import (
	"fmt"
	"os/exec"
	_ "time"
)

func normal() {
	cmd := exec.Command("sleep", "5")

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	fmt.Printf("Started process with PID %d\n", cmd.Process.Pid)

	// time.Sleep(2 * time.Second)
	fmt.Println("Doing other work...")

	// err = cmd.Wait()
	// if err != nil {
	//     fmt.Println("Error waiting for command execution:", err)
	//     return
	// }

	fmt.Println("Command execution completed.")
}
func main() {

	// normal()
	synchronus()
	// asynchronous()
	// output()
}
func synchronus() {
	cmd1 := exec.Command("sleep", "2")
	cmd2 := exec.Command("sleep", "2")

	cmd1.Run()
	cmd2.Run()

	fmt.Println("Total time: 4 sec")
}

func asynchronous() {
	cmd1 := exec.Command("sleep", "2")
	cmd2 := exec.Command("sleep", "2")

	cmd1.Start()
	cmd2.Start()

	cmd1.Wait()
	cmd2.Wait()

	fmt.Println("Total time: 2 sec")
}

func output() {
	cmd1 := exec.Command("ls", "../HTTP_programs")
	Output, err := cmd1.CombinedOutput()
	if err != nil {
		fmt.Println("has eroor")
		fmt.Println(err)

	}

fmt.Println(string(Output))
}

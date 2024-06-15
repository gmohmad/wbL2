/*
=== Interaction with the OS ===

You need to implement your own shell

built-in commands: cd/pwd/echo/kill/ps
support fork/exec commands
conveyor on pipes

Implement netcat (nc) client utility
receive data from stdin and send it to the connection (tcp/udp)
The program must pass all tests. The code must pass go vet and golint checks.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// main function runs the shell loop
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("shell> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		// If input is empty skip it
		if input == "" {
			continue
		}

		// if input is \quit, quit the shell
		if input == "\\quit" {
			break
		}

		// process commands
		commands := strings.Split(input, "|")
		// if commands are piped, call handlePipes function
		if len(commands) > 1 {
			if err := handlePipes(commands); err != nil {
				fmt.Println("Error handling pipes:", err)
			}
			continue
		}
		// else run the command
		args := strings.Fields(commands[0])
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			changeDirectory(args)
		case "pwd":
			printWorkingDirectory()
		case "echo":
			echo(args)
		case "kill":
			killProcess(args)
		case "ps":
			printProcesses()
		default:
			executeCommand(args)
		}
	}
}

// handlePipes processes commands connected by pipes
func handlePipes(commands []string) error {
	// lastOutput is needed to set the output of the command to the next one
	var lastOutput *os.File

	for i, command := range commands {
		args := strings.Fields(command)
		cmd := exec.Command(args[0], args[1:]...)

		// if it's not the first command, set the input of the command to lastOutput
		if i > 0 {
			cmd.Stdin = lastOutput
		}

		// if it's not the last command, create a new pipe and set hte outputs
		if i < len(commands)-1 {
			r, w, err := os.Pipe()
			if err != nil {
				return fmt.Errorf("error creating pipe: %v", err)
			}
			cmd.Stdout = w
			lastOutput = r
			defer w.Close()
		} else {
			// if it is the last command, the set the stdout to the os.stdout
			cmd.Stdout = os.Stdout
		}

		// set command's stderr to os.stderr
		cmd.Stderr = os.Stderr

		// start the command
		if err := cmd.Start(); err != nil {
			return fmt.Errorf("error starting command: %v", err)
		}

		// close the output
		if i > 0 {
			lastOutput.Close()
		}
		// wait for the command
		if err := cmd.Wait(); err != nil {
			return fmt.Errorf("error waiting for command: %v", err)
		}
	}
	return nil
}

// changeDirectory handles the cd command
func changeDirectory(args []string) {
	if len(args) < 2 {
		fmt.Println("cd: missing argument")
		return
	}
	if err := os.Chdir(args[1]); err != nil {
		fmt.Println("cd:", err)
	}
}

// printWorkingDirectory handles the pwd command
func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("pwd:", err)
	} else {
		fmt.Println(dir)
	}
}

// echo handles the echo command
func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

// killProcess handles the kill command
func killProcess(args []string) {
	if len(args) < 2 {
		fmt.Println("kill: missing argument")
		return
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("kill: invalid pid:", args[1])
		return
	}
	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		fmt.Println("kill:", err)
	}
}

// printProcesses handles the ps command
func printProcesses() {
	cmd := exec.Command("ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("ps:", err)
	}
}

// executeCommand handles execution of external commands
func executeCommand(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(args[0], ":", err)
	}
}

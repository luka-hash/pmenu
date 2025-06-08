// Copyright © 2025 Luka Ivanović
// This code is licensed under the terms of the MIT licence (see LICENCE for details)

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var programNames []string
	for _, prog := range programs {
		programNames = append(programNames, prog.Name)
	}
	wofiInput := strings.Join(programNames, "\n")

	cmd := exec.Command("wofi", "--dmenu", "--prompt", "Run Program", "--columns", "1")
	cmd.Stdin = strings.NewReader(wofiInput)
	
	output, err := cmd.Output()
	if err != nil {
		// Cancelled by user or wofi failed
		os.Exit(0)
	}

	selectedName := strings.TrimSpace(string(output))
	if selectedName == "" {
		os.Exit(0)
	}

	var selectedProgram *Program
	for _, prog := range programs {
		if prog.Name == selectedName {
			selectedProgram = &prog
			break
		}
	}

	if selectedProgram == nil {
		// log.Printf("Program not found: %s", selectedName)
		os.Exit(1)
	}

	executeCommand(selectedProgram.Command)
}

func executeCommand(command string) {
	// Currently we treat all commands as external

	// Split command into parts for proper execution
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.CombinedOutput()

	// Send output to notification daemon if there's output or an error
	if len(output) > 0 || err != nil {
		message := string(output)
		if err != nil {
			if message == "" {
				message = err.Error()
			} else {
				message = fmt.Sprintf("%s\nError: %s", message, err.Error())
			}
		}
		
		if message != "" {
			sendNotification("pmenu", message)
		}
	}
}

func sendNotification(title, message string) {
	cmd := exec.Command("notify-send", title, message)
	_ = cmd.Run()
}

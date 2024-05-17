package cli

import(
	"fmt"
	"goctftools/config/colors"
	"os/exec"
	"bufio"
)

func Nmap(ip string) {
	// Use pipe instead of capturing output
	fmt.Println(color.Magenta + "|> [NMAP] Start..." + color.Reset)
	cmd := exec.Command("nmap", "-A", "-p-", "-v", ip)
  
	// Open pipe for stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
	  fmt.Println(color.Red + "|> Error creating stdout pipe:", err, color.Reset)
	  return
	}
  
	// Start the command
	err = cmd.Start()
	if err != nil {
	  fmt.Println(color.Red + "|> Error starting Nmap scan:", err, color.Reset)
	  return
	}
  
	// Read from the pipe and print to terminal concurrently
	go func() {
	  scanner := bufio.NewScanner(stdout)
	  for scanner.Scan() {
		fmt.Println(scanner.Text())
	  }
	}()
  
	// Wait for command to finish
	err = cmd.Wait()
	if err != nil {
	  fmt.Println(color.Red + "|> Nmap scan error:", err)
	} else {
	  fmt.Println(color.Green + "\n|> [NMAP] Scan completed successfully.\n" + color.Reset)
	}
}
  
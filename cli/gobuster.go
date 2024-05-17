package cli

import(
	"fmt"
	"goctftools/config/colors"
	"os/exec"
	"bufio"
)

func Gobuster(ip string) {

	var httptype string
	var folderArea string
	var wordlist string

	fmt.Println(color.Cyan + "|> [gobuster] Configuration" + color.Reset)
	fmt.Print(color.Cyan + "|> http (or https): "+color.White+"http"+ color.Magenta)
	fmt.Scanln(&httptype)
	fmt.Println(color.Reset)

	if httptype == ""{
		httptype = "http"
	} else{
		httptype = "https"
	}

	fmt.Print(color.Cyan + "|> Path (eg : http://10.10.10.10"+color.White+"/admin"+color.Cyan+"): "+color.White+httptype+"://"+ip+color.Magenta)
	fmt.Scanln(&folderArea)
	fmt.Println(color.Reset)

	if(folderArea == ""){
		folderArea = httptype+"://"+ip
	} else{
		folderArea = httptype+"://"+ip+folderArea
	}

	fmt.Print(color.Cyan + "|> wordlist (default: /usr/share/wordlists/dirbuster/directory-list-2.3-small.txt): "+ color.Magenta)
	fmt.Scanln(&wordlist)
	fmt.Println(color.Reset)

	if(wordlist == ""){
		wordlist = "/usr/share/wordlists/dirbuster/directory-list-2.3-small.txt"
	}

	// Use pipe instead of capturing output
	fmt.Println(color.Magenta + "|> [gobuster] Start..." + color.Reset)
	cmd := exec.Command("gobuster", "dir", "-w", wordlist,"-x", "php,txt,html,js,json", "-u", folderArea)
  
	// Open pipe for stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
	  fmt.Println(color.Red + "|> Error creating stdout pipe:", err, color.Reset)
	  return
	}
  
	// Start the command
	err = cmd.Start()
	if err != nil {
	  fmt.Println(color.Red + "|> Error starting gobuster scan:", err, color.Reset)
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
	  fmt.Println(color.Red + "|> gobuster scan error:", err)
	} else {
	  fmt.Println(color.Green + "\n|> [gobuster] Scan completed successfully.\n" + color.Reset)
	}
}
  
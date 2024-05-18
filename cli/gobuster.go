package cli

import(
	"fmt"
	"goctftools/config/colors"
	"os/exec"
	"bufio"
)

var(
	httptype string
	folderArea string
	wordlist string
	taglist string
	reset string
)

func Gobuster(ip string) {

	if httptype != ""{
		fmt.Println(color.Cyan+"|> [Resume Information]"+color.Reset)
		fmt.Println(color.Cyan+"|> Target: "+color.White+folderArea+color.Reset)
		fmt.Println(color.Cyan+"|> Return extension types: "+color.White+taglist+color.Reset)
		fmt.Println(color.Cyan+"|> Wordlist: "+color.White+wordlist+color.Reset+"\n")

		fmt.Println(color.Cyan + "|> [0] No" + color.Reset)
		fmt.Println(color.Cyan + "|> [1] Yes"+ color.Reset)
		fmt.Print(color.Cyan + "|> Do you want reset information ? "+color.Magenta)
		fmt.Scanln(&reset)

		if(reset == "1"){
			httptype = ""
			folderArea = ""
			wordlist = ""
			taglist = ""
			reset = ""
		}
	}


	if httptype == "" && folderArea == "" && wordlist == "" && taglist == ""{
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

	if( wordlist == ""){
		wordlist = "/usr/share/wordlists/dirbuster/directory-list-2.3-small.txt"
	}

	fmt.Print(color.Cyan + "|> taglist: "+color.White+"php,css,scss,sass,txt,html,js,json"+ color.Magenta)
	fmt.Scanln(&taglist)
	fmt.Println(color.Reset)

	if(taglist == ""){
		taglist = "php,css,scss,sass,txt,html,js,json"
	} else{
		taglist = "php,css,scss,sass,txt,html,js,json"+taglist
	}
	}

	// Use pipe instead of capturing output
	fmt.Println(color.Magenta + "|> [gobuster] Start..." + color.Reset)
	cmd := exec.Command("gobuster", "dir", "-w", wordlist,"-x", taglist, "-u", folderArea)
  
	// Open pipe for stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
	  fmt.Println(color.Red + "|> [gobuster] Error creating stdout pipe:", err, color.Reset)
	  return
	}
  
	// Start the command
	err = cmd.Start()
	if err != nil {
	  fmt.Println(color.Red + "|> [gobuster] Error starting gobuster scan:", err, color.Reset)
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
	  fmt.Println(color.Red + "|> [gobuster] Scan error:", err)
	} else {
	  fmt.Println(color.Green + "\n|> [gobuster] Scan completed successfully.\n" + color.Reset)
	}
}
  
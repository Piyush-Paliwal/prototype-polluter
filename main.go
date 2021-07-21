package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"regexp"
	"os/exec"
	"flag"
)

func requirement() {
	check, _ := exec.Command("bash", "-c", "command -v page-fetch").Output()
	if len(string(check)) == 0 {
                fmt.Println("Trying to install the Requirements...")
		install, err := exec.Command("go", "get", "github.com/detectify/page-fetch").Output()
		if err != nil {
			fmt.Println("Failed to install Requirements")
			os.Exit(1)
		}
		fmt.Println("Installation done", install)
        }
	//Requirments Fulfilled
	fmt.Println("Starting.....")
}

func main() {

	//Colours
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	//Flag-Error-Handling
	flag.Usage = func() {
        fmt.Printf("Usage of our Program: \n")
        flag.PrintDefaults()  // prints default usage
	}

	//Flag-Arguments
	wordPtr := flag.Bool("v", false, "Verbose mode even output's the Not Vulnerable test")
	flag.Parse()

	//Checking Requirements
	requirement()

	//Main | Start
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line = scanner.Text()
		
		if len(line) == 0 {
			continue
		}
		
		matched, _ := regexp.MatchString(`\?`, line)
		if matched {
			excd := ("echo '"+line+"&__proto__[testparam]=testval' | page-fetch -j \"window.testparam == 'testval'? 'Vulnerable' : 'Not Vulnerable'\" -o /tmp/out | grep ^JS | awk '{print $3, $4}'")
			cmd, _ := exec.Command("bash", "-c", excd).Output()
			output := string(cmd[:])
			NotVulnerable, _ := regexp.MatchString("Not Vulnerable", output)
			if !NotVulnerable {
				fmt.Println(string(colorRed),"Vulnerable --> "+line+"&__proto__[testparam]=testval")
			} else if NotVulnerable && *wordPtr == true{
				fmt.Println(string(colorGreen),"Not Vulnerable --> "+line+"&__proto__[testparam]=testval")
			}
		} else {
			excd := ("echo '"+line+"?__proto__[testparam]=testval' | page-fetch -j \"window.testparam == 'testval'? 'Vulnerable' : 'Not Vulnerable'\" -o /tmp/out | grep ^JS | awk '{print $3, $4}'")
			cmd, _ := exec.Command("bash", "-c", excd).Output()
			output := string(cmd[:])
			NotVulnerable, _ := regexp.MatchString("Not Vulnerable", output)
			if !NotVulnerable {
				fmt.Println(string(colorRed),"Vulnerable --> "+line+"?__proto__[testparam]=testval")
			} else if NotVulnerable && *wordPtr == true{
				fmt.Println(string(colorGreen),"Not Vulnerable --> "+line+"?__proto__[testparam]=testval")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"bufio"
	"fmt"
	_ "io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	command := "curl"

	bin, lookErr := exec.LookPath(command)
	if lookErr != nil {
		log.Fatal(lookErr)
	}

	fmt.Println(bin)

	file, err := os.Open("1080P.m3u8")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()		
		if strings.HasPrefix(line, "#"){
			continue
		}
		fmt.Println(line)
		
		cmd := exec.Command(bin, "-O", line)
		runErr := cmd.Run()
		if runErr != nil {
			log.Fatal(runErr)
		}
	}
}

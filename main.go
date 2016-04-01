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
	
 
	base := "http://wpc.1765a.taucdn.net/801765A/video/uploads/videos/6aa46e5d-e0db-484a-ab83-a36c9c662fda/"
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()		
		if strings.HasPrefix(line, "#"){
			continue
		}
		url := strings.Join(base, line)
		fmt.Println(url)
		
		cmd := exec.Command(bin, "-O", url)
		runErr := cmd.Run()
		if runErr != nil {
			log.Fatal(runErr)
		}
	}
}

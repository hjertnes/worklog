package main

import (
	"fmt"
	"git.sr.ht/~hjertnes/doing/config"
	"git.sr.ht/~hjertnes/doing/utils"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func help(){
	fmt.Println("worklog is a small utlitity for timetracking")
	fmt.Println("It writes the logs to an org file and you can freely add notes to it if needed")
	fmt.Println("Usage:")
	fmt.Println("\t worklog text you want to append")
	fmt.Println("Configuration")
	fmt.Println("\t config file is stored at ~/.worklog.yml")
	fmt.Println("\t currently has only one key Path that need to be pointed at your org-roam folder")
}

func success(){
	conf, err := config.Read()
	if err != nil{
		panic(err)
	}

	text := strings.Join(os.Args[1:], " ")
	now := time.Now()
	current := now.Format("15:04:05")
	today:= now.Format("2006-01-02")

	filename := utils.ReplaceTilde(fmt.Sprintf("%sworklog.org", conf.Path))

	if !utils.Exist(filename){
		err := ioutil.WriteFile(filename, []byte(""), os.ModePerm)
		if err != nil{
			panic(err)
		}
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil{
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == ""{
		lines = lines[0:len(lines)-1]
	}
	//lines = append(lines, fmt.Sprintf("- %s: %s", current, text))
	lines = append(lines, fmt.Sprintf("* %s", text))
	lines = append(lines, fmt.Sprintf("- Started: %s %s", today, current))

	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), os.ModePerm)
	if err != nil{
		panic(err)
	}
}
func main(){
	if len(os.Args) > 1{
		success()
	} else {
		help()
	}
}

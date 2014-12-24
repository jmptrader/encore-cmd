//TODO:

//License: MIT
//Author: Gani Mendoza (itjumpstart.wordpress.com)

//This program aims to solve a subset of configuration management tasks
//Cmdfile is the only argument required for now
//Cmdfile takes inspiration from Dockerfile
//Cmdfile tasks must be sequential (no loops or conditionals)
//Cmdfile tasks are Bash commands and external programs
//Cmdfile is for humans, not machines

//Limitations of cmdfile
//It is not a shell (so no variable declaration and substitution)
//No pipe commands
//No backslash (commands must be put on each line)
//No &&
//No cd (Use GO chdir directoryname)

package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"strings"
)

func printError(err error) {

	if err != nil {
		color.Set(color.FgRed)
		os.Stderr.WriteString(fmt.Sprintf("==> ERROR: %s\n", err.Error()))
		color.Unset()
	}

}

func printOutput(outs []byte) {

	if len(outs) > 0 {
		color.Set(color.FgGreen)
		fmt.Printf("==> OUTPUT: %s\n", string(outs))
		color.Unset()
	}
}

func echoCmd(args []string) {
	var cmd *exec.Cmd
	//echo "Hello world" > echofile.txt
	strJoin := strings.Join(args, " ")
	if strings.Contains(strJoin, " > ") {

		slcSplit := strings.Split(strJoin, " > ")
		idxFile := len(slcSplit) - 1
		fmt.Println(slcSplit)

		outfile, err := os.Create(slcSplit[idxFile])
		defer outfile.Close()

		if err != nil {
			printError(err)
			return
		}

		str := strings.Replace(slcSplit[0], "\"", "", 2)

		cmd = exec.Command("echo", str)
		cmd.Stdout = outfile

		err = cmd.Start()
		if err != nil {
			printError(err)
			return
		}
		cmd.Wait()

	}

	if strings.Contains(strJoin, " >> ") {
		//http: //stackoverflow.com/questions/13513375/how-to-append-text-to-a-file-in-golang
		slcSplit := strings.Split(strJoin, " >> ")
		idxFile := len(slcSplit) - 1
		fmt.Println(slcSplit)

		outfile, err := os.OpenFile(slcSplit[idxFile], os.O_RDWR|os.O_APPEND, 0666)

		defer outfile.Close()

		str := strings.Replace(slcSplit[0], "\"", "", 2)
		_, err = outfile.WriteString(str)

		if err != nil {
			printError(err)
			return
		}

		outfile.Sync()

	} else {
		cmd = exec.Command("echo", args...)

		output, err := cmd.CombinedOutput()
		printError(err)
		printOutput(output)
	}
}

//http://stackoverflow.com/questions/18986943/in-golang-how-can-i-write-the-stdout-of-an-exec-cmd-to-a-file
//http://stackoverflow.com/questions/12907653/what-is-difference-between-string-and-string-in-golang
func runCmd(argCommand string, args []string) error {
	fmt.Println(argCommand, args)

	var cmd *exec.Cmd
	var output []byte
	var err error
	switch argCommand {
	case "echo":
		echoCmd(args)
	//http://stackoverflow.com/questions/20568515/how-to-use-sed-to-replace-a-config-files-variable
	//for some reason, sed in debian 7.2 doesn't like single quote
	case "sed":
		cmd = exec.Command("sed", args...)

		output, err = cmd.CombinedOutput()
		printError(err)
		printOutput(output)

	case "mkdir":
		dir := args[0]

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			cmd = exec.Command("mkdir", args...)

			output, err = cmd.CombinedOutput()
			printError(err)
			printOutput(output)
		}

	default:
		cmd = exec.Command(argCommand, args...)

		output, err = cmd.CombinedOutput()
		printError(err)
		printOutput(output)
	}
	return err
}

func pipe(args []string) {
	fmt.Println(len(args))

	//sh.Command("echo", "hello", "world").Command("awk", "{print $1}").Run()
	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
}

func chdirCmd(dir string) error {
	fmt.Println("chdir " + dir)

	err := os.Chdir(dir)

	if err != nil {
		printError(err)
	} else {
		printOutput([]byte("chdir to " + dir))
	}
	return err
}

func getenvCmd(key string) error {
	fmt.Println("getenv " + key)

	result := os.Getenv(key)

	if len(result) == 0 {
		err := errors.New("No environment variable named " + key)
		printError(err)
		return err
	} else {
		printOutput([]byte("getenv " + key + "=" + result))
		return nil
	}
}

func setenvCmd(key, value string) error {
	if key == "" || value == "" {
		return errors.New("Error in ENV. Key or value is blank")
	}

	fmt.Println("ENV " + key + " " + value)

	err := os.Setenv(key, value)

	if err != nil {
		printError(err)
		return err
	} else {
		printOutput([]byte("ENV " + key + "=" + value))
		return nil
	}
}

func hostenvCmd(key string) error {
	fmt.Println("hostenv " + key)

	slc := os.Environ()

	found := false

	for _, v := range slc {
		//fmt.Println(slc[k])
		pair := strings.Split(v, "=")

		if pair[0] == key {

			printOutput([]byte("hostenv: " + key + "=" + v))

			found = true
			break
		}
	}

	if !found {
		err := errors.New("No host environment variable named " + key)
		printError(err)
		return err
	} else {
		return nil
	}
}

//http://stackoverflow.com/questions/12907653/what-is-difference-between-string-and-string-in-golang
func goCmd(argCommand string, args []string) error {
	var err error
	switch argCommand {

	case "chdir":
		if len(args) != 1 {
			err = errors.New("GO chdir. Directory not specified")
			printError(err)
			return err
		} else {
			err = chdirCmd(args[0])
		}

	case "getenv":
		if len(args) != 1 {
			err = errors.New("GO getenv. Key is blank")
			printError(err)
			return err
		} else {
			err = getenvCmd(args[0])
		}

	case "hostenv":
		if len(args) != 1 {
			err = errors.New("GO setenv. Key is blank")
			printError(err)
			return err
		} else {
			err = hostenvCmd(args[0])
		}

	case "hostname":
		var str string
		str, err = os.Hostname()
		if err != nil {
			printError(err)
			return err
		} else {
			printOutput([]byte("hostname: " + str))
		}

	}

	return err
}

func processCmd(command string) error {
	var err error

	s := strings.TrimSpace(command)

	slcStr := strings.Split(s, " ")

	args := []string{}

	var argCommand string

	//value is RUN if cmdfile = RUN dpkg -s sudo
	cmd := strings.ToUpper(slcStr[0])

	if !strings.Contains(cmd, "FROM") || !strings.Contains(cmd, "MAINTAINER") {
		fmt.Println(cmd)
	}

	//http://play.golang.org/p/QWmzgIWpF8
	for i, _ := range slcStr {

		/*
			cmdfile.txt
			RUN dpkg -s sudo

			cmd := exec.Command("dpkg", "-s", "sudo")
		*/

		if i == 1 {
			argCommand = slcStr[i]
		} else if i > 1 {
			args = append(args, slcStr[i])
		}
	}

	switch cmd {

	case "RUN":
		err = runCmd(argCommand, args)

	case "GO":
		err = goCmd(argCommand, args)

	case "ENV":
		err = setenvCmd(argCommand, args[0])

	}

	return err
}

func main() {
	//file, err := os.Open("cmdfile.txt")

	if len(os.Args) != 2 {
		printError(errors.New("Please specify a cmdfile"))
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var ln string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		ln = scanner.Text()

		if !strings.Contains(ln, "#") {
			err = processCmd(ln)
		}

		if err != nil {
			break
			log.Fatal(err)
		}
	}

	fmt.Println("If any error appears, cmdfile is not completed. Press ENTER to exit")
	fmt.Scanln()
}

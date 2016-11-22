package main

import (
	"angular2-gin/server"
	"bufio"
	"log"

	"io"
	"os/exec"
)

func main() {

	go npmStart()
	server.StartSrv()

}

func npmStart() (err error) {
	cmd := exec.Command("npm", "start")
	stdout, err := cmd.StdoutPipe()
	stderr, err := cmd.StderrPipe()
	multi := io.MultiReader(stdout, stderr)

	if err != nil {
		return err
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		return err
	}
	// read command's stdout line by line
	in := bufio.NewScanner(multi)
	buf := make([]byte, 0, 64*1024)
	in.Buffer(buf, 1024*1024)

	for in.Scan() {
		log.Printf(in.Text()) // write each line to your log, or anything you need
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}
	return err
}

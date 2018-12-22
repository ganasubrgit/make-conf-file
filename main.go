package main

import (
	"os"
	"os/exec"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	SRVFILE := "/etc/systemd/system/example.service"
	//SRVFILE := "/tmp/example.service" //tmp path for testing
	f, err := os.Create(SRVFILE)
	check(err)
	defer f.Close()
	f.WriteString("[Unit]\nDescription=Restarts example.py\n")
	f.WriteString("\n[Service]\nUser=your username\nWorkingDirectory=/directory/of/script\nExecStart=/usr/bin/python3 /directory/of/script/example.py\nRestart=always")
	f.WriteString("\n\n[Install]\nWantedBy=multi-user.target\n\n")
	os.Chmod(SRVFILE, 0644)
	c := exec.Command("/bin/sh", "-c", "systemctl daemon-reload;systemctl start example.service;systemctl enable example.service")
	c.Run()
}

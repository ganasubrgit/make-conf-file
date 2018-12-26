package main

import (
	"os"
	"os/exec"
)

var APPFILE string = "/etc/PitCrew/vel.py"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exist(PATH string) {
	if _, err := os.Stat(PATH); os.IsNotExist(err) {
		println("Application not installed.", APPFILE, "file not found, Please check and try again.")
		help()
		os.Exit(1)
	}
}

//Support contact
func help() {
	println("For assistance, email support@tecknologics.com")
}

func main() {

	//check app file exist
	exist(APPFILE)
	SRVFILE := "/etc/systemd/system/veld.service"
	f, err := os.Create(SRVFILE)
	check(err)
	defer f.Close()
	f.WriteString("[Unit]\nDescription=Service demon for veld.py\n")
	f.WriteString("\n[Service]\nUser=root\nWorkingDirectory=/etc/PitCrew\nExecStart=/usr/bin/python2.7 /etc/PitCrew/vel.py\nRestart=always")
	f.WriteString("\n\n[Install]\nWantedBy=multi-user.target\n\n")
	os.Chmod(SRVFILE, 0644)
	c := exec.Command("/bin/sh", "-c", "systemctl daemon-reload;systemctl start veld.service;systemctl enable veld.service")
	c.Run()
	println("Veld service installed successfully!")
	help()
}

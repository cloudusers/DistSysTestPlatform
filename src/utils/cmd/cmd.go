package cmd

import (
	"bytes"
	"os"
	"os/exec"
)

type CmdFactory struct {
	workingDir   string
	materialsDir string
	msg          error
	status       bool
}

func (this *CmdFactory) SetWorkingDir(s string) {
	this.workingDir = s
}

func (this *CmdFactory) GetWorkingDir() string {
	return this.workingDir
}

func (this *CmdFactory) SetMaterialsDir(s string) {
	this.materialsDir = s
}

func (this *CmdFactory) GetMaterialsDir() string {
	return this.materialsDir
}

func (this *CmdFactory) SetMsg(s error) {
	this.msg = s
}

func (this *CmdFactory) GetMsg() error {
	return this.msg
}

func (this *CmdFactory) SetStatus(b bool) {
	this.status = b
}

func (this *CmdFactory) GetStatus() bool {
	return this.status
}

func (this *CmdFactory) RunCmd(scriptPath string) *CmdFactory {
	cmd := exec.Command("sh", "-c", scriptPath)

	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		this.msg = err
		this.status = false
		return this
	}

	this.status = true

	// this line is needed to extend current envs
	cmd.Env = os.Environ()

	// workinDir can be empty for `nmz init`
	if this.workingDir != "" {
		cmd.Env = append(cmd.Env, "DIR="+this.workingDir)
	}

	if this.materialsDir != "" {
		cmd.Env = append(cmd.Env, "DIR="+this.materialsDir)
	}

	return this
}

func NewCmdFactory() *CmdFactory {

	return &CmdFactory{
		workingDir:   "",
		materialsDir: "",
		msg:          nil,
		status:       false,
	}
}

var CommandFactory *CmdFactory

func init() {
	CommandFactory = NewCmdFactory()
}

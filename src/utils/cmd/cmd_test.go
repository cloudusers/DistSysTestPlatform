package cmd

import (
	"fmt"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestCmdFactory(t *testing.T) {
	f := NewCmdFactory()
	// there shouldn't any access to the file, actually
	//f.SetWorkingDir("/tmp/dummy1")
	//f.SetMaterialsDir("/tmp/dummy2")
	res := f.RunCmd("echo 42")
	if res.status {
		fmt.Println(res)
	}
	//assert.Contains(t, cmd.Env, "NMZ_WORKING_DIR=/tmp/dummy1")
}

func TestCommandFactory(t *testing.T) {
	res := CommandFactory.RunCmd("ping 10.134.99.128")
	if res.status {
		fmt.Println(res)
	} else {
		fmt.Println(res.msg)
	}
	//assert.Contains(t, cmd.Env, "NMZ_WORKING_DIR=/tmp/dummy1")
}

package component

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func ExecCmd(cmdStr string, timeOut int) (err error, stdout bytes.Buffer, stderr bytes.Buffer) {
	// 设置默认超时时间
	if timeOut == 0 {
		timeOut = 10
	}
	execDown := make(chan bool, 1)
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	go func() {
		err = cmd.Run()
		execDown <- true
	}()

	select {
	case <-execDown:
		return
	case <-time.After(time.Duration(timeOut) * time.Second):
		if killerr := cmd.Process.Kill(); killerr != nil {
			err = fmt.Errorf("EXEC CMD TIMEOUT:%d  %s\nKILL CMD ERROR:%s", timeOut, cmdStr, killerr.Error())
		} else {
			err = fmt.Errorf("EXEC CMD TIMEOUT:%d  %s", timeOut, cmdStr)
		}
		return
	}
}

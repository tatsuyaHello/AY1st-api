package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

// RunCommand は 外部コマンドを実行します
func RunCommand(command string, options, envList []string) error {
	cmd := exec.Command(command, options...)
	cmd.Env = append(os.Environ(), envList...)

	GetLogger().Debugf("%v", cmd.Args)

	// コマンドの標準エラー出力 は バッファとエラーログの両方へ流す
	var stderrBuf bytes.Buffer
	stderr := io.MultiWriter(&stderrBuf, ErrorLogWriter())

	stdoutIn, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderrIn, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		_, err := io.Copy(DebugLogWriter(), stdoutIn)
		if err != nil {
			GetLogger().Errorf("copy stdout error: %v", err)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		_, err := io.Copy(stderr, stderrIn)
		if err != nil {
			GetLogger().Errorf("copy stderr error: %v", err)
		}
		wg.Done()
	}()

	wg.Add(1)
	done := make(chan error)
	go func() {
		err = cmd.Wait()
		if err != nil {
			done <- fmt.Errorf("command %v が失敗: %v", command, err)
			return
		}

		exitCode := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
		if exitCode != 0 {
			done <- fmt.Errorf("command %v がExitCode %v で失敗", command, exitCode)
			return
		}

		wg.Done()
		close(done)
	}()

	wg.Wait()
	err = <-done
	if err != nil {
		return fmt.Errorf("message = %v, STDERR = %v", err, string(stderrBuf.Bytes()))
	}

	return nil
}

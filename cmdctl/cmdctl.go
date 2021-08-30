package cmdctl

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
)

// Status the status of process

const (
	// Stopped the stopped status
	Stopped int = iota

	// Starting the starting status
	Starting = 10

	// Running the running status
	Running = 20

	// Backoff the backoff status
	Backoff = 30

	// Stopping the stopping status
	Stopping = 40

	// Exited the Exited status
	Exited = 100

	// Fatal the Fatal status
	Fatal = 200

	// Unknown the unknown status
	Unknown = 1000
)

// ProcessConfig process config
type ProcessConfig struct {
	ExecStart        string
	User             string
	Group            string
	ExecOnce         bool
	RestartSec       int
	WorkingDirectory string
	SyncExec         bool //是否同步执行，同步执行，会等待程序结果退出
}

// Process struct
type Process struct {
	Config      ProcessConfig
	Name        string
	cmd         *exec.Cmd
	Status      int
	UserStopped bool
}

// ProcessContainer container
var ProcessContainer map[string]*Process

// Start start process
func (p *Process) Start() error {
	cmdstring := p.Config.ExecStart

	args := strings.Split(cmdstring, " ")
	// 启动时，需要设置userstopped为false，否则无法启动
	p.UserStopped = false

	go func() {
		for !p.UserStopped && !p.Config.ExecOnce {
			cmd := exec.Command(args[0], args[1:]...)
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
			p.cmd = cmd
			err := cmd.Start()
			if err == nil {
				err = cmd.Wait()
				p.Status = Exited
			}
			if p.Config.RestartSec != 0 {
				time.Sleep(time.Second * time.Duration(p.Config.RestartSec))
			}
		}
	}()

	return nil
}

// Stop stop process
func (p *Process) Stop() (err error) {
	p.UserStopped = true
	if p.cmd != nil && p.cmd.Process != nil {
		err = p.cmd.Process.Signal(syscall.SIGINT)
	}

	return err
}

// GetStatus get status
func (p *Process) GetStatus() int {
	return Running
}

// IsRunning is running
func (p *Process) IsRunning() bool {
	if p.cmd != nil && p.cmd.Process != nil {
		if runtime.GOOS == "windows" {
			proc, err := os.FindProcess(p.cmd.Process.Pid)
			return proc != nil && err == nil
		}
		return p.cmd.Process.Signal(syscall.Signal(0)) == nil
	}
	return false
}

// IsRunning is running
func IsRunning(name string) bool {
	p := ProcessContainer[name]
	if p == nil {
		return false
	}
	return p.IsRunning()
}

// Exist process exist
func Exist(name string) bool {
	p := ProcessContainer[name]
	if p != nil {
		return true
	}
	return false
}

// New new a process
func New(name string, conf *ProcessConfig) error {
	if ProcessContainer[name] != nil {
		return errors.New("process exist")
	}

	p := &Process{
		Config: *conf,
	}
	ProcessContainer[name] = p
	return nil
}

// Start start process async
func Start(name string) (err error) {
	p := ProcessContainer[name]
	if p == nil {
		return errors.New("process not exist")
	}

	err = p.Start()
	return err
}

// Stop stop process
func Stop(name string) (err error) {
	p := ProcessContainer[name]
	if p == nil {
		return errors.New("process not exist")
	}
	p.Stop()
	return err
}

// Release release process
func (p *Process) Release() {
	ProcessContainer[p.Name] = nil
}

// ExecSync exec command sync
func ExecSync(cmdstring string) (string, error) {
	args := strings.Split(cmdstring, " ")
	var cmd *exec.Cmd
	if len(args) >= 1 {
		cmd = exec.Command(args[0], args[1:]...)
	} else {
		cmd = exec.Command(args[0])
	}

	outbyte, err := cmd.Output()
	return string(outbyte), err
}

func init() {
	ProcessContainer = make(map[string]*Process, 1)
}

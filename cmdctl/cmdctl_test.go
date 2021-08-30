package cmdctl

import (
	"testing"
)

func TestNew(t *testing.T) {
	var pc ProcessConfig = ProcessConfig{
		ExecStart: "echo",
	}
	err := New("test", &pc)
	if err != nil {
		t.Errorf("test: new process 1 failed")
	}
	err = New("test", &pc)
	if err == nil {
		t.Errorf("test: new process failed")
	}

	err = New("test2", &pc)
	if err != nil {
		t.Errorf("test: new process 1 failed")
	}

	if !Exist("test2") {
		t.Errorf("test: exist test failed")
	}
}

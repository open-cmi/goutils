package devutil

import (
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var deviceIDFiles []string = []string{
	"/sys/class/dmi/id/product_uuid",
	"/sys/block/mmcblk0/device/serial",
}

// GetLinuxProductID func
func GetLinuxProductID() string {
	for _, filep := range deviceIDFiles {
		file, err := os.Open(filep)
		if err != nil {
			continue
		}

		data, _ := ioutil.ReadAll(file)
		deviceid := strings.Trim(string(data), " \r\n\t ")
		file.Close()
		return deviceid
	}
	return ""
}

// GetDarwinProductID func
func GetDarwinProductID() string {
	grep := exec.Command("grep", "IOPlatformSerialNumber")
	ioreg := exec.Command("ioreg", "-l")

	// Get ps's stdout and attach it to grep's stdin.
	pipe, err := ioreg.StdoutPipe()
	if err != nil {
		return ""
	}

	defer pipe.Close()
	grep.Stdin = pipe
	// Run ps first.
	ioreg.Start()

	// Run and get the output of grep.
	output, err := grep.Output()
	if err != nil {
		return ""
	}

	arr := strings.Split(string(output), "=")
	if len(arr) != 2 {
		return ""
	}
	deviceid := strings.Trim(arr[1], "\t\n\" ")
	return deviceid
}

// GetDeviceID func
func GetDeviceID() string {
	var deviceid string
	var sys string

	sys = runtime.GOOS
	if sys == "darwin" {
		deviceid = GetDarwinProductID()
	} else if sys == "windows" {
		deviceid = "windows_test"
	} else {
		deviceid = GetLinuxProductID()
	}

	return deviceid
}

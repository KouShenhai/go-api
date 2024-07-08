package device

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func GetSerialNo() (string, error) {
	// 优先CPU，其次MAC
	serial, err := getCPUSerial()
	if err != nil {
		return getMacAddress()
	}
	return serial, nil
}

func getCPUSerial() (string, error) {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		if strings.HasPrefix(line, "Serial") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				return fields[2], nil
			}
		}
	}
	return "", fmt.Errorf("CPU serial number not found")
}

func getMacAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, inter := range interfaces {
		if inter.Flags&net.FlagUp != 0 && inter.Flags&net.FlagLoopback == 0 {
			addr := inter.HardwareAddr.String()
			if addr != "" {
				return strings.ToUpper(addr), nil
			}
		}
	}
	return "", fmt.Errorf("no MAC address found")
}

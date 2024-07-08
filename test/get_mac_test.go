package test

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

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

func TestGetMacAddress(t *testing.T) {
	// 00:50:56:C0:00:01
	mac, err := getMacAddress()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("MAC Address:", mac)
}

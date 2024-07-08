package test

import (
	"fmt"
	"go-api/device"
	"testing"
)

func TestGetSerial(t *testing.T) {
	// 00:50:56:C0:00:01
	serialNo, err := device.GetSerialNo()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Serial No:", serialNo)
}

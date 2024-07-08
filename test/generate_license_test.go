package test

import (
	"fmt"
	"go-api/device"
	"testing"
)

func TestGenerateLicense(t *testing.T) {
	serialNo, err := device.GetSerialNo()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	expiryDate := "2024-12-31"
	filePath := "license.lic"
	// https://www.imaegoo.com/2020/aes-key-generator
	// 256-bit
	key := "cR0FfmGxU9dsNKpnxLAJ07Y96LUvDIhS"
	err = device.GenerateLicense(serialNo, expiryDate, key, filePath)
	if err != nil {
		fmt.Println("Error generating license:", err)
		return
	}
	fmt.Println("License generated successfully.")
}

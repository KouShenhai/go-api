package test

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"go-api/device"
	"io/ioutil"
	"testing"
	"time"
)

func decrypt(data []byte, passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func verifyLicense(key, filePath string) (bool, error) {
	encryptedLicense, err := ioutil.ReadFile(filePath)
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	decryptedLicense, err := decrypt(encryptedLicense, key)
	if err != nil {
		return false, err
	}
	var license device.License
	err = json.Unmarshal(decryptedLicense, &license)
	if err != nil {
		return false, err
	}
	currentSerialNo, err := device.GetSerialNo()
	if err != nil {
		return false, err
	}
	if license.SerialNo != currentSerialNo {
		return false, fmt.Errorf("SerialNo does not match")
	}
	expiryDate, err := time.Parse("2006-01-02", license.ExpiryDate)
	if err != nil {
		return false, err
	}
	if time.Now().After(expiryDate) {
		return false, fmt.Errorf("license has expired")
	}
	return true, nil
}

func TestVerifyLicense(t *testing.T) {
	key := "cR0FfmGxU9dsNKpnxLAJ07Y96LUvDIhS"
	filePath := "license.lic"
	isValid, err := verifyLicense(key, filePath)
	if err != nil {
		fmt.Println("Error verifying license:", err)
		return
	}
	if isValid {
		fmt.Println("License verification succeeded.")
	} else {
		fmt.Println("License verification failed.")
	}
}

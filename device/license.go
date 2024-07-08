package device

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type License struct {
	SerialNo   string `json:"serial_no"`
	ExpiryDate string `json:"expiry_date"`
}

func GenerateLicense(serialNo, expiryDate, key, filePath string) error {
	license := License{SerialNo: serialNo, ExpiryDate: expiryDate}
	licenseData, err := json.Marshal(license)
	if err != nil {
		return err
	}
	encryptedLicense, err := encrypt(licenseData, key)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	// 关闭文件
	_, err = file.Write(encryptedLicense)
	defer file.Close()
	defer os.Remove(filePath)
	return err
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(passphrase))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, data, nil), nil
}

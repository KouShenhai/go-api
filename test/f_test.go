package test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type License struct {
	MACAddress string `json:"mac_address"`
	ExpiryDate string `json:"expiry_date"`
}

func generateKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
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

func generateLicense(macAddress, expiryDate, key string) error {
	license := License{MACAddress: macAddress, ExpiryDate: expiryDate}
	licenseData, err := json.Marshal(license)
	if err != nil {
		return err
	}

	encryptedLicense, err := encrypt(licenseData, key)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("license.lic", encryptedLicense, 0644)
}

func main() {
	mac, err := getMacAddress()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	key, err := generateKey()
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}
	fmt.Println("Encryption Key (save this):", key)

	expiryDate := "2024-12-31"
	err = generateLicense(mac, expiryDate, key)
	if err != nil {
		fmt.Println("Error generating license:", err)
		return
	}
	fmt.Println("License generated successfully.")
}

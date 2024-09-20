package test

import (
	"fmt"
	"github.com/hirochachacha/go-smb2"
	"net"
	"testing"
)

func TestName1(t *testing.T) {
	conn, err := net.Dial("tcp", "xxx:30445")
	if err != nil {
		panic(err)
	}
	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "xxx",
			Password: "xxx",
		},
	}
	s, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	fs, err := s.Mount("//xxx/shareTest")
	if err != nil {
		panic(err)
	}
	matches, err := fs.Glob("*")
	if err != nil {
		panic(err)
	}
	for _, match := range matches {
		fmt.Println("match:", match)
	}
}

package localip

/*
Code provided by klasbo
https://github.com/TTK4145/Network-go/tree/master/network
*/
import (
	"net"
	"strings"
)

var localIP string

func LocalIP() (string, error) {
	if localIP == "" {
		conn, err := net.DialTCP("tcp4", nil, &net.TCPAddr{IP: []byte{8, 8, 8, 8}, Port: 53})
		if err != nil {
			return "", err
		}
		defer conn.Close()
		localIP = strings.Split(conn.LocalAddr().String(), ":")[0]
	}
	return localIP, nil
}

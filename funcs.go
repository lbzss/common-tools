package common_tools

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func GetNowTimeStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

func GetHostName() string {
	name, _ := os.Hostname()
	return name + "v1.0.2"
}

func GetLocalIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")

	if err != nil {
		log.Printf("net.dial err= %v", err)
		return "", fmt.Errorf("net.dial error")
	} else {
		fmt.Println(conn.LocalAddr().String())
		fmt.Println(conn.LocalAddr().Network())
		localIp := strings.Split(conn.LocalAddr().String(), ":")[0]
		conn.Close()
		return localIp, nil
	}

}

// func main() {
// 	fmt.Println(GetHostName())
// 	fmt.Println(GetLocalIp())
// 	fmt.Println(GetNowTimeStr())

// }

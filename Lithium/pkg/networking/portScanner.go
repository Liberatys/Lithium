package networking

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Scan struct {
	BaseIP string
	Scans  map[int]bool
}

func ScanPort(ip string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			return false
		}
	}
	defer conn.Close()
	return true
}

/**
Security measure for a new service
The code scans the entire given maschine and stores the state of a port
OPEN | CLOSE
So the executor of the service can check if everything is in order.
*/

func ScanPortRange(IP string, startPort int, endPort int, timeout time.Duration) Scan {
	scan := Scan{BaseIP: IP, Scans: make(map[int]bool)}
	for startPort < endPort {
		scan.Scans[startPort] = ScanPort(scan.BaseIP, startPort, timeout)
		if scan.Scans[len(scan.Scans)-1] == true {
			fmt.Println(strconv.Itoa(startPort) + "==" + strconv.FormatBool(scan.Scans[startPort]))
		}
		startPort++
	}
	return scan
}

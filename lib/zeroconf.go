package lib

import (
	"log"
	"net"
	"os"
)

func MyName() string {
	ip := MyIP()
	h, err := os.Hostname()
	if err == nil {
		as, err := net.LookupHost(h)
		if err != nil {
			return "localhost"
		}
		for _, a := range as {
			if a == ip.String() {
				return h + ".local"
			}
		}
	}
	return ip.String()
}

func MyIP() net.IP {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range ifaces {
		if i.Name == "lo" {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			break
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			return ip
		}
	}
	return net.IPv4(127, 0, 0, 1)
}

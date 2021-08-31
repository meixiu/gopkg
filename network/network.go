package network

import (
	"encoding/binary"
	"net"
	"net/http"
	"os"
	"strings"
)

// GetIP from header
func GetIP(r *http.Request) (ip string) {
	ip = "unknown"
	headers := []string{
		"X-Forwarded-For",
		"X-Real-Ip",
	}
	for _, v := range headers {
		ips := r.Header.Get(v)
		if ips == "" {
			continue
		}
		ipSlice := strings.Split(ips, ",")
		if len(ipSlice) < 0 {
			continue
		}
		ip = strings.Trim(ipSlice[0], " ")
		return ip
	}
	return ip
}

// InternalIP return internal ip.
func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

// gethostname()
func Gethostname() (string, error) {
	return os.Hostname()
}

// gethostbyname()
// Get the IPv4 address corresponding to a given Internet host name
func Gethostbyname(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}

// gethostbynamel()
// Get a list of IPv4 addresses corresponding to a given Internet host name
func Gethostbynamel(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		var ipstrs []string
		for _, v := range ips {
			if v.To4() != nil {
				ipstrs = append(ipstrs, v.String())
			}
		}
		return ipstrs, nil
	}
	return nil, err
}

// gethostbyaddr()
// Get the Internet host name corresponding to a given IP address
func Gethostbyaddr(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}

// ip2long()
// IPv4
func Ip2long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// long2ip()
// IPv4
func Long2ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}

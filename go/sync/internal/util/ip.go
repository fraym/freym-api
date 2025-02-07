package util

import (
	"fmt"
	"net"
	"os"
)

func OwnIp() (*net.IP, error) {
	host, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("unable to determine own ip: %w", err)
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, fmt.Errorf("unable to determine own ip: %w", err)
	}

	if len(ips) >= 1 {
		return &ips[0], nil
	}

	return nil, fmt.Errorf("unable to determine own ip")
}

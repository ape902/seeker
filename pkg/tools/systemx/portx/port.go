package portx

import (
	"fmt"
	"github.com/ape902/seeker/pkg/tools/systemx/netstatx"
)

const (
	TCP sockType = iota + 1
	TCP6
	UDP
	UDP6
)

type (
	sockType int8
)

func NewPorts() (map[int][]string, error) {
	ports := make(map[int][]string)
	s, err := sock(TCP, TCP6, UDP, UDP6)
	if err != nil {
		return ports, err
	}

	for i := 0; i < len(s); i++ {
		if s[i].Process == nil {
			continue
		}

		ports[s[i].Process.Pid] = append(ports[s[i].Process.Pid], fmt.Sprintf("%s;;%d", s[i].LocalAddr.IP, s[i].LocalAddr.Port))
	}

	return ports, nil
}

func sock(st ...sockType) ([]netstatx.SockTabEntry, error) {
	socks := make([]netstatx.SockTabEntry, 0)

	for _, t := range st {
		switch t {
		case TCP:
			s, err := netstatx.TCPSocks(func(e *netstatx.SockTabEntry) bool {
				return e.State == netstatx.Listen
			})
			if err != nil {
				return socks, err
			}

			socks = append(socks, s...)
		case TCP6:
			s, err := netstatx.TCP6Socks(func(e *netstatx.SockTabEntry) bool {
				return e.State == netstatx.Listen
			})

			if err != nil {
				return socks, err
			}

			socks = append(socks, s...)
		case UDP:
			s, err := netstatx.UDPSocks(func(e *netstatx.SockTabEntry) bool {
				return e.State == netstatx.Listen
			})

			if err != nil {
				return socks, err
			}

			socks = append(socks, s...)
		case UDP6:
			s, err := netstatx.UDP6Socks(func(e *netstatx.SockTabEntry) bool {
				return e.State == netstatx.Listen
			})

			if err != nil {
				return socks, err
			}

			socks = append(socks, s...)
		}
	}

	return socks, nil
}

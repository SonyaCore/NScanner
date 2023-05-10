package net

import (
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"sync"
	"time"
)

func (s Network) ScanPort(host string) (int, bool) {
	var conn net.Conn
	var err error

	address := host + ":" + strconv.Itoa(s.Port)
	conn, err = net.DialTimeout(s.Protocol, address, 20*time.Second)

	log.WithFields(
		log.Fields{
			"Address": address,
			"Port":    s.Port,
		},
	).Debug("Scanning", " ", address, s.Port)

	if err != nil {
		return 0, false
	}

	conn.Close()

	return s.Port, true

}

func (s Network) InitialScan(num int, host string, wg *sync.WaitGroup) []int {
	var Ports []int

	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(port int, host string) {
			s.Port = port
			scanner, open := s.ScanPort(host)

			if open {
				log.WithFields(
					log.Fields{
						"Address": host,
						"Port":    scanner,
					},
				).Info("Port", " ", scanner, " ", "Open")

				Ports = append(Ports, scanner)
			}
			wg.Done()

		}(i, host)
	}

	wg.Wait()

	return Ports
}

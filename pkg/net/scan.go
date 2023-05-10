package net

import (
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"time"
)

const DefaultRange = 2048

func (Scanner Network) ScanPort(host string) (int, bool) {
	var conn net.Conn
	var err error

	address := host + ":" + strconv.Itoa(Scanner.Port)
	conn, err = net.DialTimeout(Scanner.Protocol, address, 20*time.Second)

	log.WithFields(
		log.Fields{
			"Address": address,
			"Port":    Scanner.Port,
		},
	).Debug()

	if err != nil {
		return 0, false
	}

	defer conn.Close()

	return Scanner.Port, true

}

func (Scanner Network) InitialScan(num int, host string) []int {
	var Ports []int

	for i := 1; i <= num; i++ {
		Scanner.wg.Add(1)
		go func(port int, host string) {
			Scanner.Port = port
			scanner, open := Scanner.ScanPort(host)

			if open {
				log.WithFields(
					log.Fields{
						"Address": host,
						"Port":    scanner,
					},
				).Info("Port", " ", scanner, " ", "Open")

				Ports = append(Ports, scanner)
			}
			Scanner.wg.Done()

		}(i, host)
	}

	Scanner.wg.Wait()

	return Ports
}

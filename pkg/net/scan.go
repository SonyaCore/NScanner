package net

import (
	log "github.com/sirupsen/logrus"
	"net"
	"strconv"
	"sync"
	"time"
)

const DefaultRange = "1-2048"

func (Scanner Network) ScanPort(host string) (int, bool) {
	var connection net.Conn
	var err error

	address := host + ":" + strconv.Itoa(Scanner.Port)
	connection, err = net.DialTimeout(Scanner.Protocol, address, 20*time.Second)

	log.WithFields(
		log.Fields{
			"Address": address,
			"Port":    Scanner.Port,
		},
	).Debug()

	if err != nil {
		return 0, false
	}

	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(connection)

	return Scanner.Port, true

}

func (Scanner Network) InitialScan(startRangePorts, endRangePorts int, host string) []int {
	var Ports []int
	mutex := &sync.Mutex{}

	for i := startRangePorts; i <= endRangePorts; i++ {
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

				mutex.Lock()
				Ports = append(Ports, scanner)
				mutex.Unlock()

			}
			Scanner.wg.Done()

		}(i, host)
	}

	Scanner.wg.Wait()

	return Ports
}

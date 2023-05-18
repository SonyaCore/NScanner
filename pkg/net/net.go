package net

import (
	"Scanner/pkg/util"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sort"
)

func ReceiverHelper(w http.ResponseWriter, r *http.Request) {
	var Scanner Network

	err := json.NewDecoder(r.Body).Decode(&Scanner)
	if err != nil {
		return
	}

	startPort, _, _ := util.SplitPorts(Scanner.Range)

	if startPort == 0 {
		log.Warn("Scanner.Range is 0 setting port range to ", DefaultRange)

		Scanner.Range = DefaultRange
	}

	if Scanner.Host != "" {
		Scanner.SingleScan(w)
		return

	}

	if len(Scanner.HostList) >= 0 {
		Scanner.MultiScan(w)
		return
	}
}

func (Scanner Network) MultiScan(w http.ResponseWriter) {
	var Results []Result
	startPort, endPort, _ := util.SplitPorts(Scanner.Range)

	for _, host := range Scanner.HostList {
		result := Scanner.InitialScan(startPort, endPort, host)
		res := Result{Host: host}

		for _, x := range result {
			res.Ports = append(res.Ports, x)
			sort.Ints(res.Ports)
		}

		Results = append(Results, res)
	}

	data, _ := json.MarshalIndent(Results, " ", " ")

	_, err := w.Write(data)
	if err != nil {
		log.Error(err)
	}
}

func (Scanner Network) SingleScan(w http.ResponseWriter) {
	startPort, endPort, _ := util.SplitPorts(Scanner.Range)

	result := Scanner.InitialScan(startPort, endPort, Scanner.Host)
	res := Result{Host: Scanner.Host}

	for _, x := range result {
		res.Ports = append(res.Ports, x)
		sort.Ints(res.Ports)
	}

	data, _ := json.MarshalIndent(res, " ", " ")

	_, err := w.Write(data)
	if err != nil {
		log.Error(err)
	}
}

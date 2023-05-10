package net

import (
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

	if Scanner.Range == 0 {
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
	for _, host := range Scanner.HostList {
		result := Scanner.InitialScan(Scanner.Range, host)
		res := Result{Host: host}

		for _, x := range result {
			res.Ports = append(res.Ports, x)
			sort.Ints(res.Ports)
		}

		Results = append(Results, res)
	}

	data, _ := json.MarshalIndent(Results, " ", " ")

	w.Write(data)
}

func (Scanner Network) SingleScan(w http.ResponseWriter) {
	result := Scanner.InitialScan(Scanner.Range, Scanner.Host)

	res := Result{Host: Scanner.Host}

	for _, x := range result {
		res.Ports = append(res.Ports, x)
		sort.Ints(res.Ports)
	}

	data, _ := json.MarshalIndent(res, " ", " ")

	w.Write(data)
}

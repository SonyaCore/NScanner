package net

import (
	"Scanner/pkg/validate"
	"encoding/json"
	"net/http"
	"sort"
	"sync"
)

func ReceiverHelper(w http.ResponseWriter, r *http.Request) {
	var Scanner Network
	var wg sync.WaitGroup
	var Results []Result

	err := json.NewDecoder(r.Body).Decode(&Scanner)
	if err != nil {
		return
	}

	if Scanner.Range == 0 {
		Scanner.Range = 2048
	}

	err = validate.ValRange(Scanner.Range)
	if err != nil {
		_ = json.NewEncoder(w).Encode(
			&Error{err.Error(), http.StatusFailedDependency},
		)
		return
	}

	err = validate.ValProtocols(Scanner.Protocol)
	if err != nil {
		_ = json.NewEncoder(w).Encode(
			&Error{err.Error(), http.StatusFailedDependency},
		)
		return
	}

	if Scanner.Host != "" {
		SingleScan(w, Scanner, wg)
		return

	}

	if len(Scanner.HostList) >= 0 {
		MultiScan(w, Scanner, wg, Results)
		return
	}
}

func MultiScan(w http.ResponseWriter, Scanner Network, wg sync.WaitGroup, Results []Result) {
	for _, host := range Scanner.HostList {
		result := Scanner.InitialScan(Scanner.Range, host, &wg)
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

func SingleScan(w http.ResponseWriter, Scanner Network, wg sync.WaitGroup) {
	result := Scanner.InitialScan(Scanner.Range, Scanner.Host, &wg)

	res := Result{Host: Scanner.Host}

	for _, x := range result {
		res.Ports = append(res.Ports, x)
		sort.Ints(res.Ports)
	}

	data, _ := json.MarshalIndent(res, " ", " ")

	w.Write(data)
}

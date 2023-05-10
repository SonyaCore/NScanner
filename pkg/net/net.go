package net

import (
	"Scanner/pkg/validate"
	"encoding/json"
	"net/http"
	"sort"
	"sync"
)

func Receiver(w http.ResponseWriter, r *http.Request) {
	var Scanner Network
	var wg sync.WaitGroup
	//var result []int

	err := json.NewDecoder(r.Body).Decode(&Scanner)
	if err != nil {
		return
	}

	if Scanner.Range == 0 {
		Scanner.Range = 2048
	}

	err = validate.ValidateRange(Scanner.Range)
	if err != nil {
		_ = json.NewEncoder(w).Encode(
			&Error{err.Error(), http.StatusFailedDependency},
		)
		return
	}

	Res := []Result{}

	for _, host := range Scanner.Hostname {
		result := Scanner.InitialScan(Scanner.Range, host, &wg)cd
		res := Result{Host: host}

		for _, x := range result {
			res.Ports = append(res.Ports, x)
			sort.Ints(res.Ports)
		}

		Res = append(Res, res)
	}

	str, _ := json.MarshalIndent(Res, " ", " ")
	w.Write(str)
}

package net

import "sync"

type Network struct {
	Mode     string   `json:"mode"`
	Protocol string   `json:"protocol"`
	Host     string   `json:"host"`
	HostList []string `json:"list"`
	Port     int
	Range    string `json:"range"`
	wg       sync.WaitGroup
}

type Result struct {
	Host  string
	Ports []int
}

type Error struct {
	Message string
	Status  int
}

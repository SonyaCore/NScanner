package net

import "sync"

type Network struct {
	Protocol string   `json:"protocol"`
	Host     string   `json:"host"`
	HostList []string `json:"list"`
	Port     int
	Range    int `json:"range"`
	wg       sync.WaitGroup
}

type Result struct {
	Host  string
	Ports []int
}

type Results struct {
	Result []Result
}

type Error struct {
	Message string
	Status  int
}

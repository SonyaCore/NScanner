package net

type Network struct {
	Protocol string   `json:"protocol"`
	Host     string   `json:"host"`
	HostList []string `json:"list"`
	Port     int
	Range    int `json:"range"`
}

type Result struct {
	Host  string
	Ports []int
}

type Error struct {
	Message string
	Status  int
}

package net

type Network struct {
	Protocol string   `json:"protocol"`
	Hostname []string `json:"hostname"`
	Port     int
	Range    int `json:"range"`
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

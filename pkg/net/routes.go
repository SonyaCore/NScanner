package net

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/scan", ReceiverHelper)
}

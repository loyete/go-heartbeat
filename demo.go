package main

import (
	"github.com/loyete/go-heartbeat/heartbeat"
	"net/http"
)

//func main() {
//	r := gin.Default()
//	r.GET("/heartbeat", gin.WrapF(heartbeat.HeartbeatHandler))
//	r.Run(":8080")
//}

func main() {
	http.HandleFunc("/heartbeat", heartbeat.HeartbeatHandler)
	http.ListenAndServe(":8080", nil)
}

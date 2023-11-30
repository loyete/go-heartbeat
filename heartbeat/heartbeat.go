package heartbeat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type HeartbeatResponse struct {
	Status      string `json:"status"`
	RunningTime string `json:"running_time"`
}

func formatDuration(d time.Duration) string {
	d = d.Round(time.Second)
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second

	return fmt.Sprintf("%d天%d小时%d分%d秒", days, hours, minutes, seconds)
}

// HeartbeatHandler
// heartbeat
// http.HandleFunc("/heartbeat", heartbeat.HeartbeatHandler)
// use gin
// gin.GET("/heartbeat", gin.WrapF(heartbeat.HeartbeatHandler))
// return json:
//
//	{
//		"status": "OK",
//		"running_time": "25天3小时27分13秒"
//	}
func HeartbeatHandler(w http.ResponseWriter, r *http.Request) {
	uptime := formatDuration(time.Since(startTime))
	response := HeartbeatResponse{
		Status:      "running",
		RunningTime: uptime,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

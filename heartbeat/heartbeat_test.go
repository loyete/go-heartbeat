package heartbeat

import (
	"net/http"
	"testing"
	"time"
)

func TestHeartbeatHandler(t *testing.T) {
	http.HandleFunc("/heartbeat", HeartbeatHandler)
	http.ListenAndServe(":8080", nil)
}

func Test_formatDuration(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatDuration(tt.args.d); got != tt.want {
				t.Errorf("formatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

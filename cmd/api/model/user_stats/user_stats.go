package user_stats

import "net/http"

type UserStats struct {
	Ip                  string `json:"ip"`
	Path                string `json:"path"`
	ResponseTimeMs      int64  `json:"responseTimeMs"`
	Timestamp           string `json:"timestamp"`
	StatusCode          int    `json:"statusCode"`
	http.ResponseWriter `json:"-"`
}

func (stat *UserStats) WriteHeader(code int) {
	stat.StatusCode = code
	stat.ResponseWriter.WriteHeader(code)
}

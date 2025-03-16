package common

import (
	"fmt"
	"time"
)

// Common response structure for Jikan API
type JikanResponse struct {
	Data       interface{} `json:"data"`
	Pagination struct {
		LastVisiblePage int  `json:"last_visible_page"`
		HasNextPage     bool `json:"has_next_page"`
		CurrentPage     int  `json:"current_page"`
		Items           struct {
			Count   int `json:"count"`
			PerPage int `json:"per_page"`
			Total   int `json:"total"`
		} `json:"items"`
	} `json:"pagination"`
}

// APIResponse is our standardized API response format
type APIResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Count      int         `json:"count,omitempty"`
	TotalCount int         `json:"total_count,omitempty"`
	Page       int         `json:"page,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
	QueryTime  string      `json:"query_time,omitempty"`
}

// Success creates a successful API response
func Success(data interface{}, queryTime time.Duration) APIResponse {
	return APIResponse{
		Success:   true,
		Data:      data,
		QueryTime: fmt.Sprintf("%.3f seconds", queryTime.Seconds()),
	}
}

// Error creates an error API response
func Error(message string) APIResponse {
	return APIResponse{
		Success: false,
		Error:   message,
	}
}

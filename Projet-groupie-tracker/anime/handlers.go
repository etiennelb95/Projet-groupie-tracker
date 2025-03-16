package anime

import (
	"encoding/json"
	"fmt"
	"jikan-api-wrapper/common"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SearchHandler handles requests to search for anime by title
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		common.ErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get query parameters
	query := r.URL.Query().Get("q")
	if query == "" {
		common.ErrorResponse(w, "Search query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	limit := 25
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 && parsedLimit <= 100 {
			limit = parsedLimit
		}
	}

	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	// Measure execution time
	startTime := time.Now()

	// Perform the search
	results, totalCount, err := SearchAnime(query, limit, page)
	if err != nil {
		common.ErrorResponse(w, "Error searching anime: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := common.APIResponse{
		Success:    true,
		Data:       results,
		Count:      len(results),
		TotalCount: totalCount,
		Page:       page,
		TotalPages: (totalCount + limit - 1) / limit,
		QueryTime:  fmt.Sprintf("%.3f seconds", time.Since(startTime).Seconds()),
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAnimeHandler handles requests to get detailed information about a specific anime
func GetAnimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		common.ErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract anime ID from path
	path := strings.TrimPrefix(r.URL.Path, "/api/anime/")
	if path == "" {
		common.ErrorResponse(w, "Anime ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		common.ErrorResponse(w, "Invalid anime ID", http.StatusBadRequest)
		return
	}

	// Measure execution time
	startTime := time.Now()

	// Get anime details
	anime, err := GetAnimeDetails(id)
	if err != nil {
		common.ErrorResponse(w, "Error fetching anime details: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := common.Success(anime, time.Since(startTime))

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetTopAnimeHandler handles requests to get top anime
func GetTopAnimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		common.ErrorResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get query parameters
	limitStr := r.URL.Query().Get("limit")
	limit := 25
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 && parsedLimit <= 100 {
			limit = parsedLimit
		}
	}

	pageStr := r.URL.Query().Get("page")
	page := 1
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	filter := r.URL.Query().Get("filter")

	// Measure execution time
	startTime := time.Now()

	// Get top anime
	results, totalCount, err := GetTopAnime(filter, limit, page)
	if err != nil {
		common.ErrorResponse(w, "Error fetching top anime: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := common.APIResponse{
		Success:    true,
		Data:       results,
		Count:      len(results),
		TotalCount: totalCount,
		Page:       page,
		TotalPages: (totalCount + limit - 1) / limit,
		QueryTime:  fmt.Sprintf("%.3f seconds", time.Since(startTime).Seconds()),
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

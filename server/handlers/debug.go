package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("pong")
	}
}

func CpuHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		for time.Since(start) < 5*time.Second {
		}
		json.NewEncoder(w).Encode("CPU spike done")
	}
}

func DelayHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		durationStr := chi.URLParam(r, "duration")
		duration, err := strconv.Atoi(durationStr)
		if err != nil {
			http.Error(w, "Invalid duration", http.StatusBadRequest)
			return
		}
		time.Sleep(time.Duration(duration) * time.Second)
		fmt.Fprintf(w, "Delayed %d seconds\n", duration)
	}
}

func ErrorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func StatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		statusStr := chi.URLParam(r, "status")
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			http.Error(w, "Invalid status", http.StatusBadRequest)
			return
		}
		w.WriteHeader(status)
		http.Error(w, statusStr, status)
	}
}

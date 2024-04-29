package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// ========== CORS ============
	// set response sebagai json
	w.Header().Set("Content-Type", "application/json")

	// set endpoint ini boleh di akses dari mana saja
	// jika value nya = *, maka artinya semua host(server) boleh mengaksesnya
	// namun jika Access-Control-Allow-Origin = http://google.com
	// berarti hanya gogle.com yg boleh mengaksesnya
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// untuk mengizinkan method apa saja yang boleh di gunakan
	// by default, method GET dan POST sudah ter allow otomatis.
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS")

	// untuk mengizinkan header apa saja yang boleh
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Agent")
	// ======================

	// untuk mengirim http status code
	w.WriteHeader(http.StatusOK)

	// ========= UPDATE =========
	// setup random value value
	name := "User John doe"

	// ===========================

	response := map[string]interface{}{
		"name": name,
		"date": time.Now(),
	}

	log.Printf("type=%v method=%v path=%v", "[INFO]", r.Method, r.URL.Path)
	json.NewEncoder(w).Encode(response)
}

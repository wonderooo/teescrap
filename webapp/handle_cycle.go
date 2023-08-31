package webapp

import (
	"encoding/json"
	"net/http"
)

type CycleHandler struct {
	db *Db
}

type Cycle struct {
	Breeds []string `json:"breeds"`
	Styles []string `json:"styles"`
	Tags   []string `json:"tags"`
}

func HandleCreateCycle(db *Db) http.Handler {
	return &CycleHandler{
		db: db,
	}
}

func (h *CycleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var c Cycle
	if r.Method == "POST" {
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			internalServerError(w, r)
			return
		}

		h.db.cycles.append(c)

		encoded, err := json.Marshal(h.db.cycles.m)
		if err != nil {
			internalServerError(w, r)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(encoded)
	} else {
		notFound(w, r)
	}
}

package webapp

import (
	"encoding/json"
	"log"
	"net/http"
)

type ListHandler struct {
	db *Db
}

type List struct {
	Resp map[int]int `json:"response"`
}

func HandleList(db *Db) http.Handler {
	return &ListHandler{
		db: db,
	}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var ret List
		m := make(map[int]int)
		for k, v := range h.db.cycles.all() {
			m[k+1] = len(v.Jobs)
		}
		ret.Resp = m

		encoded, err := json.Marshal(ret)
		if err != nil {
			log.Fatal("Could not marshall List struct:", err)
			internalServerError(w, r)
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(encoded)
		if err != nil {
			internalServerError(w, r)
		}
	} else {
		notFound(w, r)
	}
}
package webapp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wonderooo/teescrap/uploader"
)

type CycleHandler struct {
	db *Db
}

type Cycle struct {
	Breeds []string       `json:"breeds"`
	Styles []string       `json:"styles"`
	Tags   []string       `json:"tags"`
	Jobs   []uploader.Job `json:"jobs"`
	Status string         `json:"status"`
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

		c.Status = "pending"
		c.Jobs = getJobs(c.Breeds, c.Styles, c.Tags)
		h.db.cycles.append(c)

		log.Println(c)
		upl := uploader.New(true, &c.Jobs)
		go upl.Run()

		w.WriteHeader(http.StatusCreated)
	} else {
		notFound(w, r)
	}
}

func getJobs(breeds []string, styles []string, tags []string) []uploader.Job {
	ret := make([]uploader.Job, 0)

	for _, breed := range breeds {
		for _, style := range styles {
			desc := fmt.Sprintf("%s painted in style of %s", breed, style)
			ret = append(ret,
				uploader.NewJob(
					"shared/rng.jpeg",
					tags[0],
					desc,
					uploader.ColorChoices{},
					tags...,
				),
			)
		}
	}

	return ret
}

const cycleLi = `<li class="list-group-item d-flex justify-content-between align-items-start pb-4">
<div class="ms-2 me-auto">
  <div class="fw-bold">#:param Cycle</div>
      <div class="dropdown">
          <button class="btn btn-info dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
          Jobs
          </button>
          <ul class="dropdown-menu">
            :param
          </ul>
      </div> 
  </div>
<span class="badge bg-danger rounded-pill">Failed</span>
</li>`

const jobLi = `<li>
<button class="dropdown-item btn" data-bs-toggle="modal" data-bs-target="#exampleModal">
    #:param Job
    <span class="badge bg-danger rounded-pill">Failed</span>
</button>
</li>`

function createCycle() {
    let breeds = document.getElementById("breeds-inp")?.value.split(",")
    let styles = document.getElementById("styles-inp")?.value.split(",")
    let tags = document.getElementById("tags-inp")?.value.split(",")

    fetch("http://localhost:8080/cycles", {
        method: "POST",
        body: JSON.stringify({
            breeds: breeds,
            styles: styles,
            tags: tags,
        })
    }
    ).then((response) => response.json())
    .then((json) => console.log(json));
}

function refresh() {
    fetch("http://localhost:8080/refresh", {
        method: "GET"   
    }
    ).then((response) => response.json())
    .then((json) => addCycles(json))
}

function addCycles(json) {
    let map = new Map(Object.entries(json["response"]))

    let cyclesList = document.getElementById("cycles-list")
    for (let [k, v] of map) {
        let ret = sprintf(cycleLi, k, addJobs(v))
        let node = document.createRange().createContextualFragment(ret)
        cyclesList.append(node)
    }
}

function addJobs(size) {
    let ret = ""
    for (let i=1; i<=size; i++) {
        ret += sprintf(jobLi, i)
    }

    return ret
}

function sprintf(format, ...values) {
    return values.reduce((carry, current) => carry.replace(/:param/, current), format);
}
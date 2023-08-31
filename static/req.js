function createCycle() {
    fetch("http://localhost:8080/cycles", {
        method: "POST",
        body: JSON.stringify({
            breeds: ["dupa", "dupa"],
            styles: ["dupa", "dupa"],
            tags: ["dupa", "dupa"]
        })
    }
    ).then((response) => response.json())
    .then((json) => console.log(json));
}
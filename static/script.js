function sendAnswer(answer) {
    fetch("/answer", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            answer: answer
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Something went wrong");
        }

        return response.text();
    })
    .then(message => {
        document.getElementById("response").innerText =
            message + " ❤️";
    })
    .catch(error => {
        document.getElementById("response").innerText =
            "Unable to send your answer. Please try again.";
    });
}
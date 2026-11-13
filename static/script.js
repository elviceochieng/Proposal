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
    .then(() => {
        if (answer === "Yes") {
            showCelebration();
        } else {
            document.getElementById("response").innerText =
                "Thank you for being honest, Deeh. ❤️";
        }
    })
    .catch(() => {
        document.getElementById("response").innerText =
            "Unable to send your answer. Please try again.";
    });
}

function showCelebration() {
    const response = document.getElementById("response");

    response.innerHTML = `
        <div class="celebration-message">
            <h2>YAAAAAY! ❤️🥹</h2>
            <p>
                Deeh ❤️🥹 You just made me the happiest person alive!
                I promise to make every moment with you special,
                even though we're financially unstable 😂.
                Cheers! 🥂 to our beautiful journey together! 💕
            </p>
        </div>
    `;

    response.classList.add("show");

    createConfetti();
}

function createConfetti() {
    const colors = ["#ff1493", "#70045ef3", "#ffd700", "#ff0000", "#9370db"];

    for (let i = 0; i < 100; i++) {
        const ribbon = document.createElement("div");

        ribbon.classList.add("confetti");

        ribbon.style.left = Math.random() * 100 + "vw";
        ribbon.style.backgroundColor =
            colors[Math.floor(Math.random() * colors.length)];

        ribbon.style.animationDuration =
            (Math.random() * 3 + 2) + "s";

        ribbon.style.animationDelay =
            Math.random() * 2 + "s";

        document.body.appendChild(ribbon);

        setTimeout(() => {
            ribbon.remove();
        }, 7000);
    }
}
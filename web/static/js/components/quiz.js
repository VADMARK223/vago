(() => {
    const quizEl = document.getElementById("quiz");
    if (!quizEl) {
        console.warn("[quizEl] elements not found");
        return;
    }

    const questionID = Number(quizEl.dataset.questionId);
    const buttons = document.querySelectorAll(".answer");
    const result = document.getElementById("result");
    const explanation = document.getElementById("explanation");

    const continueBtn = document.getElementById("continueBtn")
    continueBtn.addEventListener("click", () => {
        window.location.href = "/quiz";
    })

    buttons.forEach(btn => {
        btn.addEventListener("click", async () => {
            const answerID = Number(btn.dataset.id);

            // Очищаем предыдущие подсветки
            buttons.forEach(b => {
                b.classList.remove("correct", "wrong");
            });

            const resp = await fetch("/quiz/check", {
                method: "POST",
                headers: {"Content-Type": "application/json"},
                body: JSON.stringify({
                    question_id: Number(questionID),
                    answer_id: answerID
                })
            });
            const data = await resp.json();

            result.classList.remove("hidden");

            if (data.correct) {
                btn.classList.add("correct");

                result.classList.remove("bad");
                result.classList.add("ok");

                if (data.explanation) {
                    explanation.classList.remove("hidden");
                    result.innerText = "Правильно!";
                    explanation.innerText = data.explanation;
                    continueBtn.style.visibility = "visible";
                } else {
                    result.innerText = "Правильно!";
                    setTimeout(() => {
                        window.location.href = "/quiz";
                    }, 700);
                }
            } else {
                btn.classList.add("wrong");
                result.innerText = "Неправильно 😞";
                result.classList.remove("ok");
                result.classList.add("bad");
                explanation.classList.add("hidden");
                continueBtn.style.visibility = "hidden";
            }
        });
    });
})();
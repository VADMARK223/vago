(() => {
    const testEl = document.getElementById("quiz");
    if (!testEl) {
        console.warn("[quizEl] elements not found");
        return;
    }

    const questionID = Number(testEl.dataset.questionId);
    const buttons = document.querySelectorAll(".answer");
    if (buttons.length === 0) {
        console.warn("[buttons] elements with class='answer' not found");
        return;
    }
    const result = document.getElementById("result");
    const explanation = document.getElementById("explanation");

    const continueBtnNew = document.getElementById("continueBtn")
    if (!continueBtnNew) {
        console.warn("[continueBtn] elements not found");
        return;
    }

    let finished = false;

    continueBtnNew.addEventListener("click", () => {
        window.location.href = "/test";
    })

    const disableAnswers = (state) => {
        buttons.forEach(b=>b.disabled = state)
    }

    buttons.forEach(btn => {
        btn.addEventListener("click", async () => {
            if (finished) return;

            const answerID = Number(btn.dataset.id);

            // Очищаем предыдущие подсветки
            buttons.forEach(b => {
                b.classList.remove("correct", "wrong");
            });

            disableAnswers(true)

            result.classList.add("hidden");
            explanation.classList.add("hidden");
            result.classList.remove("ok", "bad");

            try {
                const resp = await fetch("/test/check", {
                    method: "POST",
                    headers: {"Content-Type": "application/json"},
                    body: JSON.stringify({
                        question_id: Number(questionID),
                        answer_id: answerID
                    })
                });

                if (!resp.ok) {
                    throw new Error("Server error");
                }

                const data = await resp.json();
                result.classList.remove("hidden");
                if (data.correct) {
                    showCorrect(btn, data.explanation)
                } else {
                    showWrong(btn)
                }

            } catch (err) {
                showError()
            }
        });
    });

    function showCorrect (btn, dataExplanation) {
        finished = true;
        btn.classList.add("correct");

        result.classList.add("ok");
        result.innerText = "Правильно!";

        if (dataExplanation) {
            explanation.classList.remove("hidden");
            explanation.innerText = dataExplanation;
        }

        continueBtnNew.hidden = false
    }

    function showWrong (btn) {
        btn.classList.add("wrong");
        disableAnswers(false)

        result.classList.add("bad");
        result.innerText = "Неправильно 😞";

        continueBtnNew.hidden = true
    }

    function showError () {
        disableAnswers(false)
        explanation.classList.add("hidden");
        result.classList.add("bad");
        result.innerText = "Ошибка проверки ответа 😕";
    }
})();
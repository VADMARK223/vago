(() => {
    const questionsSeedBtn = document.getElementById("questionsSeedBtn")
    const questionsSeedBtnText = questionsSeedBtn.textContent

    if (!questionsSeedBtn) {
        console.warn("[questionsSeedBtn] elements not found");
        return;
    }

    questionsSeedBtn.addEventListener("click", async () => {
        questionsSeedBtn.disabled = true
        questionsSeedBtn.textContent = "Running..."

        try {
            const resp = await fetch("/run_questions_seed", {method: "POST"})
            const data = await resp.json();
            if (!resp.ok) {
                alert('Ошибка: ' + data.message);
                return;
            }

            location.reload();

        } finally {
            resetQuestionsSeedBtn()
        }

        function resetQuestionsSeedBtn() {
            questionsSeedBtn.disabled = false;
            questionsSeedBtn.textContent = questionsSeedBtnText
        }
    })
})();
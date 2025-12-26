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
            await runSeed("/run_questions_seed", {method: "POST"})
            location.reload()
        } catch (err) {
            alert("Error: " + err.message)
        } finally {
            resetQuestionsSeedBtn()
        }

        function resetQuestionsSeedBtn() {
            questionsSeedBtn.disabled = false;
            questionsSeedBtn.textContent = questionsSeedBtnText
        }
    })

    async function runSeed(url, options) {
        const resp = await fetch(url, options);
        const data = await resp.json();
        if (!resp.ok) {
            throw new Error(data.error || resp.statusText);
        }

        return data
    }
})();
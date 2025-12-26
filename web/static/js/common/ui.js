(() => {
    document.querySelectorAll(".clear-btn[data-clear]").forEach(btn => {
        btn.addEventListener("click", () => {
            const inp = btn.closest(".input-wrap").querySelector("input")
            inp.value = ""
            inp.dispatchEvent(new Event("input", { bubbles: true }));
        });
    });

    document.querySelectorAll("input[data-required='true']").forEach(input => {
        input.required = true;
    });
})();
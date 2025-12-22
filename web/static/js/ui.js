document.addEventListener("DOMContentLoaded", () => {
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

    // -----------------------------------------------------scroll-----------------------------------------
    const container = document.getElementById('scrollable');
    const btn = document.getElementById('scrollTopBtn');

    console.log("scrollable container", container)
    console.log("scrollable btn", btn)

    function updateScrollButton() {
        const hasScroll = container.scrollHeight > container.clientHeight;
        const scrolled = container.scrollTop > 50;

        if (hasScroll && scrolled) {
            btn.classList.add('visible');
        } else {
            btn.classList.remove('visible');
        }
    }

    container.addEventListener('scroll', updateScrollButton);

    btn.addEventListener('click', () => {
        container.scrollTo({
            top: 0,
            behavior: 'smooth'
        });
    });

    // Проверка при загрузке
    updateScrollButton();
});
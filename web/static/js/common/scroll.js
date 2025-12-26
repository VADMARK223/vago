(()=>{
    const container = document.getElementById('scrollable');
    const btn = document.getElementById('scrollTopBtn');

    if (!container) {
        console.warn("[scrollable] elements not found");
        return;
    }

    if (!btn) {
        console.warn("[scrollTopBtn] elements not found");
        return;
    }

    function updateScrollButton() {
        const hasScroll = container.scrollHeight > container.clientHeight;
        const scrolled = container.scrollTop > 50;
        btn.classList.toggle('visible', hasScroll && scrolled);
    }

    let ticking = false;

    container.addEventListener('scroll', () => {
        if (!ticking) {
            window.requestAnimationFrame(() => {
                updateScrollButton();
                ticking = false;
            });
            ticking = true;
        }
    });

    btn.addEventListener('click', () => {
        container.scrollTo({
            top: 0,
            behavior: 'smooth'
        });
    });

    updateScrollButton();
})();
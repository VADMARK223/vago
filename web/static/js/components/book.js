(() => {
    const STORAGE_KEY = 'bookHistory';
    const backBtn = document.querySelector('[data-back]');

    if (!backBtn) {
        console.warn("[scrollable] elements not found");
        return;
    }

    const isBookPage = (url = location.pathname) =>
        url.startsWith('/book');

    const normalizeUrl = () =>
        location.pathname + location.search; // БЕЗ hash

    const loadHistory = () =>
        JSON.parse(sessionStorage.getItem(STORAGE_KEY) || '[]');

    const saveHistory = (history) =>
        sessionStorage.setItem(STORAGE_KEY, JSON.stringify(history));

    const updateBackButton = () => {
        const history = loadHistory();
        if (history.length <= 1) {
            backBtn.classList.add('disabled');
        } else {
            backBtn.classList.remove('disabled');
        }
    };

    const pushToHistory = () => {
        if (!isBookPage()) {
            sessionStorage.removeItem(STORAGE_KEY);
            return;
        }

        const history = loadHistory();
        const current = normalizeUrl();

        if (history[history.length - 1] !== current) {
            history.push(current);
            saveHistory(history);
        }

        updateBackButton();
    };

    // Клик по кнопке Назад
    backBtn.addEventListener('click', (e) => {
        e.preventDefault();
        if (backBtn.classList.contains('disabled')) return;

        const history = loadHistory();
        history.pop(); // текущая
        const prev = history[history.length - 1];

        saveHistory(history);
        updateBackButton();

        if (prev) {
            location.href = prev;
        }
    });

    // Игнорируем переходы только по якорям
    window.addEventListener('hashchange', () => {
        updateBackButton();
    });

    // Навигация браузерной кнопкой
    window.addEventListener('popstate', () => {
        pushToHistory();
    });

    // Первый заход / переход
    document.addEventListener('DOMContentLoaded', () => {
        pushToHistory();
    });

    // Сброс истории при уходе из книги
    document.addEventListener('click', (e) => {
        const link = e.target.closest('a[href]');
        if (!link) return;

        const href = link.getAttribute('href');
        if (!href || href.startsWith('#')) return;

        // абсолютный URL
        const url = new URL(href, location.origin);

        // уходим НЕ в /book
        if (!url.pathname.startsWith('/book')) {
            sessionStorage.removeItem(STORAGE_KEY);
        }
    });
})();
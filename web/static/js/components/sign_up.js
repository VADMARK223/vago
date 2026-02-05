(()=>{
    const loginInput = document.querySelector('input[name="login"]')
    if (!loginInput) {
        console.warn("[login] elements not found");
        return;
    }
    const passwordInput = document.querySelector('input[name="password"]')
    if (!passwordInput) {
        console.warn("[password] elements not found");
        return;
    }
    const usernameInput = document.querySelector('input[name="username"]')
    if (!usernameInput) {
        console.warn("[username] elements not found");
        return;
    }
    const submitBtn = document.getElementById("submitBtn")
    if (!submitBtn) {
        console.warn("[submitBtn] elements not found");
        return;
    }

    const select = document.getElementById("roleSelect");
    if (!select) {
        console.warn("[select] elements not found");
        return;
    }

    const hint = document.getElementById("roleHint");
    if (!hint) {
        console.warn("[hint] elements not found");
        return;
    }

    const ROLE_HINT = {
        moderator: "Доступны все разделы с ограниченными правами.",
        user: "Некоторые разделы не доступы."
    };

    loginInput.addEventListener("input", () => {
        const v = loginInput.value.trim();

        if (v) {
            usernameInput.value = `User${v}`.slice(0, 30);
        }

        updateState()
    })

    passwordInput.addEventListener("input", updateState);
    usernameInput.addEventListener("input", updateState);

    function updateState() {
        const login = loginInput.value.trim();
        const password = passwordInput.value.trim();
        const username = usernameInput.value.trim();
        setDisabledSubmitBtn(login, password, username);
    }

    function setDisabledSubmitBtn(login, password, username) {
        submitBtn.disabled = login === "" || password === "" || username === "";
    }

    updateState();

    function updateHint(role) {
        hint.textContent = ROLE_HINT[role];
    }

    // начальное состояние
    updateHint(select.value);

    // при изменении
    select.addEventListener("change", function () {
        updateHint(this.value);
    });
})();
(() => {
    const loginInput = document.querySelector('input[name="login"]');
    if (!loginInput) {
        console.warn("[login] elements not found");
        return;
    }

    const passwordInput = document.querySelector('input[name="password"]');
    if (!passwordInput) {
        console.warn("[password] elements not found");
        return;
    }
    const submitBtn = document.getElementById("submitBtn")
    if (!submitBtn) {
        console.warn("[submitBtn] elements not found");
        return;
    }

    loginInput.addEventListener("input", updateState);
    passwordInput.addEventListener("input", updateState);

    function updateState() {
        const login = loginInput.value.trim();
        const password = passwordInput.value.trim();
        setDisabledSubmitBtn(login, password);
    }

    function setDisabledSubmitBtn(login, password) {
        submitBtn.disabled = login === "" || password === "";
    }

    updateState();
})();
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
    const emailInput = document.querySelector('input[name="email"]')
    if (!emailInput) {
        console.warn("[email] elements not found");
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

    loginInput.addEventListener("input", () => {
        const v = loginInput.value.trim();

        if (v) {
            passwordInput.value = v;
            emailInput.value = `${v}@mail.ru`;
            usernameInput.value = `User${v}`.slice(0, 30);
        }

        updateState()
    })

    passwordInput.addEventListener("input", updateState);
    emailInput.addEventListener("input", updateState);
    usernameInput.addEventListener("input", updateState);

    function updateState() {
        const login = loginInput.value.trim();
        const password = passwordInput.value.trim();
        const email = emailInput.value.trim();
        const username = usernameInput.value.trim();
        setDisabledSubmitBtn(login, password, email, username);
    }

    function setDisabledSubmitBtn(login, password, email, username) {
        submitBtn.disabled = login === "" || password === "" || email === "" || username === "";
    }

    updateState();
})();
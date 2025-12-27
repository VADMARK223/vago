(()=>{
    window.saveAdminSection = async function (path) {
        localStorage.setItem("admin:lastSection", path);
    }

    const last = localStorage.getItem("admin:lastSection");
    if (last && location.pathname === "/admin") {
        window.location.replace(last);
    }
})();

document.addEventListener("DOMContentLoaded", () => {
    const last = localStorage.getItem("admin:lastSection");

    if (last && last !== location.pathname) {
        window.location.replace(last);
    } else {

    }
});
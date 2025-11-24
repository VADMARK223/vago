const HOSTNAME = window.location.hostname;

export function initChat(cfg) {
    const myUserId = cfg.userId;
    const token = cfg.token;
    const port= cfg.port
    console.log(`Init chat userid: ${myUserId}, PORT:${port}, token: ${token}`)

    if (!token) {
        cfg.status.textContent = "❌ Token not found. Please log in.";
        return;
    }

    const host = `ws://${HOSTNAME}:${port}/ws`;
    const url = host + "?token=" + token;

    const socket = new WebSocket(url);

    socket.onopen = () =>
        (cfg.status.textContent = "🟢 Connected (" + host + ")");
    socket.onclose = () => (cfg.status.textContent = "🔴 Disconnected");
    socket.onerror = () => (cfg.status.textContent = "❌ Error");

    socket.onmessage = (event) => {
        try {
            const msg = JSON.parse(event.data);

            if (msg.type === "message") {
                const isMine =
                    String(msg.userId) === String(myUserId);
                const text = isMine
                    ? msg.text
                    : `${msg.userId}: ${msg.text}`;
                addMessage(text, isMine);
            }
        } catch (e) {
            console.error("Bad JSON:", e);
        }
    };

    cfg.sendBtn.onclick = sendMessage;
    cfg.input.addEventListener("keydown", (e) => {
        if (e.key === "Enter") {
            e.preventDefault();
            sendMessage();
        }
    });

    function sendMessage() {
        const text = cfg.input.value.trim();
        if (!text || socket.readyState !== WebSocket.OPEN) return;

        socket.send(
            JSON.stringify({
                type: "message",
                text,
            })
        );

        cfg.input.value = "";
    }

    function addMessage(text, isMine) {
        const div = document.createElement("div");
        div.textContent = text;

        if (isMine) {
            div.classList.add("chat-my-message");
        }

        cfg.messages.appendChild(div);
        cfg.messages.scrollTop = cfg.messages.scrollHeight;
    }
}
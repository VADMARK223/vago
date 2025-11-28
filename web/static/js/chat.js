const HOSTNAME = window.location.hostname;

export function initChat(cfg) {
    const myUserId = cfg.userId;
    const token = cfg.token;
    const port = cfg.port

    console.log(`Init chat userid: ${myUserId}, PORT:${port}, token: ${token}`)
    const msgInput = cfg.input
    const sendBtn = cfg.sendBtn

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
            parseAndAddMessage(messagesDiv, msg, myUserId)
        } catch (e) {
            console.error("Bad JSON:", e);
        }
    };

    sendBtn.onclick = sendMessage;
    msgInput.addEventListener("keydown", (e) => {
        if (e.key === "Enter") {
            e.preventDefault();
            sendMessage();
        }

        updateEnableSendBtn()
    });

    msgInput.addEventListener("input", () => {
        updateEnableSendBtn()
    })

    function sendMessage() {
        const text = msgInput.value.trim();
        if (!text || socket.readyState !== WebSocket.OPEN) return;

        socket.send(
            JSON.stringify({
                type: "message",
                text,
            })
        );

        msgInput.value = "";
        updateEnableSendBtn()
    }

    function updateEnableSendBtn() {
        sendBtn.disabled = msgInput.value === ""
    }

    updateEnableSendBtn()
}

export function parseAndAddMessage(messagesDiv, msg, myUserId) {
    console.log("Row message", msg)
    if (msg.type === "message") {
        const isMine = String(msg.author) === String(myUserId);
        addMessage(messagesDiv, msg.username, msg.body, isMine, msg.sent_at, msg.author);
    } else {
        console.error(`Unknown message type:"${msg.type}"`)
    }
}

function addMessage(messagesDiv, username, text, isMine, sentAt, author) {
    const wrapper = document.createElement("div");
    wrapper.classList.add("chat-message");

    if (isMine) {
        wrapper.classList.add("chat-my-message");
    }

    // if (!isMine) {
        const header = document.createElement("div");
        header.classList.add("chat-message-header");
        header.textContent = `${author} • ${username}`;
        wrapper.appendChild(header);
    // }

    const body = document.createElement("div");
    body.classList.add("chat-message-body");
    body.textContent = text;
    wrapper.appendChild(body);

    const footer = document.createElement("div");
    footer.classList.add("chat-message-footer");
    footer.textContent = `${sentAt}`;
    wrapper.appendChild(footer);

    messagesDiv.appendChild(wrapper);
    messagesDiv.scrollTop = messagesDiv.scrollHeight;
}
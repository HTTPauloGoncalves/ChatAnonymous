document.addEventListener("DOMContentLoaded", () => {
    const modalOverlay = document.getElementById("modalOverlay");
    const modalCodnome = document.getElementById("modalCodnome");

    function abrirModal() {
        modalOverlay.style.display = "block";
        modalCodnome.style.display = "block";
        setTimeout(() => document.getElementById("codinomeInput").focus(), 50);
    }

    function fecharModal() {
        modalOverlay.style.display = "none";
        modalCodnome.style.display = "none";
    }

    modalOverlay.onclick = fecharModal;

    let CODINOME = localStorage.getItem("codinome");

    if (!CODINOME || CODINOME.trim() === "") {
        abrirModal();
    }

    const params = new URLSearchParams(window.location.search);
    const ROOM_ID = params.get("room");
    const PASSWORD = params.get("password");

    document.getElementById("room-title").innerText = "Sala #" + ROOM_ID;

    let ws = new WebSocket(`ws://localhost:8080/ws?room=${ROOM_ID}&password=${PASSWORD}`);

    ws.onopen = () => {
        document.getElementById("room-users").innerText = "Conectado";
    };

    ws.onclose = () => {
        document.getElementById("status-indicator").style.background = "#ef4444";
        document.getElementById("room-users").innerText = "Desconectado";
    };

   ws.onmessage = (event) => {
    const msg = JSON.parse(event.data);

    const username = msg.username ?? msg.data?.username ?? "Anônimo";
    const message = msg.message ?? msg.data?.message;

    if (message) {
        adicionarMensagem(message, false, username);
    }
};

    document.getElementById("messageInput").addEventListener("keypress", e => {
        if (e.key === "Enter") enviarMensagem();
    });

    window.enviarMensagem = () => {
        const input = document.getElementById("messageInput");
        const text = input.value.trim();
        if (text === "" || ws.readyState !== WebSocket.OPEN) return;

        let nomeAtual = localStorage.getItem("codinome");

        if (!nomeAtual || nomeAtual.trim() === "") {
            abrirModal();
            return;
        }

        ws.send(JSON.stringify({ username: nomeAtual, message: text }));
        adicionarMensagem(text, true, nomeAtual);
        input.value = "";
    };

    window.salvarCodinome = () => {
        const nome = document.getElementById("codinomeInput").value.trim();
        if (nome === "") return;
        localStorage.setItem("codinome", nome);
        fecharModal();
    };

    window.nomeAleatorio = () => {
        const nomes = ["Nebulosa", "ZeroAlpha", "Sombra", "Violeta", "GhostFox"];
        document.getElementById("codinomeInput").value = nomes[Math.floor(Math.random() * nomes.length)];
    };
});

function adicionarMensagem(text, own, senderName) {
    const messages = document.getElementById("messages");
    const div = document.createElement("div");
    div.className = own ? "message own" : "message";

    const time = new Date().toLocaleTimeString("pt-BR", {
        hour: "2-digit",
        minute: "2-digit"
    });

    div.innerHTML = `
        <div class="message-avatar">${senderName[0]}</div>
        <div class="message-content">
            <div class="message-sender">${senderName}</div>
            <div class="message-bubble">${text}</div>
            <div class="message-time">${time}</div>
        </div>
    `;

    messages.appendChild(div);
    messages.scrollTop = messages.scrollHeight;
}

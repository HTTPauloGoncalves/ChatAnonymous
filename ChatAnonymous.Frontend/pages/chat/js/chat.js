let ws = null;
const params = new URLSearchParams(window.location.search);

document.addEventListener("DOMContentLoaded", () => {
    const modalOverlay = document.getElementById("modalOverlay");
    const modalCodnome = document.getElementById("modalCodnome");
    const isRandom = params.get("random") === "true";

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

    const ROOM_ID = params.get("room");
    const PASSWORD = params.get("password");

    if (isRandom) {
        document.getElementById("room-title").innerText = "Chat Aleatório";
        document.getElementById("copyBtn").style.display = "none";

        ws = new WebSocket("ws://192.168.0.109:8080/ws/random");

        ws.onopen = () => {
            document.getElementById("room-users").innerText = "Procurando alguém";
            document.getElementById("divloader").style.display = "block"
            ws.send(JSON.stringify({ type: "join_random" }));
        };
    } else {
        document.getElementById("room-title").innerText = "Sala #" + ROOM_ID;

        ws = new WebSocket(
            `ws://192.168.0.109:8080/ws?room=${ROOM_ID}&password=${PASSWORD}`
        );

        ws.onopen = () => {
            document.getElementById("room-users").innerText = "Conectado";
            document.getElementById("divloader").style.display = "none"
        };
    }


    ws.onclose = () => {
        document.getElementById("status-indicator").style.background = "#ef4444";
        document.getElementById("room-users").innerText = "Desconectado";
        document.getElementById("divloader").style.display = "none";
    };

   ws.onmessage = (event) => {
    const msg = JSON.parse(event.data);

    if (msg.type === "system") {

        if (isRandom && msg.message === "Um usuário entrou.") {
            document.getElementById("room-users").innerText = "Conectado";
            document.getElementById("divloader").style.display = "none";
        }

        if (isRandom && msg.message === "Sala encerrada.") {
            document.getElementById("status-indicator").style.background = "#ef4444";
            document.getElementById("room-users").innerText = "Desconectado";
            document.getElementById("divloader").style.display = "none";

            proximoRandom()
        }

        adicionarMensagemSistema(msg.message);
        return;
    }

    if (msg.type === "chat") {
        adicionarMensagem(msg.message, false, msg.username);
    }

};

document.querySelector('emoji-picker')
        .addEventListener('emoji-click', event => document.getElementById("messageInput").value += event.detail.emoji.unicode);


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

        ws.send(JSON.stringify({
            type: "chat",
            username: nomeAtual,
            message: text
        }));

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

function saiuMensagem(text, own, senderName) {
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

function adicionarMensagemSistema(text) {
    const messages = document.getElementById("messages");
    const div = document.createElement("div");

    div.className = "system-message";

    div.innerHTML = `
        <div class="system-text"><em>${text}</em></div>
    `;

    messages.appendChild(div);
    messages.scrollTop = messages.scrollHeight;
}

function abrirModalRefresh() {
    const overlay = document.getElementById("refreshOverlay");
    const modal = document.getElementById("refreshModal");

    overlay.style.display = "block";
    modal.style.display = "block";

    setTimeout(() => {
        overlay.classList.add("show");
        modal.classList.add("show");
    }, 10);
}

function fecharModalRefresh() {
    const overlay = document.getElementById("refreshOverlay");
    const modal = document.getElementById("refreshModal");

    overlay.classList.remove("show");
    modal.classList.remove("show");

    setTimeout(() => {
        overlay.style.display = "none";
        modal.style.display = "none";
    }, 200);
}

function confirmarRefresh() {
    window.location.reload();
}

document.addEventListener("keydown", function (e) {
    if (e.key === "F5") {
        e.preventDefault();
        abrirModalRefresh();
        return;
    }

    if (e.ctrlKey && e.key.toLowerCase() === "r") {
        e.preventDefault();
        abrirModalRefresh();
        return;
    }

    if (e.ctrlKey && e.shiftKey && e.key.toLowerCase() === "r") {
        e.preventDefault();
        abrirModalRefresh();
        return;
    }

    if (e.ctrlKey && e.key === "F5") {
        e.preventDefault();
        abrirModalRefresh();
        return;
    }
});

function copiarLink() {
    const url = window.location.href;
    const btn = document.getElementById("copyBtn");

    if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(url).then(() => {
            animarBotaoCopiado(btn);
        }).catch(() => {
            fallbackCopy(url, btn);
        });
    } 
    else {
        fallbackCopy(url, btn);
    }
}

function fallbackCopy(text, btn) {
    const input = document.createElement("input");
    input.value = text;
    document.body.appendChild(input);
    input.select();
    document.execCommand("copy");
    document.body.removeChild(input);

    animarBotaoCopiado(btn);
}

function animarBotaoCopiado(btn) {
    btn.classList.add("copied");
    btn.innerText = "Link Copiado!";

    setTimeout(() => {
        btn.classList.remove("copied");
        btn.innerText = "Copiar Link";
    }, 2000);
}




function sairSala(){
    const overlay = document.getElementById("refreshOverlay");
    const modal = document.getElementById("closeModal");

    overlay.style.display = "block";
    modal.style.display = "block";

    setTimeout(() => {
        overlay.classList.add("show");
        modal.classList.add("show");
    }, 10);
}

function fecharModalClose() {
    const overlay = document.getElementById("refreshOverlay");
    const modal = document.getElementById("closeModal");

    overlay.classList.remove("show");
    modal.classList.remove("show");

    setTimeout(() => {
        overlay.style.display = "none";
        modal.style.display = "none";
    }, 200);
}

function confirmarClose() {
    if (ws && ws.readyState === WebSocket.OPEN) {
        ws.close();
    }
    window.location.href = "/ChatAnonymous.Frontend";
}

function proximoRandom(){
    const overlay = document.getElementById("refreshOverlay");
    const modal = document.getElementById("nextRandomModal");

    overlay.style.display = "block";
    modal.style.display = "block";

    setTimeout(() => {
        overlay.classList.add("show");
        modal.classList.add("show");
    }, 10);
}

function fecharModalRandom() {
    const overlay = document.getElementById("refreshOverlay");
    const modal = document.getElementById("nextRandomModal");

    overlay.classList.remove("show");
    modal.classList.remove("show");

    setTimeout(() => {
        overlay.style.display = "none";
        modal.style.display = "none";
    }, 200);
}

function confirmarRandom(){
    if (ws && ws.readyState === WebSocket.OPEN) {
        ws.close();
    }

    window.location.href = "/ChatAnonymous.Frontend/pages/chat/html/chat.html?random=true";
}


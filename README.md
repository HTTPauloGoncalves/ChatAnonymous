🟣 ChatAnonymous — Chat Anônimo, Seguro e Minimalista

Bem-vindo ao ChatAnonymous, um chat em tempo real, simples, rápido e verdadeiramente anônimo.
Sem contas. Sem histórico. Sem rastreamento. Sem armazenamento de mensagens.
Cada sala existe apenas enquanto os usuários estão nela.

Construído em Go + WebSockets, com foco em:

Segurança

Anonimato real

Experiência fluida

Baixo consumo de recursos

Zero rastros

🚀 Funcionalidades Atuais
🟢 Conexão segura via WebSockets

Comunicação em tempo real entre todos os usuários da sala.

🟣 Salas temporárias protegidas por senha

Uma sala só existe se você tiver o link + password gerado na criação.

🔒 Sem armazenamento de mensagens

Nada é salvo no servidor.
Se recarregar a página, tudo desaparece.

🧩 Codinome local (não enviado ao backend)

O nome do usuário é salvo somente no front-end via localStorage.

👤 Mensagens diferenciadas entre remetente e outros usuários

Bolhas estilizadas para tornar o chat mais agradável.

📩 Notificações internas do sistema

Ex.: “Um usuário saiu da sala”.

🟪 Estilo moderno

Interface elegante com transições suaves, modo janela e efeitos blur.

🛠️ Funcionalidades Avançadas — Implementadas recentemente
🟤 1. Modal de prevenção de refresh

Impede recarregar a página acidentalmente e mostra aviso de perda de dados.

🟤 2. Mensagens do sistema com estilo próprio

Ex.:

Alguém desconectou…

🟤 3. Quebra de linha real nas mensagens

Utilizando white-space: pre-wrap.

🟤 4. Arquitetura otimizada com Client, Room e Hub

Separação profissional entre leitura, escrita e distribuição de mensagens.

🔮 Próximas Funcionalidades (Roadmap)

(já incluídas aqui como forma oficial do projeto)

🛡️ Segurança e Anti-Abuso

Rate limit anti-spam (limitar X mensagens por segundo por usuário)

Proteção contra mensagens mal formatadas / XSS

Sanitização automática de JSON

UUID interno por cliente (sem expor ao front)

💬 Experiência de Chat

“Usuário está digitando...” em tempo real

Scroll inteligente (não empurra mensagens se usuário está lendo acima)

Timer de inatividade com aviso para encerramento da sala

Indicador de latência (ping/pong WebSocket)

🧪 Usabilidade

Suporte a short links: https://anon.chat/r/abc123

Modal unificado para confirmação de ações

Tema claro/escuro (toggle no header)

🕵️ Modos de Uso

Modo totalmente anônimo: Usuário 1 / Usuário 2 / etc

Modo permanente: salas que não expiram automaticamente

Matchmaking aleatório (entre em uma fila e conecte com outro usuário)

📷 Extras opcionais

Envio de imagens pequenas (base64 transitório)

Emojis aprimorados via picker nativo

🧱 Arquitetura
Hub

Gerencia todas as salas existentes.

Room

lista de clientes

canais de join/leave/broadcast

timer de expiração

controle de mensagens

Client

conexão WebSocket individual

ReadPump() e WritePump() separados

buffer próprio para evitar travar broadcast

Frontend

HTML/CSS minimalista com efeito glass

WebSocket nativo

modal de codinome

modal anti-refresh

gerenciamento local de nome

renderização com animação de mensagens

📦 Como rodar
Backend (Go):
go mod tidy
go run ./cmd/server

Frontend:

A pasta /public contém os arquivos HTML, CSS e JS.
Basta abrir em localhost ou qualquer servidor estático simples.

🔐 Anonimato Real

Este projeto segue o conceito de Zero Knowledge:

Nenhuma mensagem é armazenada

Nome do usuário é apenas no front-end

Servidor não registra logs de conteúdo

Apenas gerencia a conexão WebSocket

Salas expiram automaticamente

❤️ Por que este projeto existe?

Para oferecer uma alternativa realmente segura, direta e sem rastros —
diferente de mensageiros tradicionais como WhatsApp e Telegram, que ainda dependem de servidores que retêm metadados.

Este projeto preza por:

✔ Liberdade
✔ Privacidade
✔ Simplicidade
✔ Zero rastreamento

⭐ Contribuições

Sinta-se livre para:

abrir issues

sugerir novas funções

reportar bugs

enviar PRs
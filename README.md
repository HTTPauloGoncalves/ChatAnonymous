# 🟣 ChatAnonymous - Chat Anônimo, Seguro e Minimalista

O **ChatAnonymous** é um chat em tempo real que prioriza **privacidade,
simplicidade e segurança**.
Sem contas. Sem histórico. Sem rastreamento.
Nenhuma mensagem é armazenada. Cada sala existe apenas enquanto houver
usuários conectados.

Construído em **Go + WebSockets**, o projeto foi desenhado para
entregar:

-   🔐 **Anonimato real**
-   ⚡ **Baixa latência e alta performance**
-   🧪 **UX suave e moderna**
-   🧩 **Arquitetura limpa e escalável**
-   🕳 **Zero rastros**

------------------------------------------------------------------------

## 🚀 Funcionalidades Atuais

### 🟢 Conexão WebSocket em tempo real

Mensagens instantâneas entre todos os usuários conectados à sala.

### 🟣 Salas temporárias protegidas por senha

Acesso somente com **ID + password** gerados automaticamente.

### 🔒 Nenhum armazenamento de mensagens

O servidor não grava **nada**.
Ao recarregar a página, o conteúdo desaparece para sempre.

### 🧩 Codinome local

O nome do usuário é salvo *apenas no front* (via `localStorage`)
garantindo anonimato.

### 👤 UI responsiva com mensagens estilizadas

Diferencia automaticamente o remetente dos demais usuários.

### 📩 Eventos do sistema

Exemplo:
 *Um usuário saiu da sala*

### 🪟 Design moderno

Glassmorphism, animações suaves e interface minimalista.

------------------------------------------------------------------------

## 🛠️ Funcionalidades Avançadas (já implementadas)

### 🟤 Modal anti-refresh

Evita recarregamento acidental que causaria perda total da sessão.

### 🟤 Mensagens do sistema com estilo próprio

Visual diferenciado e discreto.

### 🟤 Quebra de linha real nas mensagens

Renderização correta de textos longos e multilinhas.

### 🟤 Arquitetura com Client, Room e Hub

Separação clara entre leitura, escrita, distribuição e gerenciamento.

------------------------------------------------------------------------

# 🔮 Roadmap - Próximas Funcionalidades

## 🛡️ Segurança & Anti-Abuso

-   Rate-limit anti-spam
-   Proteção contra XSS e mensagens malformadas
-   Sanitização automática no backend
-   UUID interno para clientes (sem expor ao front)

## 💬 Experiência do Chat

-   Indicador "Usuário está digitando..."
-   Scroll inteligente (somente desce se estiver no final)
-   Aviso de inatividade com contagem regressiva
-   Indicador de latência (ping/pong)

## 🧪 Usabilidade

-   Short links do tipo `anon.chat/r/abc123`
-   Modal unificado para confirmações
-   Tema claro/escuro

## 🕵️ Modos de Uso

-   Modo totalmente anônimo (Usuário 1, Usuário 2, etc.)
-   Salas permanentes opcionais
-   Matchmaking aleatório (modo Omegle seguro)

## 📷 Extras opcionais

-   Envio de imagens pequenas (base64 transitório, não persistido)
-   Emoji picker

------------------------------------------------------------------------

# 🧱 Arquitetura do Projeto

### **Hub**

Gerencia o conjunto de salas vivas no servidor.

### **Room**

-   Lista de clientes
-   Canais: `Join`, `Leave`, `Broadcast`
-   Timer de expiração
-   Encerramento seguro

### **Client**

-   Conexão individual WebSocket
-   `ReadPump` e `WritePump` isolados
-   Buffer próprio para evitar travamentos no broadcast

### **Frontend**

-   HTML/CSS com Glass Effect
-   WebSocket nativo
-   Modal de codinome
-   Modal anti-refresh
-   Animações + responsividade
-   Renderização das mensagens com bolhas estilizadas

------------------------------------------------------------------------

# 📦 Como rodar

## Backend (Go):

``` bash
go mod tidy
go run ./cmd/server
```

## Frontend:

Arquivos na pasta `/public`.\
Basta abrir um servidor estático simples ou usar `Live Server` no
VSCode.

------------------------------------------------------------------------

# 🔐 Filosofia de Anonimato

O ChatAnonymous segue a ideia de **Zero Knowledge**:

-   Nenhuma mensagem é salva
-   Nome do usuário nunca vai para o backend
-   Nenhum log de conteúdo é registrado
-   Somente metadados mínimos para manter a sala funcionando
-   Salas expiram automaticamente

------------------------------------------------------------------------

# ❤️ Por que esse projeto existe?

Criado como uma alternativa segura e direta a mensageiros tradicionais,
que apesar de criptografados, **mantêm metadados, números de telefone e
histórico de conexões**.

O ChatAnonymous é para quem quer:

-   ✔ Privacidade extrema
-   ✔ Comunicação efêmera
-   ✔ Zero dependência de empresas
-   ✔ Uma ferramenta simples e útil

------------------------------------------------------------------------

# ⭐ Contribuições

Contribua com:

-   Issues
-   Sugestões
-   Correções
-   Pull Requests

------------------------------------------------------------------------

# Context7 Documentation: Go Google Cloud Function Template

**Source:** Este documento é um resumo estruturado do ficheiro `ai_instructions.md` para uso por sistemas de processamento de contexto como o Context7.

## 1. Overview

*   **Purpose:** Template para gerar Google Cloud Functions (GCF) completas e funcionais escritas em Go.
*   **Core Functionality:** Fornece uma base com autenticação JWT, conexão a banco de dados SQL (MySQL/PostgreSQL), configuração via variáveis de ambiente e estrutura de projeto organizada.
*   **Generation Trigger:** A geração de uma função específica é iniciada a partir de um Nome de Função, Descrição e Schema de Banco de Dados fornecidos pelo utilizador, utilizando as diretrizes em `ai_instructions.md`.

## 2. Core Technologies & Libraries

*   **Language:** Go (versão 1.21+ recomendada)
*   **Cloud Platform:** Google Cloud Functions (GCF)
*   **Framework:** `github.com/GoogleCloudPlatform/functions-framework-go`
*   **Database:** SQL (com driver específico, ex: `github.com/go-sql-driver/mysql` ou `github.com/lib/pq`)
*   **Authentication:** JWT (`github.com/golang-jwt/jwt/v5`)
*   **Configuration:** Variáveis de Ambiente (leitura com `os.Getenv`, opcionalmente `github.com/joho/godotenv` para `.env` local)

## 3. Expected Project Structure (Generated)

```
go_lang_gcloud_function_template/
├── internal/           # Lógica interna da aplicação
│   ├── handler/        # Manipuladores de requests HTTP
│   ├── model/          # Structs de dados (mapeamento DB/JSON)
│   └── store/          # Lógica de acesso ao banco de dados (CRUD)
├── pkg/                # Código partilhável (ex: conexão DB)
│   └── database/
│       └── connection.go # Lógica de conexão DB (se refatorada)
├── .env.example        # Ficheiro de exemplo para variáveis de ambiente locais
├── env.yaml            # Ficheiro (NÃO COMMITADO) para variáveis de ambiente no deploy
├── go.mod              # Definição do módulo Go e dependências
├── go.sum              # Checksums de dependências
├── main.go             # Ponto de entrada principal (HTTP handler, init, routing)
└── README.md           # Documentação para utilizador final (setup, deploy, endpoints)
```

## 4. Key Features & Implementation Details

*   **JWT Authentication:**
    *   Implementada na função `validateJWT` dentro de `main.go`.
    *   Verifica o header `Authorization: Bearer <token>` em cada request.
    *   A chave secreta (`jwtSecret`) é lida da variável de ambiente `JWT_SECRET_KEY` na inicialização (`init` func).
*   **Database Connection:**
    *   Implementada na função `connectToDB` (em `main.go` ou `pkg/database/connection.go`).
    *   Lê credenciais (`DB_USER`, `DB_PASSWORD`, `DB_HOST`, `DB_PORT`, `DB_NAME`) de variáveis de ambiente.
    *   Utiliza o pacote `database/sql` e o driver SQL apropriado.
    *   Idealmente, estabelece um pool de conexões reutilizável (configurado na `init` func).
*   **Configuration:**
    *   Primariamente via Variáveis de Ambiente.
    *   `.env.example`: Modelo para configuração local (usado com `godotenv`).
    *   `env.yaml`: Ficheiro usado com `gcloud functions deploy --env-vars-file` para definir variáveis no ambiente GCF.
*   **HTTP Handling & Routing:**
    *   A função `HandleFunction` em `main.go` é o ponto de entrada principal registrado no `functions-framework`.
    *   Realiza a validação JWT inicial.
    *   Contém um `switch` (ou lógica similar) para rotear requests baseado no método HTTP (`r.Method`) e path (`r.URL.Path`) para handlers específicos (localizados em `internal/handler/`).
*   **Code Organization:**
    *   `internal/`: Código específico da aplicação (não importável por outros projetos).
        *   `model`: Structs Go.
        *   `store`: Funções de interação com o banco de dados.
        *   `handler`: Funções que lidam com a lógica HTTP (decodificação, validação, chamadas ao store, resposta).
    *   `pkg/`: Código reutilizável (opcional, ex: pacote `database`).

## 5. Generation Workflow (Based on `ai_instructions.md`)

1.  **Input:** AI receives Function Name, Description, DB Schema.
2.  **Base Project Creation:**
    *   Create directory structure.
    *   Run `go mod init`.
    *   Run `go get` for core dependencies (framework, db driver, jwt, godotenv).
    *   Create `.env.example`.
    *   Create base `main.go` (with `init`, `main`, `HandleFunction`, `connectToDB`, `validateJWT`, env var loading, initial routing).
    *   Create base `README.md`.
3.  **Specific Function Implementation:**
    *   Create model struct in `internal/model/` based on Schema.
    *   Create store functions in `internal/store/` for DB operations.
    *   Create handler function in `internal/handler/`.
    *   Update router in `main.go` to call the new handler.
    *   Run `go mod tidy`.
    *   Update `README.md` with the new endpoint documentation and cURL example.

## 6. Usage & Deployment (Generated Function)

*   **Local Testing:** Use `funcframework --target=HandleFunction --source=.` (requer `go install .../funcframework`). Assume `.env` existe para variáveis locais.
*   **Deployment:** Use `gcloud functions deploy ... --env-vars-file env.yaml ...` (conforme detalhado no `README.md` e `ai_instructions.md`).

## 7. Example Endpoint (Generated: `POST /user-profiles`)

*   **Purpose:** Creates a new user profile.
*   **Method:** `POST`
*   **Path:** `/user-profiles`
*   **Auth:** Requires `Authorization: Bearer <token>` header.
*   **Request Body (JSON):** `model.UserProfile` struct (e.g., `{"user_id": 123, "full_name": "...", "email": "...", "bio": "..."}`)
*   **Success Response (201):** `{"message": "...", "id": <new_id>}`
*   **cURL Example (Postman format):**
    ```bash
    curl --location --request POST '{{server}}/user-profiles' \
    --header 'Authorization: Bearer {{token}}' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "user_id": 123,
        "full_name": "Nome Completo Postman",
        "email": "postman@exemplo.com",
        "bio": "Criado via cURL/Postman."
    }'
    ``` 

## 8. Example `go.mod` (Generated):

```mod
module go_lang_gcloud_function_template

go 1.21 // Ou versão compatível

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.12.1 // Versão exemplo
	github.com/go-sql-driver/mysql v1.8.1             // Versão exemplo
	github.com/golang-jwt/jwt/v5 v5.2.1                // Versão exemplo
	github.com/joho/godotenv v1.5.1                   // Versão exemplo
)

// indiretas podem aparecer aqui após 'go mod tidy'
```

## 9. Example `env.example` (Generated):

```dotenv
# .env.example
DB_USER="root"
DB_PASSWORD=""
DB_HOST="localhost"
DB_PORT="3306"
DB_NAME="user_db" # Exemplo
JWT_SECRET_KEY="change_this_secret_key"
```

## 10. Example `env.yaml` (Generated):

```yaml
# env.yaml (NÃO COMMITAR)
DB_USER: 'usuario_db'
DB_PASSWORD: 'senha_super_secreta'
DB_HOST: 'ip_ou_cloud_sql_socket'
DB_PORT: '3306'
DB_NAME: 'nome_db'
JWT_SECRET_KEY: 'sua_chave_jwt_muito_segura_aqui'
```

## 11. Example `main.go` (Base Generated Code):

```go
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"go_lang_gcloud_function_template/internal/handler" // Assumindo que handlers estão em internal
	"go_lang_gcloud_function_template/internal/store"   // Assumindo que store está em internal

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/go-sql-driver/mysql" // Import implícito para registro do driver
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"          // Para dev local
)

var (
	db        *sql.DB
	jwtSecret []byte
)

// init carrega configuração e inicializa conexão DB reutilizável (se aplicável)
func init() {
	// Carregar .env apenas para desenvolvimento local, ignora erro se não existir
	_ = godotenv.Load() 

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Carregar segredo JWT do ambiente
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		log.Fatal("Variável de ambiente JWT_SECRET_KEY não definida")
	}
	jwtSecret = []byte(secret)

	// Configurar e conectar ao banco de dados
	var err error
	db, err = connectToDB()
	if err != nil {
		log.Fatalf("Erro fatal ao conectar ao banco de dados: %v", err)
	}
	// Opcional: Configurar pool de conexões
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Registrar a função HTTP principal
	funcframework.RegisterHTTPFunction("/", HandleFunction)
}

func main() {
	// O framework cuida do início do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão para GCF HTTP
	}
	log.Printf("Servidor iniciado na porta %s", port)
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("Erro ao iniciar funcframework: %v", err)
	}
}

// HandleFunction é o ponto de entrada principal
func HandleFunction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1. Validar JWT
	if !validateJWT(r) {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// 2. Roteamento (exemplo simples)
	w.Header().Set("Content-Type", "application/json") // Definir Content-Type padrão

	switch {
	// Rota para User Profiles (Adicionada pelo workflow de geração)
	case strings.HasPrefix(r.URL.Path, "/user-profiles"):
		if r.Method == http.MethodPost {
			handler.CreateUserProfileHandler(w, r, db) // Chamar handler específico
		} else {
			http.Error(w, `{"error": "Method Not Allowed"}`, http.StatusMethodNotAllowed)
		}
		
	// Adicionar outras rotas aqui (ex: /products, /orders)

	default:
		http.Error(w, `{"error": "Not Found"}`, http.StatusNotFound)
	}
}

// connectToDB lê variáveis de ambiente e conecta ao DB
func connectToDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return nil, fmt.Errorf("variáveis de ambiente do DB (DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME) não estão completamente definidas")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	
	log.Printf("Conectando ao DB: %s@%s:%s/%s", dbUser, dbHost, dbPort, dbName)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com DB: %w", err)
	}

	// Verificar a conexão
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = conn.PingContext(ctx)
	if err != nil {
		conn.Close() // Fechar a conexão se o ping falhar
		return nil, fmt.Errorf("erro ao fazer ping no DB: %w", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso.")
	return conn, nil
}

// validateJWT extrai e valida o token JWT do header Authorization
func validateJWT(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("Cabeçalho Authorization ausente")
		return false // Sem header, não autorizado
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		log.Println("Formato inválido do cabeçalho Authorization")
		return false // Formato inválido
	}
	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validar o algoritmo de assinatura esperado (ex: HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("algoritmo de assinatura inesperado: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		log.Printf("Erro ao parsear/validar token: %v", err)
		return false // Erro na validação
	}

	if !token.Valid {
		 log.Println("Token JWT inválido")
		return false // Token inválido
	}

	return true // Token válido
}

## 12. Example `internal/model/user_profile.go` (Generated for CreateUserProfile):

```go
package model

import "time"

// UserProfile representa a estrutura da tabela user_profiles
type UserProfile struct {
	ID        int64     `json:"id,omitempty"` // omitempty para criação
	UserID    int64     `json:"user_id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Bio       *string   `json:"bio,omitempty"` // Ponteiro para permitir NULL/omitempty
	CreatedAt time.Time `json:"created_at,omitempty"`
}
```

## 13. Example `internal/store/user_profile_store.go` (Generated for CreateUserProfile):

```go
package store

import (
	"database/sql"
	"fmt"
	"log"
	"go_lang_gcloud_function_template/internal/model"
)

// CreateUserProfile insere um novo perfil de utilizador na base de dados.
func CreateUserProfile(db *sql.DB, profile model.UserProfile) (int64, error) {
	if db == nil {
		 return 0, fmt.Errorf("conexão DB é nil")
	}

	query := `INSERT INTO user_profiles (user_id, full_name, email, bio) VALUES (?, ?, ?, ?)`
	
	result, err := db.Exec(query, profile.UserID, profile.FullName, profile.Email, profile.Bio)
	if err != nil {
		// Idealmente, verificar erros específicos (ex: entrada duplicada)
		log.Printf("Erro ao inserir user profile para user_id %d: %v", profile.UserID, err)
		return 0, fmt.Errorf("falha ao criar perfil: %w", err) // Encadeamento de erro
	}

	id, err := result.LastInsertId()
	if err != nil {
		 log.Printf("Erro ao obter LastInsertId para user_id %d: %v", profile.UserID, err)
		return 0, fmt.Errorf("perfil criado, mas falha ao obter ID: %w", err)
	}
	
	log.Printf("User profile criado com sucesso com ID: %d para UserID: %d", id, profile.UserID)
	return id, nil
}
```

## 14. Example `internal/handler/user_profile_handler.go` (Generated for CreateUserProfile):

```go
package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"go_lang_gcloud_function_template/internal/model"
	"go_lang_gcloud_function_template/internal/store"
)

// CreateUserProfileHandler processa requisições POST para /user-profiles
func CreateUserProfileHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		 http.Error(w, `{"error": "Method Not Allowed"}`, http.StatusMethodNotAllowed)
		 return
	}

	var profile model.UserProfile

	// Decodificar JSON
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		log.Printf("Erro ao decodificar corpo da requisição para criar perfil: %v", err)
		http.Error(w, `{"error": "Corpo da requisição inválido"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validação básica de entrada
	if profile.UserID == 0 || profile.FullName == "" || profile.Email == "" {
		log.Printf("Campos obrigatórios em falta para criar perfil: %+v", profile)
		 http.Error(w, `{"error": "Campos obrigatórios em falta: user_id, full_name, email"}`, http.StatusBadRequest)
		 return
	}

	// Chamar store para criar
	newID, err := store.CreateUserProfile(db, profile)
	if err != nil {
		log.Printf("Erro vindo do store ao criar perfil para user_id %d", profile.UserID)
		http.Error(w, fmt.Sprintf(`{"error": "Falha ao criar perfil de utilizador: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	// Retornar sucesso
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Perfil de utilizador criado com sucesso",
		"id":      newID, // Pode ser 0 se LastInsertId falhar
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		 log.Printf("Erro ao encodificar resposta de sucesso para criar perfil: %v", err)
	}
}
```

## 15. Example `README.md` (Generated):

```markdown
# Google Cloud Functions Template com GoLang

Template para criar funções no Google Cloud Functions usando GoLang, com JWT e MySQL.

## Funcionalidades Base

*   Autenticação JWT via header `Authorization: Bearer <token>`.
*   Conexão com banco de dados MySQL (configurável via env vars).
*   Estrutura básica de projeto com `internal` e `pkg` (opcional).
*   Logging configurado.

## Configuração

1.  **Variáveis de Ambiente:** Copie `.env.example` para `.env` e preencha com suas credenciais de banco de dados e uma chave secreta JWT segura.
    ```bash
    cp .env.example .env 
    # Edite o ficheiro .env com os seus valores
    ```
    **Importante:** NÃO comitar o ficheiro `.env` no Git. Use variáveis de ambiente diretamente no ambiente de deploy (Cloud Functions, etc.).

2.  **Dependências:** Instale as dependências Go:
    ```bash
    go mod tidy
    ```

3.  **Banco de Dados:** Certifique-se de que o schema necessário existe no seu banco de dados MySQL. Veja o schema de exemplo em `ai_instructions.md`.

## Variáveis de Ambiente Necessárias

*   `DB_USER`: Utilizador do banco de dados.
*   `DB_PASSWORD`: Senha do banco de dados.
*   `DB_HOST`: Host do banco de dados.
*   `DB_PORT`: Porta do banco de dados (ex: 3306).
*   `DB_NAME`: Nome do banco de dados.
*   `JWT_SECRET_KEY`: Chave secreta para assinar e validar tokens JWT.

## Teste Local

Use o framework de funções do Google Cloud (requer instalação: `go install github.com/GoogleCloudPlatform/functions-framework-go/funcframework`):

```bash
# Carrega .env automaticamente se existir e inicia o servidor
funcframework --target=HandleFunction --source=. 
```
O servidor iniciará na porta 8080 por padrão (ou na definida por `PORT`).

## Endpoints

*   **POST /user-profiles**: Cria um novo perfil de utilizador.
    *   **Header Obrigatório:** `Authorization: Bearer <seu_token_jwt_valido>`
    *   **Corpo (JSON):**
        ```json
        {
          "user_id": 123,
          "full_name": "Nome Completo",
          "email": "email@exemplo.com",
          "bio": "Uma breve biografia." 
        }
        ```
    *   **Resposta Sucesso (201 Created):**
        ```json
        {
          "message": "Perfil de utilizador criado com sucesso",
          "id": 42 
        }
        ```
    *   **Respostas Erro:** 400 (Bad Request), 401 (Unauthorized), 500 (Internal Server Error).
    *   **Exemplo cURL (Postman):**
        ```bash
        curl --location --request POST '{{server}}/user-profiles' \
        --header 'Authorization: Bearer {{token}}' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "user_id": 123,
            "full_name": "Nome Completo Postman",
            "email": "postman@exemplo.com",
            "bio": "Criado via cURL/Postman."
        }'
        ```
        *Substitua `{{server}}` pela URL da sua função (local ou cloud) e `{{token}}` pelo seu token JWT válido.*

    *   _(Adicionar outros endpoints aqui conforme são implementados)_

## Deploy no Google Cloud Functions

1.  **Crie um ficheiro `env.yaml`** (NÃO o adicione ao Git!) com as suas variáveis de ambiente:
    ```yaml
    # env.yaml
    DB_USER: 'usuario_db'
    DB_PASSWORD: 'senha_super_secreta'
    DB_HOST: 'ip_ou_cloud_sql_socket'
    DB_PORT: '3306'
    DB_NAME: 'nome_db'
    JWT_SECRET_KEY: 'sua_chave_jwt_muito_segura_aqui'
    ```
    **Nota:** Para Cloud SQL, `DB_HOST` geralmente é o caminho do socket (ex: `/cloudsql/seu-projeto:sua-regiao:sua-instancia`).

2.  **Execute o comando deploy:**
    ```bash
    gcloud functions deploy <NOME_DA_SUA_FUNCAO> \
        --runtime go121 \
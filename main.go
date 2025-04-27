package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"cloud.google.com/go/functions/metadata"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

// Configuração do banco de dados MySQL
const (
	dbUser     = "root"
	dbPassword = ""
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "animals_db"
)

// Chave secreta para JWT (em produção, use uma variável de ambiente ou secret manager)
var jwtSecret = []byte("my_secret_key")

func init() {
	// Configuração de logs
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Função principal do Google Cloud Functions
func HandleFunction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Verificar autenticação JWT
	if !validateJWT(r) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
		return
	}

	// Conectar ao banco de dados
	db, err := connectToDB()
	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
		return
	}
	defer db.Close()

	// Roteamento básico
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/animals" {
			listAllAnimals(w, r, db)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Not Found"})
}

// Função para validar JWT
func validateJWT(r *http.Request) bool {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return false
	}

	// Remover 'Bearer ' do token, se presente
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		log.Printf("Erro ao validar token: %v", err)
		return false
	}

	return true
}

// Função para conectar ao banco de dados MySQL
func connectToDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Função de exemplo para listar todos os animais com paginação
func listAllAnimals(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Implementação de paginação
	page := 1
	limit := 10
	if p := r.URL.Query().Get("page"); p != "" {
		if n, err := strconv.Atoi(p); err == nil && n > 0 {
			page = n
		}
	}
	if l := r.URL.Query().Get("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil && n > 0 {
			limit = n
		}
	}

	offset := (page - 1) * limit

	query := "SELECT id, name, species FROM animals LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Erro ao consultar animais: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	animals := []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Species string `json:"species"`
	}{}

	for rows.Next() {
		var animal struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Species string `json:"species"`
		}
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Species); err != nil {
			log.Printf("Erro ao escanear animal: %v", err)
			continue
		}
		animals = append(animals, animal)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animals)
}

// Função para registrar a função no Google Cloud Functions
func init() {
	functions.HTTP("HandleFunction", HandleFunction)
} 
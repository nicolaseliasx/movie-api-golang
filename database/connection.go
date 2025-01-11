package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("A variável de ambiente MONGO_URI não está definida")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	// Criar um contexto com timeout para a conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Garantir que o contexto seja cancelado após a execução

	// Conectar ao MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	// Verificar a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Erro ao verificar a conexão com o MongoDB: %v", err)
	}

	// Selecionar o banco de dados (substitua "movieapi" pelo nome do seu banco)
	DB = client.Database("movieapi")

	log.Println("Conectado ao MongoDB com sucesso!")
}

package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Name  string
	Email string
}

func ConnectNosql() {
	// Créer une instance de contexte
	ctx := context.Background()

	// Configuration de l'URI de la base de données MongoDB
	uri := "mongodb://root:example@mongo:27017/"

	// Configuration des options de connexion
	clientOptions := options.Client().ApplyURI(uri)

	// Connexion à la base de données
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Vérification de la connexion
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Récupérer la collection "users"
	usersCollection := client.Database("bdGolang").Collection("api-golang")

	// Créer un utilisateur
	user := User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// Insérer le document dans la collection "users"
	result, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	// Afficher l'ID du document inséré
	fmt.Println("Document inséré avec l'ID :", result.InsertedID)

	// Déconnexion de la base de données
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

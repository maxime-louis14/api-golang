# API Golang pour Scrapeur Golang

Ce projet contient une API en Go utilisée dans le cadre du projet Scrapeur Golang, principalement pour la gestion des recettes.

## Structure du projet

- **cmd/** : Fichiers principaux de l'application.
- **database/** : Scripts et configurations de la base de données.
- **models/** : Définitions des modèles de données.
- **routes/** : Définition des routes et des contrôleurs.
- **.gitattributes** : Configuration Git.
- **Dockerfile** : Configuration pour Docker.
- **README.md** : Ce fichier.
- **data.json** : Stockage NoSQL.
- **docker-compose.yml** : Configuration pour Docker Compose.
- **go.mod** et **go.sum** : Gestion des dépendances Go.

## Installation

1. Assurez-vous d'avoir Go installé.
2. Clonez le repository :

   ```bash
   git clone <url-du-repository>
   ```
## Configuration
- Configurez l'application en ajustant les fichiers de configuration appropriés.

Lancement
Pour démarrer l'API :
```bash
go run main.go
```
Ou avec Docker Compose :
```bash
docker-compose up
```

## Contribution
Les contributions sont les bienvenues ! Suivez ces étapes pour contribuer :

- Fork du repository
- Créez une branche pour votre fonctionnalité (git checkout -b feature/NomDeLaFonctionnalité)
- Commit de vos changements (git commit -am 'Ajout d'une nouvelle fonctionnalité')
- Push vers la branche (git push origin feature/NomDeLaFonctionnalité)
- Créez un nouveau Pull Request
- Assurez-vous de tester votre code avant de soumettre le Pull Request.

## Licence
Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.

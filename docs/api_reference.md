# Documentation de l'API QAlpuch

Ce document fournit une référence détaillée pour l'API REST QAlpuch.

**URL de base :** `/api/v1`

### Paramètres d'URL

- `:id` : Représente un identifiant numérique (Integer).
- `:cuid` : Représente un identifiant unique de type CUID (String).

## Authentification

La plupart des endpoints nécessitent une authentification via un **JSON Web Token (JWT)**. Pour obtenir un jeton, vous
devez utiliser les endpoints `/login` ou `/register`.

Les requêtes authentifiées doivent inclure un en-tête `Authorization`:

`Authorization: Bearer <votre_jeton_jwt>`

Certains endpoints nécessitent un rôle utilisateur spécifique (`admin` ou `user`). Cela sera indiqué dans la description
de l'endpoint.

## Réponses d'Erreur

L'API utilise les codes de statut HTTP standard pour indiquer le succès ou l'échec d'une requête. En cas d'erreur, le
corps de la réponse sera un objet JSON contenant un message explicatif.

**Format d'Erreur Général :**

```json
{
  "success": false,
  "message": "Description de l'erreur."
}
```

- **400 Bad Request :** La requête était malformée (ex: JSON invalide, paramètres manquants).
- **401 Unauthorized :** L'authentification a échoué ou est requise.
- **403 Forbidden :** L'utilisateur authentifié n'a pas les droits nécessaires pour accéder à la ressource.
- **404 Not Found :** La ressource demandée n'a pas été trouvée.
- **500 Internal Server Error :** Une erreur inattendue est survenue côté serveur.

---

## Auth

Endpoints pour l'authentification et la gestion des comptes utilisateurs.

### Connexion

`POST /login`

Authentifie un utilisateur et retourne un JWT.

**Corps de la requête :** (`application/json`)

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Réponse de succès (200 OK) :**

```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Inscription

`POST /register`

Crée un nouveau compte utilisateur.

**Corps de la requête :** (`application/json`)

```json
{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "password123"
}
```

**Réponse de succès (201 Created) :**

```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "id": 1,
    "name": "newuser",
    "email": "newuser@example.com",
    "role": "user",
    "created_at": "...",
    "updated_at": "..."
  }
}
```

### Déconnexion

`GET /logout`

Déconnecte l'utilisateur actuel. (L'implémentation invalide probablement le jeton si une liste noire est utilisée).

**Authentification :** JWT Utilisateur requis.

### Changer le mot de passe

`POST /change-password`

Permet à un utilisateur authentifié de changer son mot de passe.

**Authentification :** JWT Utilisateur requis.

**Corps de la requête :** (`application/json`)

```json
{
  "oldPassword": "password123",
  "newPassword": "newStrongPassword456"
}
```

---

## Users (Utilisateurs)

Endpoints pour la gestion des comptes utilisateurs.

### Obtenir tous les utilisateurs

`GET /users`

Récupère une liste de tous les utilisateurs.

**Authentification :** Rôle Administrateur requis.

**Réponse de succès (200 OK) :**

```json
{
    "success": true,
    "message": "Users retrieved successfully",
    "data": [
        {
            "id": 1,
            "name": "admin",
            "email": "admin@example.com",
            "role": "admin"
        },
        {
            "id": 2,
            "name": "user1",
            "email": "user1@example.com",
            "role": "user"
        }
    ]
}
```

### Obtenir un utilisateur par ID

`GET /users/:id`

Récupère un seul utilisateur par son ID.

**Authentification :** Rôle Administrateur requis.

### Mettre à jour un utilisateur

`PUT /users/:id`

Met à jour les informations d'un utilisateur. Un administrateur peut mettre à jour n'importe quel utilisateur. Un
utilisateur normal ne peut mettre à jour que ses propres informations.

**Authentification :** JWT Utilisateur requis.

**Corps de la requête :** (`application/json`)

```json
{
  "name": "Nom Mis à Jour",
  "email": "updated@example.com",
  "role": "user"
}
```

### Supprimer un utilisateur

`DELETE /users/:id`

Supprime un utilisateur par son ID.

**Authentification :** Rôle Administrateur requis.

### Supprimer l'utilisateur actuel

`DELETE /users/me`

Permet à un utilisateur de supprimer son propre compte.

**Authentification :** JWT Utilisateur requis.

### Créer un utilisateur (Admin)

`POST /users`

Crée un nouvel utilisateur.

**Authentification :** Rôle Administrateur requis.

**Corps de la requête :** (`application/json`)

```json
{
  "name": "Nom Utilisateur",
  "email": "utilisateur@example.com",
  "role": "user"
}
```

### Rechercher des utilisateurs (Admin)

`GET /users/search`

Recherche des utilisateurs par nom ou email.

**Authentification :** Rôle Administrateur requis.

**Paramètres de la requête :**

- `q` (string, requis): Le terme de recherche.

---

## Files (Fichiers)

Endpoints pour la gestion des fichiers.

### Téléverser un fichier

`POST /files/upload`

Téléverse un fichier.

**Authentification :** JWT Utilisateur requis.

**Corps de la requête :** (`multipart/form-data`)

- **file**: Le fichier à téléverser.
- **name**: Le nom du fichier.

### Obtenir les métadonnées d'un fichier

`GET /files/:cuid`

Récupère les métadonnées d'un fichier spécifique.

**Authentification :** JWT Utilisateur requis.

**Réponse de succès (200 OK) :**

```json
{
    "success": true,
    "message": "File retrieved successfully",
    "data": {
        "id": "clvb2qabc000008l21234abcd",
        "filename": "mon-fichier.txt",
        "path": "/uploads/mon-fichier.txt",
        "size": 1024,
        "mimeType": "text/plain"
    }
}
```

### Télécharger un fichier

`GET /files/:cuid/download`

Télécharge le contenu binaire d'un fichier.

**Authentification :** JWT Utilisateur requis.

### Lister les fichiers d'un utilisateur

`GET /files`

Liste tous les fichiers appartenant à l'utilisateur authentifié.

**Authentification :** JWT Utilisateur requis.

### Supprimer un fichier

`DELETE /files/:cuid`

Supprime un fichier.

**Authentification :** JWT Utilisateur requis.

### Renommer un fichier

`PUT /files/:cuid`

Renomme un fichier.

**Authentification :** JWT Utilisateur requis.

**Corps de la requête :** (`application/json`)

```json
{
  "name": "nouveau-nom.txt"
}
```

---

## Tasks (Tâches)

Endpoints pour la création et la gestion des tâches de traitement.

### Obtenir les tâches d'un utilisateur

`GET /tasks`

Récupère une liste de toutes les tâches créées par l'utilisateur authentifié, y compris leurs journaux (logs).

**Authentification :** JWT Utilisateur requis.

**Réponse de succès (200 OK) :**

```json
{
    "success": true,
    "message": "Tasks retrieved successfully",
    "data": [
        {
            "id": "clvb2r9qj000108l23456defg",
            "config": "{}",
            "status": "completed",
            "source_file_id": "clvb2qabc000008l21234abcd",
            "result_file_id": "clvb2sxyz000208l25678efgh",
            "logs": [
                {
                    "id": "01H8Y3J0Z4QWERTY1234567890",
                    "task_status": "processing",
                    "message": "Task processing started",
                    "created_at": "2025-09-17T10:05:00Z",
                    "updated_at": "2025-09-17T10:05:00Z"
                }
            ],
            "created_at": "2025-09-17T10:00:00Z",
            "updated_at": "2025-09-17T10:06:00Z"
        }
    ]
}
```

### Créer une tâche

`POST /tasks`

Crée une nouvelle tâche à partir d'un fichier existant.

**Authentification :** JWT Utilisateur requis.

**Corps de la requête :** (`application/json`)

```json
{
  "file_id": "clv9p5n3h000008l9c4f2g8h1",
  "config": "{ \"param\": \"value\" }"
}
```

**Réponse de succès (201 Created) :**

```json
{
    "success": true,
    "message": "Task created successfully",
    "data": {
        "id": "clvb3abcd000108l2hij1234k",
        "config": "{ \"param\": \"value\" }",
        "status": "pending",
        "sourceFileId": "clv9p5n3h000008l9c4f2g8h1"
    }
}
```

### Supprimer une tâche

`DELETE /tasks/:cuid`

Supprime une tâche. Un utilisateur ne peut supprimer que les tâches qu'il a créées et qui ne sont pas en cours de
`traitement` ou `terminées`. Un administrateur peut supprimer n'importe quelle tâche.

**Authentification :** JWT Utilisateur requis.

---

## Workers

Endpoints pour la gestion des workers et de la file d'attente des tâches.

### Obtenir tous les Workers

`GET /worker`

Récupère la liste de tous les workers enregistrés.

**Authentification :** Rôle Administrateur requis.

### Obtenir un Worker par ID

`GET /worker/:cuid`

Récupère les informations d'un worker spécifique, y compris les tâches qu'il a traitées.

**Authentification :** Rôle Administrateur requis.

### Créer un Worker

`POST /worker`

Crée un nouveau worker et retourne un jeton d'enregistrement à usage unique.

**Authentification :** Rôle Administrateur requis.

**Corps de la requête :** (`application/json`)

```json
{
  "name": "Mon-Nouveau-Worker"
}
```

### Supprimer un Worker

`DELETE /worker/:cuid`

Supprime un worker par son ID.

**Authentification :** Rôle Administrateur requis.

### Enregistrer un Worker

`POST /worker/register`

Utilisé par un worker pour échanger son jeton d'enregistrement contre un JWT persistant.

**Corps de la requête :** (`application/json`)

```json
{
  "token": "a1b2c3d4-e5f6-7890-1234-567890abcdef"
}
```

### Obtenir une tâche en attente

`GET /tasks/pending`

Assigne une tâche en attente à un worker authentifié.

**Authentification :** JWT Worker requis.

### Mettre à jour le statut d'une tâche

`PATCH /tasks/:cuid`

Permet à worker de mettre à jour le statut d'une tâche qu'il traite.

**Authentification :** JWT Worker requis.

**Corps de la requête :** (`application/json`)

```json
{
  "status": "processing",
  "status_message": "Mise à jour depuis le worker..."
}
```

### Téléverser le résultat d'une tâche

`POST /tasks/:cuid/result`

Permet à worker de téléverser le résultat d'une tâche terminée.

**Authentification :** JWT Worker requis.

**Corps de la requête :** (`multipart/form-data`)

- **file**: Le fichier de résultat.


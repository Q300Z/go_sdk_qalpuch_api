# Guide d'Intégration des Workers

Ce guide explique comment connecter un worker de traitement à l'API QAlpuch. Le processus implique qu'un administrateur
crée d'abord un enregistrement de worker pour générer un jeton d'enregistrement, que le worker utilise ensuite pour
s'authentifier et recevoir un JWT persistant pour la communication avec l'API.

## Étape 1 : Création du Worker (Tâche Administrateur)

Un administrateur doit d'abord créer une entité de worker dans le système. Cette action génère un **Jeton
d'Enregistrement** à usage unique dont le worker aura besoin pour son authentification initiale.

**Endpoint :** `POST /api/v1/worker`
**Authentification :** JWT Administrateur requis.

#### Corps de la Requête

```json
{
  "name": "Mon-Nouveau-Worker-Video"
}
```

#### Réponse de Succès (201 Created)

L'API retourne les détails du nouveau worker, y compris le `token` crucial. **Ceci est le jeton d'enregistrement.
Copiez-le et fournissez-le de manière sécurisée à l'application du worker.**

```json
{
  "success": true,
  "message": "Worker created successfully",
  "data": {
    "id": "clv9p3b5a000008l9g2h7d4e2",
    "name": "Mon-Nouveau-Worker-Video",
    "token": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
    "created_at": "2025-09-16T20:15:00Z",
    "updated_at": "2025-09-16T20:15:00Z"
  }
}
```

---

## Étape 2 : Enregistrement et Authentification du Worker

L'application du worker utilise le `registrationToken` de l'étape 1 pour s'enregistrer auprès de l'API et recevoir un *
*JWT (JSON Web Token)**. Ce JWT doit être utilisé pour toutes les requêtes API ultérieures.

**Endpoint :** `POST /api/v1/worker/register`

#### Corps de la Requête

Fournissez le jeton d'enregistrement reçu de l'administrateur.

```json
{
  "token": "a1b2c3d4-e5f6-7890-1234-567890abcdef"
}
```

#### Réponse de Succès (200 OK)

L'API retourne un JWT. Stockez ce jeton de manière sécurisée dans votre worker.

```json
{
  "success": true,
  "message": "Worker authenticated successfully",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

## Étape 3 : Effectuer des Requêtes Authentifiées

Pour toutes les futures requêtes vers des endpoints protégés, le worker doit inclure le JWT dans l'en-tête
`Authorization`.

**Format de l'en-tête :**
`Authorization: Bearer <votre_jeton_jwt>`

---

## Étape 4 : Cycle de Vie du Traitement des Tâches

Une fois authentifié, le worker peut commencer à récupérer et à traiter des tâches.

### 4.1. Obtenir une Tâche en Attente

Le worker interroge cet endpoint pour demander une tâche à l'API.

**Endpoint :** `GET /api/v1/tasks/pending`
**Authentification :** JWT Worker requis.

#### Réponse de Succès (200 OK)

Si une tâche est disponible, l'API l'assigne au worker, change son statut à `processing`, et retourne les détails de la
tâche. Le worker doit alors commencer à traiter le `sourceFile`.

```json
{
  "success": true,
  "message": "Task assigned successfully",
  "data": {
    "id": "clv9p8qjk000108l9e7g3f2a1",
    "config": "{ \"param\": \"value\" }",
    "status": "processing",
    "sourceFileId": "clv9p5n3h000008l9c4f2g8h1",
    "resultFileId": null,
    "workerId": "clv9p3b5a000008l9g2h7d4e2",
    "createdAt": "2025-09-16T20:20:00Z",
    "updatedAt": "2025-09-16T20:25:00Z"
  }
}
```

Si aucune tâche n'est en attente, l'API retournera un `404 Not Found`.

### 4.2. Mettre à Jour le Statut de la Tâche (Optionnel)

Le worker peut envoyer des mises à jour de statut ou signaler des erreurs non fatales pendant le traitement.

**Endpoint :** `PATCH /api/v1/tasks/:cuid`
**Authentification :** JWT Worker requis.

#### Corps de la Requête

```json
{
  "status": "processing",
  "status_message": "Traitement en cours, 50% terminé."
}
```

Pour marquer la tâche comme échouée, utilisez le statut `failed`.

```json
{
  "status": "failed",
  "status_message": "Le traitement a échoué : le fichier source est corrompu."
}
```

### 4.3. Téléverser le Résultat de la Tâche

Lorsque le traitement est terminé, le worker téléverse le fichier résultant.

**Endpoint :** `POST /api/v1/tasks/:cuid`
**Authentification :** JWT Worker requis.

#### Corps de la Requête

Cet endpoint attend une requête `multipart/form-data` contenant le fichier de résultat.

- **Nom du champ de formulaire :** `file`
- **Contenu :** Les données binaires du fichier de résultat.

Après un téléversement réussi, l'API marque automatiquement le statut de la tâche comme `completed`.

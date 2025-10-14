# GEMINI.md — SDK Go pour l’API REST qalpuch

## 1. Contexte & Objectifs

Ce dépôt contient le SDK officiel en Go pour consommer l’API REST qalpuch.
Objectifs :

- Fournir une interface idiomatique Go pour toutes les ressources de l’API (`Users`, `Files`, `Tasks`, `Workers`).
- Respecter le **Standard Go Project Layout** afin d’assurer une structure claire et extensible.
- Appliquer les principes de conception **GoF (Gang of Four)**, notamment Factory, Strategy et Adapter, afin d’obtenir
  un code maintenable et évolutif.

## 2. Architecture

- **`cmd/`** : binaries et CLI de test/démo.
- **`internal/`** : logique privée au SDK.
- **`pkg/`** : modules publics exposés (clients, modèles, utils).
- **`api/`** : définition des services API (Users, Files, Tasks, Workers).
- **`examples/`** : snippets pour intégration rapide.
- **`docs/`** : documentation générée + guides (`api_reference.md`, `worker_integration.md`).

Chaque ressource de l’API aura :

- Un **client struct** (e.g. `UserClient`, `FileClient`).
- Des **méthodes claires** mappant directement les endpoints REST (`CreateUser`, `GetFiles`, `RegisterWorker`, etc.).
- Une **gestion d’erreurs centralisée** basée sur les codes HTTP et la structure JSON standardisée.

## 3. Fichiers & Dossiers Clés

- `api_reference.md` : référence des endpoints REST.
- `worker_integration.md` : protocole spécifique d’intégration des workers.
- `pkg/clients/` : code générant et gérant les requêtes HTTP.
- `pkg/models/` : définitions des structs Go représentant les entités (`User`, `File`, `Task`, `Worker`).
- `pkg/errors/` : gestion des erreurs communes.

## 4. Conventions de Code

- Respect du **formatage Go (`go fmt`)** et du **linting (`golangci-lint`)**.
- Nommage clair et cohérent : les méthodes reflètent les endpoints (`CreateTask`, `UploadFile`).
- Chaque méthode expose un contexte (`context.Context`) pour cancellations et deadlines.
- Tests unitaires obligatoires pour chaque ressource, avec couverture >80%.

## 5. Workflow de Développement

1. Ajouter ou modifier une ressource API en suivant la doc (`api_reference.md`).
2. Implémenter la struct + méthodes dans `pkg/clients/<ressource>.go`.
3. Ajouter les modèles dans `pkg/models/<ressource>.go`.
4. Écrire les tests correspondants dans `pkg/tests/`.
5. Documenter l’usage dans `examples/`.

## 6. Contraintes & Limites

- Ne pas casser la compatibilité publique (`pkg/`).
- Ne pas modifier `api_reference.md` ni `worker_integration.md` : ces fichiers sont sources d’autorité.
- Toujours passer par les interfaces définies, ne jamais coder directement contre `net/http` dans le code utilisateur.

## 7. Patterns GoF appliqués

- **Factory** : création des clients (e.g. `NewUserClient(baseURL, token)`).
- **Adapter** : uniformiser les réponses API JSON → structs Go.
- **Strategy** : gestion de la réauthentification ou du retry HTTP configurable.

## 8. Protocoles IA / Contributions Assistées

Lorsqu’un agent IA intervient sur ce dépôt :

- Lire attentivement `api_reference.md` et `worker_integration.md`.
- Ne pas toucher au `internal/` sans validation humaine.
- Pour tout ajout de ressource :

1. Générer un plan (structs + méthodes).
2. Proposer un patch de code.
3. Ajouter tests + exemples.

## 9. Zones ouvertes

- Support futur du streaming temps réel (`WebSockets` ou `SSE`).

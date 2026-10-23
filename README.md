# learn_go_with_api

API REST en Go pour l'inscription de clients, construite selon une architecture hexagonale (ports & adapters).

## Stack

- Go 1.26
- [Gin](https://github.com/gin-gonic/gin) — framework HTTP
- [GORM](https://gorm.io/) + PostgreSQL ([gorm.io/driver/postgres](https://pkg.go.dev/gorm.io/driver/postgres))
- [godotenv](https://github.com/joho/godotenv) — chargement des variables d'environnement
- [google/uuid](https://github.com/google/uuid)

## Architecture

```
cmd/
  main.go                          point d'entrée

internal/
  domain/
    model/                         entités métier (Customer) et erreurs du domaine
    rule/                          règles de validation (ex. format d'email)
  application/
    port/in/                       interfaces d'entrée (use cases)
    port/out/                      interfaces de sortie (repository, hasher)
    service/                       implémentation des use cases (RegisterService)
  adapter/
    web/in/controllers/            controllers HTTP (Gin)
    web/in/dto/                    DTOs de requête/réponse
    web/out/persistance/           implémentation GORM/PostgreSQL du repository
```

Le sens des dépendances va des adapters vers le domaine : les controllers dépendent des ports (`application/port/in`), les services dépendent des ports de sortie (`application/port/out`), et le domaine ne dépend de rien d'autre dans le projet.

## Configuration

Copier `.env.example` en `.env` et renseigner les variables :

```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=learn_go_with_api
DB_PORT=5432
DB_SSLMODE=disable
```

La connexion à la base (et l'auto-migration de l'entité `Customer`) est établie dans [internal/adapter/web/out/persistance/config/gorm_config.go](internal/adapter/web/out/persistance/config/gorm_config.go).

## Lancer le projet

```bash
go mod download
go run ./cmd
```

Le serveur démarre par défaut sur `http://localhost:8080`.

## Fonctionnalité : inscription d'un client

- Route : `POST /register`
- Use case : [internal/application/service/register.go](internal/application/service/register.go)
- Controller : [internal/adapter/web/in/controllers/auth_controller.go](internal/adapter/web/in/controllers/auth_controller.go)
- Hash du mot de passe : [internal/adapter/security/bcrypt_password_hasher.go](internal/adapter/security/bcrypt_password_hasher.go) (bcrypt)

Règles appliquées lors de l'inscription :

| Cas | Erreur domaine | Code HTTP |
|---|---|---|
| Email déjà utilisé | `model.ErrEmailAlreadyUsed` | 409 Conflict |
| Format d'email invalide | `model.ErrInvalidEmailFormat` | 400 Bad Request |
| Requête mal formée (JSON invalide) | — | 400 Bad Request |
| Autre erreur (persistance, hashing, ...) | — | 500 Internal Server Error |

## Tests

```bash
go test ./...
```

## Build

```bash
go build ./...
```

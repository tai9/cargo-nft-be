# CARGO NFT BACKEND 

## FE integration:
- Repo: https://github.com/tai9/cargo-nft-fe
- Preview: https://cargo-nft-fe.vercel.app/

## How to start project ?

To start backend service, you need to install docker and golang-migrate first.

1. Install golang-migrate

```sh
brew install golang-migrate
```

2. Create dockerdb container

```sh
make postgres
```

3. Create database instance

```sh
make createdb
```

We also provide a cli to drop database if you need to refresh project

```sh
make dropdb
```
 
 4. Migrate schema and master data

```sh
make migrateup
```

5. Generate sqlc

```sh
make sqlc
```

6. Start web service

```sh
make server
```

## Usefull commands

### Migrations:
1. Create migrations:
   ```sh
   make migrations name={ARGS}
   ```
2. Migration up:
    ```sh
    make migrationup
    ```
3. Migration down:
   ```sh
   make migrationdown
   ```

### Deployment:
1. Docker build:
   ```sh
   make docker.build
   ```
2. Push docker image to registry:
   ```sh
   make docker.push
   ```
3. Deploy to staging env:
   ```sh
   make deploy.staging
   ```

### Authorization
Library: https://github.com/casbin/casbin

Model:
- `permission`: define all system permissions. Permission can have `type` or not. 
Currently, there is only one `type`, this is `CRUD`, it represents 4 action: read, update, create, delete with 
following bitmask. In contrast, if `type` is `null`, this means action is `can`.


|action | bit |
|---|---|
| create | 8 // 0b1000 |
| read   | 4 // 0b0100 |
| update | 2 // 0b0010 |
| delete | 1 // 0b0001 |

## CONVENTION
- Package name: https://go.dev/blog/package-names


## Server Quick Start
1. `cd api`
2. install golang
3. install hivemind: https://github.com/DarthSim/hivemind
4. `cp .env.example .env`
5. make run


## API Client Quick start (must be done before running the widget)
1. `cd api/client`
2. `npm i`

## Run tests

```bash
make test

```

## Environment

hivemind looks for `.env` and there is an example of this file you can copy `.env.example`

## Firestore Emulator

### Quick 'N Dirty Firestore Emulator (port 8080)
```bash
curl -O  https://storage.googleapis.com/firebase-preview-drop/emulator/cloud-firestore-emulator-v1.11.7.jar
java -jar cloud-firestore-emulator-v1.11.7.jar
```

### Full-Feature Emulator (port 8080 & 4000)
boots an additional gui server on port `4000` which allows you to brows, view, and edit data in the emulation db.
```bash
curl -sL https://firebase.tools | bash
firebase login
firebase emulators:start
```


## Codegen

### dependencies (incomplete atm)
swagger client codegen Mac setup
```bash
npm install @openapitools/openapi-generator-cli -g
```

### Generate client & protocol
this is required whenever api.proto is updated and/or lib/swagger/merge.json
```bash
make proto
```

### Generate dependency injections
generates wire_gen.go from the spec in wire.go
this is required when wiring up new dependencies into our server container
```bash
make wire
```
### Swagger Documents
http://localhost:5100/swagger

### LINKS
https://docs.microsoft.com/en-us/sql/relational-databases/security/encryption/always-encrypted-cryptography?view=sql-server-ver15

Deterministic encryption is more effective in concealing patterns, compared to alternatives, such as using a pre-defined IV value.


https://tools.ietf.org/html/rfc5297

https://github.com/miscreant/meta/wiki/AES-SIV



```
import (
	"github.com/disintegration/imaging"
)

func ProvideFoo() {
	// Resize srcImage to size = 128x128px using the Lanczos filter.
	dstImage128 := imaging.Resize(srcImage, 128, 128, imaging.Lanczos)

}

```

### pubsub emulator
may be needed in the future...

```
gcloud components install pubsub-emulator
gcloud components update
```

job message examples

```
{"kind": "CREATE_WYRE_ACCOUNT_FOR_USER", "relatedIDs": ["<user_id>"]}
```


## secrets management

### load up dev env from local into cloud secret in correct project env

```
bin/load-snap-env-secret --project silken-phalanx-305703 --data-file=secrets/.env.dev
```

### cleanup old versions
```
bin/get-active-versions-from-snap-env-secret --project silken-phalanx-305703 | xargs -n1 bin/delete-snap-env-version
```

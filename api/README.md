## Server Quick Start
1. `cd client`
2. install golang
3. install hivemind: https://github.com/DarthSim/hivemind
4. `cp .env.example .env`
5. make run


# Client Quick start
1. `cd client`
2. `cp .env.example .env`
3. `npm i`
4. `npm run dev`

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
firebase login
firebase emulators:start
```


## Swagger Client Codegen

### Mac setup
```bash
npm install @openapitools/openapi-generator-cli -g
```

### Generate client
```bash
openapi-generator-cli generate -g typescript -i lib/swagger/swagger.json -o client
```


## Quick Start

1. install golang
2. install hivemind: https://github.com/DarthSim/hivemind
3. make run

## Environment

grpcserver & hivemind looks for `.env` and there is an example of this file you can copy `.env.example`

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


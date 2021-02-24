1. install golang
2. install hivemind: https://github.com/DarthSim/hivemind
3. `bin/build && hivemind`


NOTES

grpcserver looks for `.env` and there is an example of this file you can copy `.env.example`

# Firestore Emulator
> curl -O  https://storage.googleapis.com/firebase-preview-drop/emulator/cloud-firestore-emulator-v1.11.7.jar
> java -jar cloud-firestore-emulator-v1.11.7.jar

## Full-Feature Emulator
boots a server on port `4000`
> firebase login
> firebase emulators:start
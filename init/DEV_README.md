# Snap Wallet Init

This directory contains the initialization script `index.ts` which constructs an iframe for web.
The `index.html` file is a working demo of the widget being launched using the `Snap` class from `index.ts`

## Up and Running

### Environment

Make sure that you have a `.env` file in the `init/` directory. Environment variable examples can be found in `.env.example`.

### Build

Run the build script in order to output `dist/`. This is where the script will live locally.

```bash
  npm run build
```

### Serve

Serve the init project directory

- Python 2
  ```bash
  python -m SimpleHTTPServer
  ```
- Python 3
  ```bash
  python3 -m http.server
  ```

and finally browse to [server](http://localhost:8000)

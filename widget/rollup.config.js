import svelte from 'rollup-plugin-svelte'
import commonjs from '@rollup/plugin-commonjs'
import resolve from '@rollup/plugin-node-resolve'
import livereload from 'rollup-plugin-livereload'
import { terser } from 'rollup-plugin-terser'
import sveltePreprocess from 'svelte-preprocess'
import typescript from '@rollup/plugin-typescript'
import css from 'rollup-plugin-css-only'
import replace from '@rollup/plugin-replace'
import dotenv from 'dotenv'
import svelteSVG from 'rollup-plugin-svelte-svg'
import { nanoid } from 'nanoid'
import { writeFileSync } from 'fs'

// Dotenv will also pick up api .env if path not defined
dotenv.config({ path: './.env' })

const production = !process.env.ROLLUP_WATCH

function serve() {
  let server

  function toExit() {
    if (server) server.kill(0)
  }

  return {
    writeBundle() {
      if (server) return
      server = require('child_process').spawn(
        'npm',
        ['run', 'start', '--', '--dev'],
        {
          stdio: ['ignore', 'inherit', 'inherit'],
          shell: true,
        },
      )

      process.on('SIGTERM', toExit)
      process.on('exit', toExit)
    },
  }
}

const BuildID = nanoid()
const jsBundleName = production ? `bundle.${BuildID}.js` : 'bundle.js'
const cssBundleName = production ? `bundle.${BuildID}.css` : 'bundle.css'

if (production) {
  writeFileSync(
    'dist/info.json',
    JSON.stringify({
      buildID: BuildID,
      jsBundleName,
      cssBundleName,
    }),
  )
}

export default {
  input: 'src/main.ts',
  output: {
    sourcemap: !production,
    format: 'iife',
    name: 'snap',
    dir: production ? 'dist' : 'public/build',
    entryFileNames: jsBundleName,
  },
  plugins: [
    replace({
      preventAssignment: true,
      __ENV: JSON.stringify({
        DEBUG: process.env.DEBUG,
        API_BASE_URL: process.env.API_BASE_URL,
        API2_BASE_URL: process.env.API2_BASE_URL,
        WYRE_BASE_URL: process.env.WYRE_BASE_URL,
      }),
    }),
    svelteSVG(),
    svelte({
      preprocess: sveltePreprocess({ sourceMap: !production, sass: true }),
      compilerOptions: {
        // enable run-time checks when not in production
        dev: !production,
      },
      onwarn: (warning, handler) => {
        // don't warn 'A11y: A form label must be associated with a control'
        if (
          ['a11y-label-has-associated-control', 'a11y-autofocus'].includes(
            warning.code,
          )
        )
          return
        // let Rollup handle all other warnings normally
        handler(warning)
      },
    }),
    // we'll extract any component CSS out into
    // a separate file - better for performance
    css({ output: cssBundleName }),

    // If you have external dependencies installed from
    // npm, you'll most likely need these plugins. In
    // some cases you'll need additional configuration -
    // consult the documentation for details:
    // https://github.com/rollup/plugins/tree/master/packages/commonjs
    resolve({
      browser: true,
      dedupe: ['svelte'],
    }),
    commonjs(),
    typescript({
      sourceMap: !production,
      inlineSources: !production,
    }),

    // In dev mode, call `npm run start` once
    // the bundle has been generated
    !production && serve(),

    // Watch the `public` directory and refresh the
    // browser on changes when not in production
    !production && livereload('public'),

    // If we're building for production (npm run build
    // instead of npm run dev), minify
    production && terser(),
  ],
  watch: {
    clearScreen: false,
  },
}

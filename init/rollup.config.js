// rollup.config.js
import typescript from '@rollup/plugin-typescript'
import resolve from 'rollup-plugin-node-resolve'
import replace from '@rollup/plugin-replace'
import commonjs from '@rollup/plugin-commonjs'
import dotenv from 'dotenv'
import { nanoid } from 'nanoid'
import { writeFileSync } from 'fs'

dotenv.config()

if (!process.env.WIDGET_URL)
  throw new Error('Please set a WIDGET_URL and rebuild: `npm run build`')

const BuildID = nanoid()
const initBundleName = `init.${BuildID}.js`
const initBundlePath = `/widget/dist/${initBundleName}`
const firebaseHostingConfig = {
  hosting: {
    headers: [],
    rewrites: [],
    redirects: [
      {
        source: `/widget/dist/init.js`,
        destination: initBundlePath,
        type: 302,
      },
    ],
  },
}

writeFileSync(
  'dist/info.json',
  JSON.stringify({
    initBundleName,
    initBundlePath,
  })
)
writeFileSync(
  'dist/firebase-hosting-config.json',
  JSON.stringify(firebaseHostingConfig)
)

export default {
  input: 'index.ts',
  output: {
    dir: './dist',
    format: 'umd',
    name: 'init.js',
    entryFileNames: initBundleName,
  },
  plugins: [
    replace({
      preventAssignment: true,
      _ENV: JSON.stringify({
        WIDGET_URL: process.env.WIDGET_URL,
        API_BASE_URL: process.env.API_BASE_URL,
      }),
    }),
    typescript(),
    resolve({
      jsnext: true,
      browser: true,
    }),
    commonjs(),
  ],
}

// rollup.config.js
import typescript from '@rollup/plugin-typescript'
import resolve from 'rollup-plugin-node-resolve'
import replace from '@rollup/plugin-replace'
import dotenv from 'dotenv'

dotenv.config()

export default {
  input: 'index.ts',
  output: {
    dir: './dist',
    format: 'umd',
    name: 'init.js',
  },
  plugins: [
    replace({
      preventAssignment: true,
      _ENV: JSON.stringify({
        WIDGET_URL: process.env.WIDGET_URL,
      }),
    }),
    typescript(),
    resolve({
      jsnext: true,
      browser: true,
    }),
  ],
}

// rollup.config.js
import typescript from '@rollup/plugin-typescript'
import resolve from 'rollup-plugin-node-resolve'
import replace from '@rollup/plugin-replace'
import dotenv from 'dotenv'

dotenv.config()

if (!process.env.WIDGET_URL) throw new Error('Please set a WIDGET_URL and rebuild: `npm run build`')

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

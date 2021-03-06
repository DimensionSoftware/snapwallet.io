// rollup.config.js
import typescript from '@rollup/plugin-typescript'
import resolve from 'rollup-plugin-node-resolve'

export default {
  input: 'index.ts',
  output: {
    dir: './dist',
    format: 'umd',
    name: 'init.js',
  },
  plugins: [
    typescript(),
    resolve({
      jsnext: true,
      browser: true,
    }),
  ],
}

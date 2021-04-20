const sveltePreprocess = require('svelte-preprocess')

// This file is only used by Jest
module.exports = {
  preprocess: sveltePreprocess({
    sass: true,
  }),
}

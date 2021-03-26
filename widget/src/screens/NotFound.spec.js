import { mount } from 'cypress-svelte-unit-test'

import NotFound from '../../src/screens/NotFound.svelte'

it('div contains oops', () => {
  mount(NotFound, {
  })
  cy.contains('div', 'Oops')
})

import { mount } from 'cypress-svelte-unit-test'

import NotFound from '../../src/screens/NotFound.svelte'

it('shows greeting', () => {
  mount(NotFound, {
  })
  //cy.contains('h1', 'Not')
})

import NotFound from '../../src/screens/NotFound.svelte'
import { render } from '@testing-library/svelte'

test('shows proper heading when rendered', () => {
  const { getByText } = render(NotFound, {})

  expect(
    getByText("Oops, we're sorry. The page requested is missing."),
  ).toBeInTheDocument()
})

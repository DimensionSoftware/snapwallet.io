import NotFound from '../../src/screens/NotFound.svelte'
import App from '../../src/App.svelte'

import { render } from '@testing-library/svelte'

jest.mock('svelte-spa-router', () => ({
  push: () => {},
  pop: () => {},
  location: () => 'somewhere',
}))

jest.mock('svelte-spa-router/wrap', () => ({
  wrap: () => {},
}))

test('shows proper heading when rendered', () => {
  const { getByText } = render(NotFound, {})

  expect(
    getByText("Oops, we're sorry. The page requested is missing."),
  ).toBeInTheDocument()
})

// TODO: move this somewhere once Jest is running properly
test('shows proper thing', () => {
  const { getByText } = render(App, {
    product: {
      videoURL:
        'https://mkpcdn.com/videos/d3a277f4e6f1212c900a1da4ec915aa9_675573.mp4',
      destinationAmount: 0.04,
      destinationTicker: 'ETH',
      destinationAddress: '0xf636B6aA45C554139763Ad926407C02719bc22f7',
      title: 'The Crown',
      author: 'Patrick Mahomes',
    },
  })

  expect(getByText('By Patrick Mahomes')).toBeInTheDocument()
})

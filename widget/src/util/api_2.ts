// api v2 helpers
export const get = (...params) => req.apply(null, ['GET', ...params])
export const post = (...params) => req.apply(null, ['POST', ...params])
export const req = async (method, path, body) => {
  const authorization = 'Bearer ' + (await window.AUTH_MANAGER.getAccessToken())
  // TODO x-snap-wallet-api-key, etc...
  const res = await fetch(`${__ENV['API2_BASE_URL']}/v1/${path}`, {
    method,
    // mode: 'cors',
    headers: {
      accept: '*/*',
      'accept-encoding': 'gzip, deflate, br',
      'content-type': 'application/json',
      connection: 'keep-alive',
      authorization,
      'x-snap-wallet-api-key': 'b31e0ded-ee1a-41a8-8860-13737acff8a7',
    },
    body: JSON.stringify(body),
  })
  return res.json()
}

// api v2 helpers
export const get = (...params) => req.apply(null, ['GET', ...params])
export const post = (...params) => req.apply(null, ['POST', ...params])
export const req = async (method, path, body) => {
  const authorization = 'Bearer ' + (await window.AUTH_MANAGER.getAccessToken())
  // TODO setup api v2 paths in deploy-web, etc...
  console.log('method', method, 'path', path, 'body', JSON.stringify(body))
  const // res = await fetch(`${__ENV['API_BASE_URL']}/v1/${path}`, {
    // res = await fetch(`https://snap-api2-4eumnbid2a-uc.a.run.app/v1/${path}`, {
    res = await fetch(`http://localhost:3001/v1/${path}`, {
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

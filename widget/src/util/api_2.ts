// api v2 helpers
export const get = (...params) => req('get', params)
export const post = (...params) => req('post', params)
export const req = async (method, path) => {
  const authorization =
      'Bearer ' + (await window.AUTH_MANAGER.getAccessToken()),
    // TODO setup api v2 paths in deploy-web, etc...
    // const res = await fetch(`${__ENV['API_BASE_URL']}/v1/${path}`, {
    res = await fetch(`https://snap-api2-4eumnbid2a-uc.a.run.app/v1/${path}`, {
      method,
      mode: 'cors',
      headers: {
        accept: '*/*',
        'accept-encoding': 'gzip, deflate, br',
        connection: 'keep-alive',
        authorization,
        'X-Snap-Wallet-Api-Key': 'ac8f0e56-e03f-408a-88f6-dea8ef6e2c9d',
      },
    })
  return res.json()
}

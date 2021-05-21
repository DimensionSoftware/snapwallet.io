import {
  createConfiguration,
  FluxApi,
  RequestContext,
  SecurityAuthentication,
  ServerConfiguration,
  TokenExchangeResponse,
  TokenMaterial,
} from 'api-client'
import { JWT_TOKENS_KEY, JWT_TOKENS_LOCK_KEY } from './constants'
import { Logger, parseJwt } from './util'

export class AuthManager {
  // to avoid cycle
  private readonly unauthenticatedAPI = genAPIClient()
  private readonly prelogoutThreshold = 5 * 60 * 1000

  // to avoid duplicate calls, we can only use a refresh token to exchange once
  private tokenExchangePromise?: Promise<TokenExchangeResponse>

  private setTokens(access: string, refresh: string) {
    if (access && refresh) {
      Logger.debug('New access token being set...', access)
      Logger.debug('New refresh token being set...', refresh)
      window.localStorage.setItem(JWT_TOKENS_KEY, [access, refresh].join(':'))
    } else {
      window.localStorage.removeItem(JWT_TOKENS_KEY)
    }
  }

  private tokenIsLocked(): boolean {
    const lockedAtStr = window.localStorage.getItem(JWT_TOKENS_LOCK_KEY)
    if (!lockedAtStr) {
      return false
    }

    const lockedAt = parseInt(lockedAtStr) // epoch ms
    if (isNaN(lockedAt)) {
      this.clearTokenLock('Token lock cleared: invalid timestamp')
      return false
    }

    // wait 10 seconds before force unlocking
    const cutoffTime = Date.now() - 10 * 1000
    if (lockedAt < cutoffTime) {
      this.clearTokenLock('Token lock cleared: 10 seconds elapsed')
      return false
    }

    return true
  }

  private clearTokenLock(reason: string = 'Token lock cleared') {
    Logger.debug(reason)
    return window.localStorage.removeItem(JWT_TOKENS_LOCK_KEY)
  }

  private setTokenLock() {
    Logger.debug('Token lock set')
    return window.localStorage.setItem(
      JWT_TOKENS_LOCK_KEY,
      Date.now().toString(),
    )
  }

  // returns access, refresh
  private getCurrentTokens(): [string, string] {
    const data = window.localStorage.getItem(JWT_TOKENS_KEY)
    if (data) {
      return data.split(':') as [string, string]
    }

    return ['', '']
  }

  private parseTokenClaims(token: string): { [k: string]: any } | null {
    if (!token) {
      return null
    }

    const parsed = parseJwt(token)
    if (!parsed) {
      Logger.warn('could not parse jwt')
      return null
    }

    return parsed
  }

  private tokenIsExpired(token: string): boolean {
    const parsed = this.parseTokenClaims(token)
    if (!parsed) {
      Logger.debug('parseTokenClaims returned empty parsed value')
      return true
    }

    const exp = parseInt(parsed.exp)
    if (isNaN(exp)) {
      Logger.debug('exp is NaN', 'original:', parsed.exp, 'parsed:', exp)
      return true
    }

    const tokenExpiresAt = exp * 1000

    return addEpochBuffer(Date.now()) > tokenExpiresAt
  }

  private tokenIsExpiredSoon(token: string): boolean {
    const parsed = this.parseTokenClaims(token)
    if (!parsed) {
      Logger.debug('parseTokenClaims returned empty parsed value')
      return true
    }

    const exp = parseInt(parsed.exp)
    if (isNaN(exp)) {
      Logger.debug(
        'refresh exp is NaN',
        'original:',
        parsed.exp,
        'parsed:',
        exp,
      )
      return true
    }

    const tokenExpiresAt = exp * 1000

    return Date.now() + this.prelogoutThreshold > tokenExpiresAt
  }

  // exchanges and updates tokens -- and makes sure only one is in flight at a time to not violate RTR
  private async tokenExchange(token: string): Promise<[string, string]> {
    if (!this.tokenExchangePromise) {
      Logger.debug('No tokenExchangePromise found')

      Logger.debug('Setting token exchange promise')
      this.tokenExchangePromise = this.unauthenticatedAPI.fluxTokenExchange({
        refreshToken: token,
      })
      Logger.debug('Awaiting token exchange promise')
      const resp = await this.tokenExchangePromise

      Logger.debug('Setting current access token', resp.tokens.accessToken)
      Logger.debug('Setting current refresh token', resp.tokens.refreshToken)
      this.setTokens(resp.tokens.accessToken, resp.tokens.refreshToken)

      Logger.debug('Nullifying tokenExchangePromise')
      this.tokenExchangePromise = null

      return [resp.tokens.accessToken, resp.tokens.refreshToken]
    } else {
      Logger.debug('Awaiting an existing tokenExchangePromise')
      await this.tokenExchangePromise

      return this.getCurrentTokens()
    }
  }

  // returns '' if token is non-refreshable or completely expired
  // if refreshable and expired: we will refresh the token before returning it
  public async getAccessToken(): Promise<string> {
    const [access, refresh] = this.getCurrentTokens()
    if (access && refresh) {
      if (this.tokenIsExpired(access)) {
        Logger.debug('Access token is expired')

        if (this.tokenIsExpired(refresh)) {
          Logger.debug('Refresh token is expired')

          this.logout()
          return ''
        }

        Logger.debug('Running token exchange')

        if (!this.tokenIsLocked()) {
          this.setTokenLock()
          const accessToken = (await this.tokenExchange(refresh))[0]
          this.clearTokenLock()
          return accessToken
        } else {
          // todo, timeout instead of infinite recursion if lock isnt cleared
          await delay(500)
          return this.getAccessToken()
        }
      } else {
        return access
      }
    }

    return ''
  }

  // for first time (unrefreshed, interactive) login to populate access/refresh tokens
  public login(tokens: TokenMaterial) {
    this.setTokens(tokens.accessToken, tokens.refreshToken)
  }

  public logout() {
    Logger.debug('Logout was called. Resetting access and refresh tokens.')
    this.setTokens('', '')

    Logger.debug('Dispatching logout event.')
    window.dispatchEvent(new Event('logout'))
  }

  public viewerIsLoggedIn(): boolean {
    Logger.debug('viewerIsLoggedIn called')

    const [_, refresh] = this.getCurrentTokens()

    return !this.tokenIsExpired(refresh)
  }

  // will return '' if user is not logged in
  public viewerUserID(): string {
    const [_, refresh] = this.getCurrentTokens()

    const parsed = this.parseTokenClaims(refresh)
    if (!parsed) {
      return ''
    }

    if (!parsed.sub) {
      return ''
    }

    return parsed.sub
  }

  public getSessionExpiration(): number {
    const [_, refresh] = this.getCurrentTokens()
    const parsed = this.parseTokenClaims(refresh)
    if (!parsed) {
      Logger.debug('parseRefreshTokenClaims returned empty parsed value')
      return 0
    }

    const exp = parseInt(parsed.exp)
    if (isNaN(exp)) {
      Logger.debug(
        'refresh exp is NaN',
        'original:',
        parsed.exp,
        'parsed:',
        exp,
      )
      return 0
    }

    const refreshTokenExpiresAt = exp * 1000
    return refreshTokenExpiresAt
  }

  // watch for session changes so a logout message can be emitted
  public watch() {
    setInterval(() => {
      const [_, refresh] = this.getCurrentTokens()

      if (refresh === '') {
        Logger.warn('refresh token was empty string, returning early...')

        return
      }

      if (this.tokenIsExpired(refresh)) {
        Logger.warn(
          'auth watcher is logging out user due to expired refresh token',
        )

        this.logout()
        return
      } else {
        if (this.tokenIsExpiredSoon(refresh)) {
          window.dispatchEvent(new Event('prelogout'))
        }
      }
    }, 30 * 1000)
  }
}

export class FluxBearerAuthentication implements SecurityAuthentication {
  public constructor(private readonly manager: AuthManager) {}

  public getName(): string {
    return 'Bearer'
  }

  public async applySecurityAuthentication(context: RequestContext) {
    const token = await this.manager.getAccessToken()

    if (token) {
      context.setHeaderParam('Authorization', `Bearer ${token}`)
      return
    }

    Logger.warn(
      'route needs authentication but no token was provided by the auth manager',
    )
    throw new Error('Please sign in.')
  }
}

export function genAPIClient(authManager?: AuthManager): FluxApi {
  const config = createConfiguration({
    baseServer: new ServerConfiguration(__ENV.API_BASE_URL, {}),
  })
  if (authManager) {
    config.authMethods.Bearer = new FluxBearerAuthentication(authManager)
  }

  return new FluxApi(config)
}

function addEpochBuffer(epoch: number): number {
  return epoch + 5 * 1000
}

function delay(ms: number): Promise<void> {
  return new Promise(resolve => {
    setTimeout(resolve, ms)
  })
}

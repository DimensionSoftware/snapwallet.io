import {
  createConfiguration,
  FluxApi,
  RequestContext,
  SecurityAuthentication,
  ServerConfiguration,
  TokenExchangeResponse,
  TokenMaterial,
} from 'api-client'
import { JWT_ACCESS_TOKEN_KEY, JWT_REFRESH_TOKEN_KEY } from './constants'
import { Logger, parseJwt } from './util'

export class AuthManager {
  // to avoid cycle
  private readonly unauthenticatedAPI = genAPIClient()
  private readonly prelogoutThreshold = 5 * 60 * 1000

  private sessionExpiresAt = 0

  constructor() {
    const parsed = this.parseRefreshTokenClaims()
    if (!parsed) {
      return
    }
    if (!parsed.exp) {
      return
    }
    this.sessionExpiresAt = parseInt(parsed.exp) * 1000
  }

  // to avoid duplicate calls, we can only use a refresh token to exchange once
  private tokenExchangePromise?: Promise<TokenExchangeResponse>

  private setCurrentAccessToken(newToken: string) {
    if (newToken) {
      window.localStorage.setItem(JWT_ACCESS_TOKEN_KEY, newToken)
    } else {
      window.localStorage.removeItem(JWT_ACCESS_TOKEN_KEY)
    }
  }
  private setCurrentRefreshToken(newToken: string) {
    if (newToken) {
      window.localStorage.setItem(JWT_REFRESH_TOKEN_KEY, newToken)
    } else {
      window.localStorage.removeItem(JWT_REFRESH_TOKEN_KEY)
    }
  }

  private getCurrentAccessToken(): string {
    return window.localStorage.getItem(JWT_ACCESS_TOKEN_KEY) || ''
  }
  private getCurrentRefreshToken(): string {
    return window.localStorage.getItem(JWT_REFRESH_TOKEN_KEY) || ''
  }

  private parseRefreshTokenClaims(): { [k: string]: any } | null {
    const token = this.getCurrentRefreshToken()
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

  private parseAccessTokenClaims(): { [k: string]: any } | null {
    const token = this.getCurrentAccessToken()
    if (!token) {
      Logger.debug('getCurrentAccessToken returned empty token')
      return null
    }

    const parsed = parseJwt(token)
    if (!parsed) {
      Logger.warn('could not parse jwt')
      return null
    }

    return parsed
  }

  private accessTokenIsExpired(): boolean {
    const parsed = this.parseAccessTokenClaims()
    if (!parsed) {
      Logger.debug('parseAccessTokenClaims returned empty parsed value')
      return true
    }

    const exp = parseInt(parsed.exp)
    if (isNaN(exp)) {
      Logger.debug('exp is NaN', 'original:', parsed.exp, 'parsed:', exp)
      return true
    }

    const accessTokenExpiresAt = exp * 1000

    return addEpochBuffer(Date.now()) > accessTokenExpiresAt
  }

  private refreshTokenIsExpired(): boolean {
    return addEpochBuffer(Date.now()) > this.sessionExpiresAt
  }

  private refreshTokenIsExpiredSoon(): boolean {
    return Date.now() + this.prelogoutThreshold > this.sessionExpiresAt
  }

  // exchanges and updates tokens -- and makes sure only one is in flight at a time to not violate RTR
  private async tokenExchange(): Promise<void> {
    if (!this.tokenExchangePromise) {
      Logger.debug('No tokenExchangePromise found')
      const token = this.getCurrentRefreshToken()
      if (!token) {
        Logger.debug('getRefreshToken returned empty token')
        return
      }

      Logger.debug('Setting token exchange promise')
      this.tokenExchangePromise = this.unauthenticatedAPI.fluxTokenExchange({
        refreshToken: token,
      })
      Logger.debug('Awaiting token exchange promise')
      const resp = await this.tokenExchangePromise

      Logger.debug('Setting current access token', resp.tokens.accessToken)
      this.setCurrentAccessToken(resp.tokens.accessToken)
      Logger.debug('Setting current refresh token', resp.tokens.refreshToken)
      this.setCurrentRefreshToken(resp.tokens.refreshToken)

      Logger.debug('Parsing refresh token claims')
      const parsed = this.parseRefreshTokenClaims()
      if (!parsed) {
        throw new Error('could not parse refresh token claims')
      }
      if (!parsed.exp) {
        throw new Error('refresh token claims lacks an expiration')
      }

      this.sessionExpiresAt = parseInt(parsed.exp) * 1000
      Logger.debug('Set session expires at to', this.sessionExpiresAt)

      Logger.debug('Nullifying tokenExchangePromise')
      this.tokenExchangePromise = null
    } else {
      Logger.debug('Awaiting an existing tokenExchangePromise')
      await this.tokenExchangePromise
    }
  }

  // returns '' if token is non-refreshable or completely expired
  // if refreshable and expired: we will refresh the token before returning it
  public async getAccessToken(): Promise<string> {
    if (this.accessTokenIsExpired()) {
      Logger.debug('Access token is expired')
      if (this.refreshTokenIsExpired()) {
        Logger.debug('Refresh token is expired')
        return ''
      } else {
        Logger.debug('Running token exchange')
        await this.tokenExchange()
      }
    }

    return this.getCurrentAccessToken()
  }

  // for first time (unrefreshed, interactive) login to populate access/refresh tokens
  public login(tokens: TokenMaterial) {
    this.setCurrentAccessToken(tokens.accessToken)
    this.setCurrentRefreshToken(tokens.refreshToken)

    const parsed = this.parseRefreshTokenClaims()
    if (!parsed) {
      throw new Error('could not parse refresh token claims')
    }
    if (!parsed.exp) {
      throw new Error('refresh token claims lacks an expiration')
    }
    this.sessionExpiresAt = parseInt(parsed.exp) * 1000
  }

  public logout() {
    Logger.debug('Logout was called. Resetting access and refresh tokens.')
    this.setCurrentAccessToken('')
    this.setCurrentRefreshToken('')
    Logger.debug('Clearing session expires at.')
    this.sessionExpiresAt = 0
    Logger.debug('Dispatching logout event.')
    window.dispatchEvent(new Event('logout'))
  }

  public viewerIsLoggedIn(): boolean {
    Logger.debug('viewerIsLoggedIn called')
    return !this.refreshTokenIsExpired()
  }

  // will return '' if user is not logged in
  public viewerUserID(): string {
    const parsed = this.parseRefreshTokenClaims()
    if (!parsed) {
      return ''
    }

    if (!parsed.sub) {
      return ''
    }

    return parsed.sub
  }

  public getSessionExpiration(): number {
    return this.sessionExpiresAt
  }

  // watch for session changes so a logout message can be emitted
  public watch() {
    setInterval(() => {
      if (this.getCurrentRefreshToken() == '') {
        return
      }

      if (this.refreshTokenIsExpired()) {
        Logger.warn(
          'auth watcher is logging out user due to expired refresh token',
        )

        this.logout()
        return
      } else {
        if (this.refreshTokenIsExpiredSoon()) {
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
    throw new Error('Please login and try again.')
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

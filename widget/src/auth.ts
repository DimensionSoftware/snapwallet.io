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
  private readonly prelogoutThreshold = 60 * 1000

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
      return true
    }

    const exp = parseInt(parsed.exp)
    if (isNaN(exp)) {
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
      const token = this.getCurrentRefreshToken()
      if (!token) {
        return
      }

      this.tokenExchangePromise = this.unauthenticatedAPI.fluxTokenExchange({
        refreshToken: token,
      })
      const resp = await this.tokenExchangePromise
      this.tokenExchangePromise = null

      this.setCurrentAccessToken(resp.tokens.accessToken)
      this.setCurrentRefreshToken(resp.tokens.refreshToken)

      const parsed = this.parseRefreshTokenClaims()
      if (!parsed) {
        throw new Error('could not parse refresh token claims')
      }
      if (!parsed.exp) {
        throw new Error('refresh token claims lacks an expiration')
      }
      this.sessionExpiresAt = parseInt(parsed.exp) * 1000
    } else {
      await this.tokenExchangePromise
    }
  }

  // returns '' if token is non-refreshable or completely expired
  // if refreshable and expired: we will refresh the token before returning it
  public async getAccessToken(): Promise<string> {
    if (this.accessTokenIsExpired()) {
      if (this.refreshTokenIsExpired()) {
        return ''
      } else {
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
    this.setCurrentAccessToken('')
    this.setCurrentRefreshToken('')
    this.sessionExpiresAt = 0
    window.dispatchEvent(new Event('logout'))
  }

  public viewerIsLoggedIn(): boolean {
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
      if (this.refreshTokenIsExpired()) {
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
    }
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

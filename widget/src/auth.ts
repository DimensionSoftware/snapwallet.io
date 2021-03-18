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
import { parseJwt } from './util'

export class AuthManager {
  // to avoid cycle
  private readonly unauthenticatedAPI = genAPIClient()
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

  private accessTokenIsExpired(): boolean {
    const token = this.getCurrentAccessToken()
    if (!token) {
      return true
    }

    const parsed = parseJwt(token)
    if (!parsed) {
      console.log('WARNING: could not parse jwt')
      return true
    }

    const isTimeLeft = parsed.exp > Math.floor(add10seconds(Date.now()) / 1000)

    return !isTimeLeft
  }

  private refreshTokenIsExpired(): boolean {
    const token = this.getCurrentRefreshToken()
    if (!token) {
      return true
    }

    const parsed = parseJwt(token)
    if (!parsed) {
      console.log('WARNING: could not parse jwt')
      return true
    }

    const isTimeLeft = parsed.exp > Math.floor(add10seconds(Date.now()) / 1000)

    return !isTimeLeft
  }

  // exchanges and updates tokens
  private async tokenExchange(): Promise<void> {
    if (!this.tokenExchangePromise) {
      this.tokenExchangePromise = this.unauthenticatedAPI.fluxTokenExchange({
        refreshToken: this.getCurrentRefreshToken(),
      })
      const resp = await this.tokenExchangePromise
      this.tokenExchangePromise = null

      this.setCurrentAccessToken(resp.tokens.accessToken)
      this.setCurrentRefreshToken(resp.tokens.refreshToken)
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
  }

  public logout() {
    this.setCurrentAccessToken('')
    this.setCurrentRefreshToken('')
  }

  public viewerIsLoggedIn(): boolean {
    return !this.refreshTokenIsExpired()
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

function add10seconds(epoch: number): number {
  return epoch + (10 * 1000)
}
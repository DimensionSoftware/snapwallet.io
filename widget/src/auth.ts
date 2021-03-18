import {
  createConfiguration,
  FluxApi,
  RequestContext,
  SecurityAuthentication,
  ServerConfiguration,
  TokenMaterial,
} from 'api-client'
import { JWT_ACCESS_TOKEN_KEY, JWT_REFRESH_TOKEN_KEY } from './constants'
import { parseJwt } from './util'

export class AuthManager {
  // to avoid cycle
  private readonly unauthenticatedAPI = genAPIClient()

  private setCurrentAccessToken(newToken: string) {
    return window.localStorage.setItem(JWT_ACCESS_TOKEN_KEY, newToken)
  }
  private setCurrentRefreshToken(newToken: string) {
    return window.localStorage.setItem(JWT_REFRESH_TOKEN_KEY, newToken)
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

    const isTimeLeft = parsed.exp > Math.floor(Date.now() / 1000)

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

    const isTimeLeft = parsed.exp > Math.floor(Date.now() / 1000)

    return !isTimeLeft
  }

  // exchanges and updates tokens
  private async tokenExchange(): Promise<void> {
    const resp = await this.unauthenticatedAPI.fluxTokenExchange({
      refreshToken: this.getCurrentRefreshToken(),
    })

    this.setCurrentAccessToken(resp.tokens.accessToken)
    this.setCurrentRefreshToken(resp.tokens.refreshToken)
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
}

export class FluxBearerAuthentication implements SecurityAuthentication {
  public constructor(private readonly manager: AuthManager) {}

  public getName(): string {
    return 'Bearer'
  }

  public applySecurityAuthentication(context: RequestContext) {
    if (this.manager.getAccessToken()) {
      context.setHeaderParam('Authorization', `Bearer ${this.accessToken}`)
    }
  }
}

export function genAPIClient(fba?: FluxBearerAuthentication): FluxApi {
  const config = createConfiguration({
    baseServer: new ServerConfiguration(__ENV.API_BASE_URL, {}),
  })
  if (fba) {
    config.authMethods.Bearer = fba
  }

  return new FluxApi(config)
}

const api = genAPIClient()
const manager = new AuthManager(api, fba)

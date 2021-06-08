// @ts-ignore
import QR from 'qr-creator'
// @ts-ignore
import { createConfiguration, FluxApi, ServerConfiguration } from 'api-client'

export type UserIntent = 'buy' | 'sell' | 'donate'
export type SrcDst = 'source' | 'destination'
export enum WidgetEnvironments {
  // ** development ** is only an option for explicitness
  // Simply provide the environment variable INIT_API_BASE_URL for dev
  DEVELOPMENT = 'development',
  SANDBOX = 'sandbox',
  PRODUCTION = 'production',
}

export enum WidgetURLs {
  PRODUCTION = 'https://snapwallet.io/widget',
  SANDBOX = 'https://sandbox.snapwallet.io/widget',
}

export enum APIBaseURLs {
  PRODUCTION = 'https://api.snapwallet.io',
  SANDBOX = 'https://sandbox-api.snapwallet.io',
}

declare global {
  var _ENV: {
    WIDGET_URL: string
    INIT_API_BASE_URL: string
  }
}

interface QROptions {
  element: HTMLElement
  foregroundColor?: string
  backgroundColor?: string
  pixelSize?: number
}

interface IWallet {
  asset: string
  address?: string
}

interface IProduct {
  imageURL?: string
  videoURL?: string
  destinationAmount: number
  destinationTicker: string
  destinationAddress: string
  title: string
}

interface IConfig {
  onMessage?: (e: any) => any
  wallets: IWallet[]
  appName: string
  payee?: string
  intent?: UserIntent
  focus?: boolean
  sourceAmount?: number
  theme?: { [cssProperty: string]: string }
  product?: IProduct
  defaultDestinationAsset?: string
  displayAmount?: SrcDst
  environment: WidgetEnvironments
  baseURL?: WidgetURLs
}

class Snap {
  IFRAME_ID = '__SNAP_IFRAME'
  events = {
    EXIT: '__SNAP_EXIT',
    SUCCESS: '__SNAP_SUCCESS',
    RESIZE: '__SNAP_RESIZE',
    DEMO_CURRENCY_SELECTED: '__SNAP_DEMO_CURRENCY_SELECTED',
  }
  onMessage = (e: any) => {}
  private originalConfig: IConfig
  wallets: IWallet[] = []
  appName: string = 'Snap Wallet'
  payee: string = ''
  intent: UserIntent = 'buy'
  baseURL?: WidgetURLs
  environment: WidgetEnvironments = WidgetEnvironments.PRODUCTION
  focus: boolean = true
  theme?: { [cssProperty: string]: string }
  sourceAmount?: number
  product?: IProduct
  defaultDestinationAsset?: string
  displayAmount?: SrcDst
  private API: FluxApi

  constructor(args: IConfig) {
    this.setConfig(args)
    this.originalConfig = this.getConfig()
    this.API = this.genAPIClient()
  }

  setConfig = (config: IConfig) => {
    Object.assign(this, { ...this.originalConfig, ...config })
  }

  getConfig = (): IConfig => {
    const baseURL = this.getBaseURL()
    return {
      baseURL,
      wallets: this.wallets || [],
      appName: this.appName,
      intent: this.intent,
      payee: this.payee,
      focus: this.focus,
      theme: this.theme || {},
      product: this.product,
      sourceAmount: this.sourceAmount,
      defaultDestinationAsset: this.defaultDestinationAsset,
      displayAmount: this.displayAmount,
      environment: this.environment,
    }
  }

  configToQueryString = () => {
    return encodeURIComponent(JSON.stringify(this.getConfig()))
  }

  openWeb = (config?: IConfig) => {
    config && this.setConfig(config)

    const iframe = document.createElement('iframe')
    iframe.id = this.IFRAME_ID
    // TODO: toggle URL per env
    iframe.src = this.generateURL()
    iframe.classList.add('snapWallet')
    iframe.frameBorder = '0'
    iframe.style.backgroundColor = 'transparent'
    iframe.style.position = 'fixed'
    iframe.style.top = '0'
    iframe.style.right = '0'
    iframe.style.bottom = '0'
    iframe.style.left = '0'
    iframe.style.width = '100%'
    iframe.style.height = '100%'
    iframe.style.border = 'none !important'
    iframe.style.zIndex = '1000000000'
    iframe.style.boxSizing = 'border-box'
    iframe.allow = 'camera:*;microphone:*;'
    iframe.scrolling = 'no'
    window.addEventListener('message', this.handleMessage, false)
    document.body.appendChild(iframe)
    document.body.classList.add('sw-loaded') // css triggers
    if (window.onSnapWalletOpen) window.onSnapWalletOpen(this.IFRAME_ID, config) // cb
  }

  closeWeb = () => {
    window.removeEventListener('message', () => {}, false)
    const iframe = document.getElementById(this.IFRAME_ID)
    iframe?.remove()
    document.body.classList.remove('sw-loaded') // css triggers
    if (window.onSnapWalletClose) window.onSnapWalletClose(this.IFRAME_ID) // cb
  }

  openNative = () => {
    // TODO: add RN WV launch logic
  }

  createQR = async (qrOpts: QROptions, config?: IConfig) => {
    const text = await this.getShortURL(config)
    QR.render(
      {
        text,
        radius: 0.0, // 0.0 to 0.5
        ecLevel: 'H', // L, M, Q, H
        fill: qrOpts.foregroundColor || '#485460',
        background: qrOpts.backgroundColor, // transparent default
        size: qrOpts.pixelSize || 128, // in pixels
      },
      qrOpts.element
    )
  }

  generateURL = (config?: IConfig) => {
    config && this.setConfig(config)
    const qs = `?ts=${Date.now()}&config=${this.configToQueryString()}`
    return `${this.getBaseURL()}/${qs}#/`
  }

  getShortURL = async (config?: IConfig): Promise<string> => {
    config && this.setConfig(config)
    const res = await this.API.fluxWidgetGetShortUrl(this.getConfig())
    return res.url!
  }

  private handleMessage = (event: any) => {
    try {
      const { data = '{}' } = event
      const msg = JSON.parse(data)
      this.onMessage && this.onMessage(msg)
    } catch (e) {
      console.error('SnapWallet:error', 'unable to parse message:', event, e)
    }
  }

  private genAPIClient = (): FluxApi => {
    const apiBaseURL = this.getAPIBaseURL()
    const config = createConfiguration({
      baseServer: new ServerConfiguration(apiBaseURL, {}),
    })
    return new FluxApi(config)
  }

  /**
   * Get the widget URL based on the environment
   * @returns Widget URL
   */
  private getBaseURL = (): WidgetURLs => {
    // Allow dev to override hardcoded 'sandbox' env for web
    if (_ENV.WIDGET_URL) return _ENV.WIDGET_URL as WidgetURLs
    if (this.environment === WidgetEnvironments.SANDBOX) {
      return WidgetURLs.SANDBOX
    }
    return WidgetURLs.PRODUCTION
  }

  /**
   * Get the API base URL based on the environment
   * @returns API URL
   */
  private getAPIBaseURL = (): APIBaseURLs => {
    // Allow dev to override hardcoded 'sandbox' env for web
    if (_ENV.INIT_API_BASE_URL) return _ENV.INIT_API_BASE_URL as APIBaseURLs
    if (this.environment === WidgetEnvironments.SANDBOX) {
      return APIBaseURLs.SANDBOX
    }
    return APIBaseURLs.PRODUCTION
  }
}

;(globalThis as any).Snap = Snap

export default Snap

import QR from 'qr-creator'
import { createConfiguration, FluxApi, ServerConfiguration } from 'api-client'

declare global {
  var _ENV: {
    WIDGET_URL: string
    API_BASE_URL: string
  }
}

type UserIntent = 'buy' | 'sell'

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
  intent: UserIntent
  focus: boolean
  sourceAmount?: number
  theme?: { [cssProperty: string]: string }
  product?: IProduct
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
  wallets: IWallet[] = []
  appName: string = 'Snap Wallet'
  intent: UserIntent = 'buy'
  baseURL: string = _ENV.WIDGET_URL
  focus: boolean = true
  theme?: { [cssProperty: string]: string }
  sourceAmount?: number
  product?: IProduct
  private API: FluxApi

  constructor(args: IConfig) {
    this.setConfig(args)
    this.API = this.genAPIClient()
  }

  setConfig = (config: IConfig) => {
    this.onMessage = config.onMessage || this.onMessage
    this.wallets = config.wallets || this.wallets
    this.appName = config.appName || this.appName
    this.intent = config.intent || this.intent
    this.focus = config.focus ?? this.focus
    this.theme = config.theme || this.theme
    this.product = config.product || this.product
    this.sourceAmount = config.sourceAmount || this.sourceAmount
  }

  getConfig = (): IConfig => {
    return {
      wallets: this.wallets,
      appName: this.appName,
      intent: this.intent,
      focus: this.focus,
      theme: this.theme,
      product: this.product,
      sourceAmount: this.sourceAmount,
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
  }

  closeWeb = () => {
    window.removeEventListener('message', () => {}, false)
    const iframe = document.getElementById(this.IFRAME_ID)
    iframe?.remove()
  }

  openNative = () => {
    // TODO: add RN WV launch logic
  }

  createQR = async (qrOpts: QROptions, config?: IConfig) => {
    const text = await this.getShortURL(config)
    QR.render(
      {
        text,
        radius: 0.5, // 0.0 to 0.5
        ecLevel: 'H', // L, M, Q, H
        fill: qrOpts.foregroundColor || '#485460',
        background: qrOpts.backgroundColor, // transparent defautl
        size: qrOpts.pixelSize || 128, // in pixels
      },
      qrOpts.element
    )
  }

  generateURL = (config?: IConfig) => {
    config && this.setConfig(config)
    const qs = `?ts=${Date.now()}&config=${this.configToQueryString()}`
    return `${this.baseURL}/${qs}#/`
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
    const config = createConfiguration({
      baseServer: new ServerConfiguration(_ENV.API_BASE_URL, {}),
    })
    return new FluxApi(config)
  }
}

;(globalThis as any).Snap = Snap

export default Snap

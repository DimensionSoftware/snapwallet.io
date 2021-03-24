type UserIntent = 'buy' | 'sell'

interface IWallet {
  asset: string
  address?: string
}

interface IConfig {
  onMessage?: (e: any) => any
  wallets: IWallet[]
  appName: string
  intent: UserIntent
}

class Snap {
  IFRAME_ID = '__SNAP_IFRAME'
  events = {
    EXIT: '__SNAP_EXIT',
  }
  onMessage = (e: any) => {}
  wallets: IWallet[] = []
  appName: string = 'Snap Wallet'
  intent: UserIntent = 'buy'

  constructor(args: IConfig) {
    this.setConfig(args)
  }

  setConfig = (config: IConfig) => {
    this.onMessage = config.onMessage || this.onMessage
    this.wallets = config.wallets || this.wallets
    this.appName = config.appName || this.appName
    this.intent = config.intent || this.intent
  }

  configToQueryString = () => {
    return encodeURIComponent(
      JSON.stringify({
        wallets: this.wallets,
        appName: this.appName,
        intent: this.intent,
      })
    )
  }

  openWeb = (config?: IConfig) => {
    config && this.setConfig(config)
    const qs = `?config=${this.configToQueryString()}`
    const iframe = document.createElement('iframe')
    iframe.id = this.IFRAME_ID
    // TODO: toggle URL per env
    iframe.src = `http://localhost:5000/${qs}#/`
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

  private handleMessage = (event: any) => {
    const { data = '{}' } = event
    const msg = JSON.parse(data)
    this.onMessage && this.onMessage(msg)
  }
}

;(globalThis as any).Snap = Snap

export default Snap

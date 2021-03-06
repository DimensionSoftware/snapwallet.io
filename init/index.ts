class Flux {
  IFRAME_ID = '__FLUX_IFRAME'
  events = {
    EXIT: '__FLUX_EXIT',
  }
  onExit = (e: any) => {}

  constructor(args: { onExit?: (e: any) => any }) {
    this.onExit = args.onExit || this.onExit
  }

  openWeb = () => {
    const iframe = document.createElement('iframe')
    iframe.id = this.IFRAME_ID
    // TODO: toggle URL per env
    iframe.src = 'http://localhost:5000/#/'
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
    window.addEventListener(
      'message',
      (event) => {
        const { data = '{}' } = event
        const msg = JSON.parse(data)
        this.onExit && this.onExit(msg)
      },
      false
    )
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
}

;(globalThis as any).Flux = Flux

export default Flux

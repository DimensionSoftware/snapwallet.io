class Flux {
  events = {
    EXIT: '__FLUX_EXIT',
  }

  openWeb = () => {
    const iframe = document.createElement('iframe')
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
    document.body.appendChild(iframe)
  }

  openNative = () => {
    // TODO: add RN WV launch logic
  }
}

;(globalThis as any).Flux = Flux

export default Flux

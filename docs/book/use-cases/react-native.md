# React Native

Below is a minimal React Native example demonstrating how to simply launch Snap Wallet within your own app! Swap our bare config out for your custom settings. Both Apple iOS and Google Android platforms share exactly the same API.

```tsx
import React, { useState } from 'react'
import { Button, SafeAreaView, StatusBar, StyleSheet, View } from 'react-native'
import { WebView } from 'react-native-webview'

// Import Snap Wallet and relevant enum(s)
import Snap, { WidgetEnvironments } from '@snapwallet/init'

// Configure the Snap Wallet instance
const snap = new Snap({
  environment: WidgetEnvironments.SANDBOX,
  appName: 'My App',
  wallets: [
    {
      asset: 'btc',
      address: 'ms6k9Mdsbq5ZkoXakJexxjGjpH2PbSQdWK',
    },
  ],
})

// Create a Snap Wallet URI for the WebView component
// This function also accepts runtime configuration
const uri = snap.generateURL()

// Create a WebView message listener for incoming Snap Wallet messages
const onMessage = async (event: any) => {
  try {
    const { data = '{}' } = event.nativeEvent
    const snapWalletMsg: { data: any; event: any } = JSON.parse(data)

    switch (snapWalletMsg.event) {
      case snap.events.EXIT:
        props.onExit && props.onExit(snapWalletMsg)
        break
      case snap.events.SUCCESS:
        console.log('Success', snapWalletMsg)
        break
      default:
        break
    }
  } catch (e) {
    console.error('Error processing Snap Wallet message', e)
  }
}

// The Snap Wallet WebView component
const SnapWallet = (props: { onExit?: (msg?: any) => any }) => (
  <WebView style={styles.webView} source={{ uri }} onMessage={onMessage} />
)

// Create a button for opening the Snap Wallet WebView
const SnapWalletButton = (props: { onPress: (e: any) => void }) => {
  return <Button onPress={props.onPress} title="Buy Cryptocurrency" />
}

const App = () => {
  const [isSnapWalletVisible, setSnapWalletIsVisible] = useState(false)
  return (
    <>
      <StatusBar barStyle="dark-content" />
      <SafeAreaView style={styles.safeView}>
        <View style={styles.container}>
          {isSnapWalletVisible ? (
            <SnapWallet onExit={(_msg) => setSnapWalletIsVisible(false)} />
          ) : (
            <View style={styles.buttonContainer}>
              <SnapWalletButton
                onPress={(_e: any) => setSnapWalletIsVisible(true)}
              />
            </View>
          )}
        </View>
      </SafeAreaView>
    </>
  )
}

const styles = StyleSheet.create({
  safeView: {
    backgroundColor: '#f8f8f8',
  },
  container: {
    height: '100%',
    width: '100%',
    display: 'flex',
    backgroundColor: '#f8f8f8',
  },
  webView: {
    height: '100%',
    width: '100%',
    backgroundColor: '#f8f8f8',
  },
  buttonContainer: {
    flex: 1,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
  },
})

export default App
```

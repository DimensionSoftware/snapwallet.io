import React, {useState} from 'react';
import {Button, SafeAreaView, StatusBar, StyleSheet, View} from 'react-native';
import {WebView} from 'react-native-webview';
import Snap, {WidgetEnvironments} from '@snapwallet/init';

// Configure the Snap Wallet instance
const snap = new Snap({
  environment: WidgetEnvironments.SANDBOX,
  appName: 'Some App',
  wallets: [],
});

// Create a Snap Wallet URI
const uri = snap.generateURL();

const SnapWallet = (props: {onExit?: (msg?: any) => any}) => (
  <WebView
    style={styles.webView}
    source={{uri}}
    onMessage={(event: any) => {
      try {
        const {data = '{}'} = event.nativeEvent;
        const snapWalletMsg: {data: any; event: any} = JSON.parse(data);

        switch (snapWalletMsg.event) {
          case snap.events.EXIT:
            props.onExit && props.onExit(snapWalletMsg);
            break;
          case snap.events.SUCCESS:
            console.log('Success', snapWalletMsg);
            break;
          default:
            break;
        }
      } catch (e) {
        console.error('Error processing Snap Wallet message', e);
      }
    }}
  />
);

const SnapWalletButton = (props: {onPress: (e: any) => void}) => {
  return <Button onPress={props.onPress} title="Buy Cryptocurrency" />;
};

const App = () => {
  const [isSnapWalletVisible, setSnapWalletIsVisible] = useState(false);
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
  );
};

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
});

export default App;

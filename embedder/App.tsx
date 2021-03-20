import React, {useState} from 'react';
import {Button, SafeAreaView, StatusBar, StyleSheet, View} from 'react-native';
import {WebView} from 'react-native-webview';

enum ChildMessages {
  EXIT = '__SNAP_EXIT',
}

const Flux = (props: {onExit?: (msg?: any) => any}) => (
  <WebView
    style={styles.webView}
    source={{uri: 'http://localhost:5000/#/'}}
    onMessage={(event) => {
      const {data = '{}'} = event.nativeEvent;
      const fluxMsg = JSON.parse(data);

      switch (fluxMsg.event) {
        case ChildMessages.EXIT:
          props.onExit && props.onExit(fluxMsg);
          break;
        default:
          console.warn('Unknown Flux msg', fluxMsg);
          break;
      }
    }}
  />
);

const FluxButton = (props: {onPress: (e: any) => void}) => {
  return <Button onPress={props.onPress} title="Buy Cryptocurrency" />;
};

const App = () => {
  const [isFluxVisible, setFluxIsVisible] = useState(false);
  return (
    <>
      <StatusBar barStyle="dark-content" />
      <SafeAreaView style={styles.safeView}>
        <View style={styles.container}>
          {isFluxVisible ? (
            <Flux onExit={(_msg) => setFluxIsVisible(false)} />
          ) : (
            <View style={styles.buttonContainer}>
              <FluxButton onPress={(_e: any) => setFluxIsVisible(true)} />
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

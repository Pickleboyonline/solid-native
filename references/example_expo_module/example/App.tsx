import { StyleSheet, Text, View } from 'react-native';

import * as TestExpoModule from 'test-expo-module';

export default function App() {
  return (
    <View style={styles.container}>
      <Text>{TestExpoModule.hello()}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});

import { View, Text } from "solid-native/core";


export function DukatapeApp() {
    return (
        <View style={{ flex: 1 }}>
            <View style={{ flex: 1, backgroundColor: '#328c59' }} />
            <View style={{ flex: 1, backgroundColor: '#4287f5', padding: 20 }} >
                {/* <View style={{
                    position: 'absolute',
                    width: 100,
                    height: 100,
                    top: 20,
                    right: 20,
                    backgroundColor: '#328c59'
                }} /> */}
                <Text style={{
                    textTransform: 'capitalize'
                }}>
                    Hello my name is Imran
                </Text>

                {/* <View style={{
                    width: 100,
                    height: 100,
                    backgroundColor: '#328c59'
                }} /> */}
            </View>
        </View>
    )
}
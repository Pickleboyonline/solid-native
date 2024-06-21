import { View } from "solid-native/core";


export function DukatapeApp() {
    return (
        <View style={{ flex: 1 }}>
            <View backgroundColor='#328c59' style={{ flex: 1 }} />
            <View backgroundColor='#4287f5' style={{ flex: 1 }} >
                <View backgroundColor='#328c59' style={{
                    position: 'absolute',
                    width: 100,
                    height: 100,
                    top: 20,
                    right: 20
                }} />
            </View>
        </View>
    )
}
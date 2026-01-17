import { View, Text, Image } from "solid-native";

export function App() {
  return (
    <View style={{ flex: 1, alignItems: "center", justifyContent: "center" }}>
      <Text style={{ fontSize: 24, fontWeight: "bold", color: "#333" }}>
        Hello Solid Native!
      </Text>
      <View style={{ height: 20 }} />
      <Image
        source={{ uri: "https://picsum.photos/200" }}
        style={{ width: 200, height: 200 }}
        resizeMode="cover"
      />
      <View style={{ height: 20 }} />
      <Text style={{ fontSize: 14, color: "#666" }}>
        Built with Rust + SolidJS + SwiftUI
      </Text>
    </View>
  );
}

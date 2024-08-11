import { View, Text } from "solid-native/core";

export function DuktapeApp() {
  return (
    <View
      style={{
        flex: 1,
        backgroundColor: "#328c59",
        paddingTop: 100,
        flexDirection: "row",
      }}
    >
      Hello my name is Imran
      <View
        style={{
          backgroundColor: "#9d29ba",
          height: 100,
          width: 100,
        }}
      />
    </View>
  );
}

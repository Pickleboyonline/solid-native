import { createSignal, onMount } from "solid-js";
import { print, View, Text, Button } from "solid-native/core";
import { createEffect } from "solid-js";

export function App() {
  const [count, setCount] = createSignal(0);

  setInterval(() => setCount(count() + 1), 1000);

  const [isBold, setIsBold] = createSignal<"bold" | "normal">("bold");
  const [bool, setBool] = createSignal(false);

  onMount(() => {
    print("App Mounted!");
  });

  const flipValues = () => {
    setBool(!bool());
    if (isBold() === "bold") {
      setIsBold("normal");
    } else {
      setIsBold("bold");
    }
  };

  createEffect(() => {
    print("bool has changed!" + bool());
  });

  return (
    <View>
      <Text>
        {"Bruh " + count() + " "}
        <Text
          style={{
            fontWeight: isBold(),
            textDecorationLine: "underline",
          }}
        >
          Hello World:
        </Text>
      </Text>
      <Text
        style={{
          color: bool() ? "#4287f5" : "#8cb512",
          fontWeight: "bold",
          fontStyle: "italic",
        }}
      >
        HELLO!
      </Text>
      <Button title="Flip Values" onPress={flipValues} />
    </View>
  );
}

import { createSignal, onMount } from "solid-js";
import { log, View, Text, Button } from "solid-native/core";
import { createEffect } from "solid-js";

export function App() {
  const [count, setCount] = createSignal(0);

  setInterval(() => setCount(count() + 1), 1000);

  const [isBold, setIsBold] = createSignal<"bold" | "normal">("bold");
  const [bool, setBool] = createSignal(false);

  onMount(() => {
    log("App Mounted!");
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
    log("bool has changed!" + bool());
  });

  return (
    <View style={{ flex: 1 }}>
      <View style={{
        flex: 1,
        backgroundColor: '#32a852',
        flexDirection: 'row',
        // marginTop: 100
      }}>
        <Text>
          {"Counter: " + count() + " "}
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
      <View style={{
        flex: 1,
        backgroundColor: '#7a2c91',
        paddingTop: 100
      }}>
        I'm purple
      </View>
      <View style={{
        flex: 1,
        backgroundColor: '#4287f5',
        paddingTop: 100
      }}>
        I'm blue
      </View>
    </View>
  );
}

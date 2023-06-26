import { Button, Text, VStack } from "solid-native/core";
import { createSignal } from "solid-js";
import { FontWeight } from "../../core/views/mod.ts";

export function TextExample() {
  const [isBold, setIsBold] = createSignal(false);

  return (
    <VStack>
      <Text
        // TODO: Make shadow component structure
        text="Hi Imran!"
        // fontWeight={FontWeight[isBold() ? "bold" : "regular"]}
      />
      <Button
        title="Toggle Bold!"
        onPress={() => {
          setIsBold(!isBold());
        }}
      />
    </VStack>
  );
}

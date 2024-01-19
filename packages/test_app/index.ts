import { print, getNativeModule } from "solid-native/core";
// import { App } from "./App.tsx";
// // deno-lint-ignore ban-ts-comment
// // @ts-ignore
// render(App, SolidNativeCore.getRootElement());

type Renderer = {
  print: (str: string) => void;
  getRootView: () => string;
  printCB: (str: () => string) => void;
};

const renderer = getNativeModule<Renderer>("SNRender");

renderer.print("Hello!" + renderer.getRootView());

renderer.printCB(() => "hi");

// const s = ;

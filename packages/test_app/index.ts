import { SolidNativeRenderer, print, render, Button } from "solid-native/core";
import { App } from "./App.tsx";
print("Hello World from S!");
// const rootView = SolidNativeRenderer.getRootView();

// print("RootView: " + rootView);

// const text = SolidNativeRenderer.createNodeByName("sn_text");

// SolidNativeRenderer.insertBefore(rootView, text);

// SolidNativeRenderer.setProp(text, "text", "Hello World!");

// const button = SolidNativeRenderer.createNodeByName("sn_button");
// SolidNativeRenderer.setProp(button, "title", "Click Me!");

// SolidNativeRenderer.insertBefore(rootView, button);

render(App);

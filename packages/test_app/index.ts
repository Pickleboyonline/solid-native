import { SolidNativeRenderer, print } from "solid-native/core";

print("Hello World from JS!");

const rootView = SolidNativeRenderer.getRootView();

print("RootView: " + rootView);

const text = SolidNativeRenderer.createNodeByName("sn_text");

SolidNativeRenderer.insertBefore(rootView, text);

SolidNativeRenderer.setProp(text, "text", "Hello World!");

const button = SolidNativeRenderer.createNodeByName("sn_button");
SolidNativeRenderer.setProp(button, "title", "Click Me!");

SolidNativeRenderer.insertBefore(rootView, button);

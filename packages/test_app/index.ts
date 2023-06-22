import { SolidNativeCore, getNativeModule, render } from "../core/mod.ts";
import { App } from "./App.tsx";

const print = getNativeModule<(str: string) => void>("print");

render(App, SolidNativeCore.getRootElement());
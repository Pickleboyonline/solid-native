import { SolidNativeCore, render } from "../core/mod.ts";
import { App } from "./App.tsx";
// deno-lint-ignore ban-ts-comment
// @ts-ignore
render(App, SolidNativeCore.getRootElement());

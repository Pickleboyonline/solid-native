import * as esbuild from "esbuild";
import * as esbuildDenoLoader from "@luca/esbuild-deno-loader";
import { solidPlugin } from "esbuild-plugin-solid";
import { resolve } from "https://deno.land/std@0.208.0/path/mod.ts";

const outfile = resolve(Deno.cwd(), "dist/bundle.js");

// Ensure dist directory exists
try {
  await Deno.mkdir(resolve(Deno.cwd(), "dist"), { recursive: true });
} catch {
  // Directory already exists
}

const esbuildPlugins = (() => {
  const configPath = import.meta
    .resolve("./../../../deno.json")
    .replace("file://", "");

  // @ts-ignore: Versions are OK, think its some Deno Global.URL type mismatch
  return [
    solidPlugin({
      solid: {
        moduleName: "solid-native-renderer",
        generate: "universal",
      },
    }),
    ...esbuildDenoLoader.denoPlugins({
      configPath,
    }),
  ] as esbuild.Plugin[];
})();

const result = await esbuild.build({
  plugins: esbuildPlugins,
  entryPoints: ["src/index.ts"],
  bundle: true,
  outfile,
  target: "ES2020",
  sourcemap: true,
});

console.log("Bundle created:", outfile);
if (result.errors.length > 0) {
  console.error("Errors:", result.errors);
}

esbuild.stop();

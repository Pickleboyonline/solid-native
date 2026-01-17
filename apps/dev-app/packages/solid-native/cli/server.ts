import { Application, Router } from "@oak/oak";
import * as esbuild from "esbuild";
import * as esbuildDenoLoader from "@luca/esbuild-deno-loader";
import { solidPlugin } from "esbuild-plugin-solid";

const app = new Application();
const router = new Router();

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

router
  .get("/", async (ctx) => {
    console.log("Request: /");

    const result = await esbuild.build({
      plugins: esbuildPlugins,
      entryPoints: ["src/index.ts"],
      bundle: true,
      write: false,
      target: "ES2020",
      outdir: "out",
    });

    const contents = result.outputFiles?.[0].text;
    ctx.response.body = contents;
  })
  .get("/bundle.js", async (ctx) => {
    console.log("Request: /bundle.js");

    const result = await esbuild.build({
      plugins: esbuildPlugins,
      entryPoints: ["src/index.ts"],
      bundle: true,
      write: false,
      target: "ES2020",
      outdir: "out",
    });

    const contents = result.outputFiles?.[0].text;
    ctx.response.headers.set("Content-Type", "application/javascript");
    ctx.response.body = contents;
  })
  .get("/source", async (ctx) => {
    console.log("Request: /source");

    const result = await esbuild.build({
      plugins: esbuildPlugins,
      entryPoints: ["src/index.ts"],
      bundle: true,
      write: false,
      sourcemap: true,
      target: "ES2020",
      outdir: "out",
    });

    const contents = result.outputFiles[0].text;
    ctx.response.body = contents;
  })
  .get("/health", (ctx) => {
    ctx.response.body = "OK";
  });

app.use(router.routes());
app.use(router.allowedMethods());

const port = 8080;
console.log(`Starting dev server on http://localhost:${port}`);
app.listen({ port });

import { Application, Router } from "@oak/oak";
import * as esbuild from "esbuid";
import * as esbuildDenoLoader from "@luca/esbuild-deno-loader";
import { solidPlugin } from "esbuild-plugin-solid";

const app = new Application();

const router = new Router();

const esbuildPlugins = (() => {
  const configPath = import.meta
    .resolve("./../../deno.json")
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
    console.log("Request!");

    const result = await esbuild.build({
      // plugins: esbuildPlugins,
      plugins: esbuildPlugins,
      entryPoints: ["packages/test_app/index.ts"],
      bundle: true,
      write: false,
      // Actually converts to es5
      target: "ES2020",
      outdir: "out",
    });

    const contents = result.outputFiles?.[0].text;
    ctx.response.body = contents;
  })
  .get("/source", async (ctx) => {
    console.log("Request Source!");

    const result = await esbuild.build({
      plugins: esbuildPlugins,
      entryPoints: ["packages/test_app/index.ts"],
      bundle: true,
      write: false,
      sourcemap: true,
      target: "ES2020",
      outdir: "out",
    });

    const contents = result.outputFiles[0].text;
    ctx.response.body = contents;
  });

app.use(router.routes());
app.use(router.allowedMethods());

const port = 8080;
console.log(`listening on http://localhost:${port}`);
app.listen({ port });

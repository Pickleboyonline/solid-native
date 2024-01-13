import { Application, Router } from "https://deno.land/x/oak@v12.5.0/mod.ts";
import * as esbuild from "https://deno.land/x/esbuild@v0.18.6/mod.js";

import * as esbuildDenoLoader from "https://deno.land/x/esbuild_deno_loader@0.8.1/mod.ts";
import { solidPlugin } from "npm:esbuild-plugin-solid@0.5.0";

const app = new Application();

const router = new Router();
router
  .get("/", async (ctx) => {
    const configPath = import.meta
      .resolve("./../../deno.json")
      .replace("file://", "");

    console.log();

    console.log("Request!");

    const result = await esbuild.build({
      plugins: [
        solidPlugin({
          solid: {
            moduleName: "solid-native-renderer",
            generate: "universal",
          },
        }),
        ...esbuildDenoLoader.denoPlugins({
          configPath,
        }),
      ],
      entryPoints: ["packages/test_app/index.ts"],
      bundle: true,
      write: false,
      target: "ios16",
      outdir: "out",
    });

    const contents = result.outputFiles[0].text;
    ctx.response.body = contents;
  })
  .get("/source", async (ctx) => {
    const configPath = import.meta
      .resolve("./../../deno.json")
      .replace("file://", "");

    console.log("Request!");

    const result = await esbuild.build({
      plugins: [
        solidPlugin({
          solid: {
            moduleName: "solid-native-renderer",
            generate: "universal",
          },
        }),
        ...esbuildDenoLoader.denoPlugins({
          configPath,
        }),
      ],
      entryPoints: ["packages/test_app/index.ts"],
      bundle: true,
      write: false,
      sourcemap: true,
      target: "ios16",
      outdir: "out",
    });

    const contents = result.outputFiles[0].text;
    ctx.response.body = contents;
  });

app.use(router.routes());
app.use(router.allowedMethods());

app.listen({ port: 8080 });

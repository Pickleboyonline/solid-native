import { Application, Router } from "@oak/oak";
import * as esbuild from "esbuid";
import * as esbuildDenoLoader from "@luca/esbuild-deno-loader";
import { solidPlugin } from "esbuild-plugin-solid";

import { transform } from "npm:@swc/core";
import { Buffer } from "node:buffer";

const app = new Application();

const router = new Router();

/**
 * Convert it to ES5 for duktape. Source maps not really working properly
 * Code on how to do it better: https://github.com/noyobo/esbuild-plugin-es5/blob/main/src/index.ts
 * @returns
 */
const swcPlugin = (): esbuild.Plugin => {
  return {
    name: "swc-plugin",
    setup(build) {
      build.onEnd(async (result) => {
        if (result.errors.length) return;

        for (const file of result.outputFiles || []) {
          if (file.path.endsWith(".js")) {
            try {
              const transformed = await transform(file.text, {
                sourceMaps: (build.initialOptions.sourcemap) ? "inline" : false,
                // minify: true,
              });
              file.contents = Buffer.from(transformed.code, "utf-8");
            } catch (error) {
              console.error("SWC transform error:", error);
              result.errors.push(
                { text: (error as Error).message } as esbuild.Message,
              );
            }
          }
        }
      });
    },
  };
};

const swcPluginWithSourceMaps = (): esbuild.Plugin => {
  return {
    name: "swc-plugin",
    setup(build) {
      build.onEnd(async (result) => {
        if (result.errors.length) return;

        for (const file of result.outputFiles || []) {
          if (file.path.endsWith(".js")) {
            try {
              const transformed = await transform(file.text);
              file.contents = Buffer.from(transformed.code, "utf-8");
            } catch (error) {
              console.error("SWC transform error:", error);
              result.errors.push(
                { text: (error as Error).message } as esbuild.Message,
              );
            }
          }
        }
      });
    },
  };
};

const esbuildPlugins = (() => {
  const configPath = import.meta
    .resolve("./../../deno.json")
    .replace("file://", "");

  // @ts-ignore: Versions are OK, think its some Deno Global.URL type mismatch
  return [
    swcPlugin(),
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
      target: "es6",
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
      target: "ios16",
      outdir: "out",
    });

    const contents = result.outputFiles[0].text;
    ctx.response.body = contents;
  });

app.use(router.routes());
app.use(router.allowedMethods());

app.listen({ port: 8080 });

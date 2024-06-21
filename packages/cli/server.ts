import { Application, Router } from "@oak/oak";
import * as esbuild from "esbuid";
import * as esbuildDenoLoader from "@luca/esbuild-deno-loader";
import { solidPlugin } from "esbuild-plugin-solid";

import { Options, Output } from "npm:@swc/types";
import { transform, transformFile } from "npm:@swc/core";
import { Buffer } from "node:buffer";

const app = new Application();

const router = new Router();

/**
 * Convert it to ES5 for duktape. Source maps not really working properly
 * Code on how to do it better: https://github.com/noyobo/esbuild-plugin-es5/blob/main/src/index.ts
 * TODO: Get sourcemaps working
 * @returns
 */
const swcPlugin = (): esbuild.Plugin => {
  return {
    name: "swc-plugin",
    setup(build) {
      // build.onLoad({ filter: /\.([tj]sx?|mjs)$/ }, async (args) => {
      //   const opts: Options = {
      //     env: {
      //       mode: "usage",
      //       coreJs: "3.22",
      //     },
      //   };
      //   const results = (await transformFile(args.path, opts)) as Output;

      //   return {
      //     contents: results.code.replaceAll(
      //       `import "core-js`,
      //       `import "npm:core-js`,
      //     ),
      //     loader: "js",
      //   };
      // });

      build.onEnd(async (result) => {
        if (result.errors.length) return;

        for (const file of result.outputFiles || []) {
          if (file.path.endsWith(".js")) {
            try {
              // Deno.writeTextFile("data.js", file.text);
              const transformed = await transform(file.text, {
                sourceMaps: (build.initialOptions.sourcemap) ? "inline" : false,
                jsc: {
                  target: "es5",
                },
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
      target: "es5",
      outdir: "out",
    });

    const contents = result.outputFiles[0].text;
    ctx.response.body = contents;
  });

app.use(router.routes());
app.use(router.allowedMethods());

app.listen({ port: 8080 });

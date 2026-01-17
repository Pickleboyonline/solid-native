import * as esbuild from "esbuild";
import { resolve } from "https://deno.land/std@0.208.0/path/mod.ts";

const entryPoint = resolve(Deno.cwd(), "src/index.ts");
const outfile = resolve(Deno.cwd(), "dist/bundle.js");

// Ensure dist directory exists
try {
  await Deno.mkdir(resolve(Deno.cwd(), "dist"), { recursive: true });
} catch {
  // Directory already exists
}

const result = await esbuild.build({
  entryPoints: [entryPoint],
  bundle: true,
  outfile,
  format: "esm",
  platform: "neutral",
  target: "es2020",
  jsx: "automatic",
  jsxImportSource: "solid-js",
  minify: false,
  sourcemap: true,
  external: [],
  define: {
    "process.env.NODE_ENV": '"development"',
  },
});

console.log("Bundle created:", outfile);
if (result.errors.length > 0) {
  console.error("Errors:", result.errors);
}

esbuild.stop();

import * as esbuild from "esbuild";
import { resolve } from "https://deno.land/std@0.208.0/path/mod.ts";

const PORT = 8080;
const HOST = "0.0.0.0";

// Build the bundle
async function buildBundle(): Promise<string> {
  const entryPoint = resolve(Deno.cwd(), "src/index.ts");

  const result = await esbuild.build({
    entryPoints: [entryPoint],
    bundle: true,
    write: false,
    format: "esm",
    platform: "neutral",
    target: "es2020",
    jsx: "automatic",
    jsxImportSource: "solid-js",
    minify: false,
    define: {
      "process.env.NODE_ENV": '"development"',
    },
  });

  if (result.errors.length > 0) {
    console.error("Build errors:", result.errors);
    throw new Error("Build failed");
  }

  return result.outputFiles[0].text;
}

// Start the server
console.log(`Starting dev server on http://${HOST}:${PORT}`);

let cachedBundle: string | null = null;

Deno.serve({ port: PORT, hostname: HOST }, async (request) => {
  const url = new URL(request.url);
  const path = url.pathname;

  console.log(`${request.method} ${path}`);

  // CORS headers for development
  const headers = {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET, POST, OPTIONS",
    "Access-Control-Allow-Headers": "Content-Type",
  };

  if (request.method === "OPTIONS") {
    return new Response(null, { headers });
  }

  if (path === "/" || path === "/bundle.js") {
    try {
      // Always rebuild in dev mode for now
      // TODO: Add file watching for automatic rebuilds
      console.log("Building bundle...");
      cachedBundle = await buildBundle();
      console.log("Bundle built successfully");

      return new Response(cachedBundle, {
        headers: {
          ...headers,
          "Content-Type": "application/javascript",
        },
      });
    } catch (error) {
      console.error("Build error:", error);
      return new Response(`Build error: ${error}`, {
        status: 500,
        headers,
      });
    }
  }

  if (path === "/health") {
    return new Response("OK", { headers });
  }

  return new Response("Not found", { status: 404, headers });
});

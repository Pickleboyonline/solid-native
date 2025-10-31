#!/usr/bin/env -S deno run --allow-all

/**
 * Start CLI
 * Runs the built application
 */

import { logger } from "../utils/logger.ts";

async function start() {
  logger.info("Starting Solid Native application...");

  // TODO: Implement actual start logic
  // This would:
  // 1. Load built bundle
  // 2. Initialize native runtime
  // 3. Start app on device/simulator
  // 4. Stream logs to console

  logger.info("App started! (placeholder)");
}

if (import.meta.main) {
  await start();
}

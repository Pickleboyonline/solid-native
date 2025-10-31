#!/usr/bin/env -S deno run --allow-all

/**
 * Doctor CLI
 * Diagnoses environment and dependency issues
 */

import { logger } from "../utils/logger.ts";

async function checkEnvironment() {
  logger.info("🔍 Checking Solid Native environment...\n");

  const checks = [
    { name: "Deno version", check: () => checkDenoVersion() },
    { name: "Configuration file", check: () => checkConfig() },
    { name: "Dependencies", check: () => checkDependencies() },
    { name: "AI configuration", check: () => checkAI() },
  ];

  for (const { name, check } of checks) {
    try {
      await check();
      logger.info(`✅ ${name}`);
    } catch (error) {
      logger.error(
        `❌ ${name}: ${error instanceof Error ? error.message : String(error)}`,
      );
    }
  }

  logger.info("\n✨ Environment check complete!");
}

async function checkDenoVersion() {
  const version = Deno.version.deno;
  if (!version) {
    throw new Error("Deno not found");
  }
  // Deno is installed
}

async function checkConfig() {
  const configPath = "./solid-native.config.ts";
  try {
    await Deno.stat(configPath);
  } catch {
    throw new Error("solid-native.config.ts not found");
  }
}

async function checkDependencies() {
  const denoJson = "./deno.json";
  try {
    await Deno.stat(denoJson);
  } catch {
    throw new Error("deno.json not found");
  }
}

async function checkAI() {
  const hasAnthropicKey = !!Deno.env.get("ANTHROPIC_API_KEY");
  const hasOpenAIKey = !!Deno.env.get("OPENAI_API_KEY");

  if (!hasAnthropicKey && !hasOpenAIKey) {
    throw new Error("No AI API keys configured");
  }
}

if (import.meta.main) {
  await checkEnvironment();
}

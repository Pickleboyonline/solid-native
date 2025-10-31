#!/usr/bin/env -S deno run --allow-read --allow-write --allow-env

/**
 * AI CLI - Natural language interface for development tasks
 * Allows developers to use natural language commands for common tasks
 */

import { Command } from "cliffy";
import { createAIClient } from "../ai/client.ts";
import { AIGenerator } from "../ai/generator.ts";

const generateCommand = new Command()
  .name("generate")
  .alias("gen")
  .description("Generate code from a description")
  .arguments("<type:string> <description:string>")
  .action(async (options, type, description) => {
    const client = createAIClient({
      enabled: true,
      provider: "anthropic",
      model: "claude-sonnet-4-5-20250929",
      features: {
        componentGeneration: true,
        codeCompletion: false,
        typeGeneration: false,
        codeReview: false,
        naturalLanguageCommands: true,
        docGeneration: false,
      },
      context: {
        includeProjectStructure: false,
        includeGitHistory: false,
        includeDependencies: false,
        maxContextSize: 50000,
      },
      promptsDir: "./src/ai/prompts",
      schemasDir: "./src/ai/schemas",
    });

    const generator = new AIGenerator(client);

    switch (type) {
      case "component":
        console.log("🤖 Generating component...\n");
        console.log(await generator.generateComponent(description));
        break;
      default:
        console.error(`Unknown type: ${type}`);
        Deno.exit(1);
    }
  });

const reviewCommand = new Command()
  .name("review")
  .description("Review code and get AI feedback")
  .arguments("<file:string>")
  .action(async (options, file) => {
    const code = await Deno.readTextFile(file);

    const client = createAIClient({
      enabled: true,
      provider: "anthropic",
      model: "claude-sonnet-4-5-20250929",
      features: {
        componentGeneration: false,
        codeCompletion: false,
        typeGeneration: false,
        codeReview: true,
        naturalLanguageCommands: true,
        docGeneration: false,
      },
      context: {
        includeProjectStructure: false,
        includeGitHistory: false,
        includeDependencies: false,
        maxContextSize: 50000,
      },
      promptsDir: "./src/ai/prompts",
      schemasDir: "./src/ai/schemas",
    });

    const generator = new AIGenerator(client);

    console.log("🤖 Reviewing code...\n");
    console.log(await generator.reviewCode(code));
  });

const command = new Command()
  .name("ai")
  .version("0.1.0")
  .description("AI-powered development assistant")
  .command("generate", generateCommand)
  .command("review", reviewCommand);

if (import.meta.main) {
  await command.parse(Deno.args);
}

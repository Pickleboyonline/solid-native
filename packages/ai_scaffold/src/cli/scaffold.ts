#!/usr/bin/env -S deno run --allow-read --allow-write --allow-env

/**
 * Scaffold CLI - AI-powered project scaffolding
 * Natural language interface for generating components and code
 */

import { Command } from "cliffy";
import { createAIClient } from "../ai/client.ts";
import { AIGenerator } from "../ai/generator.ts";

const command = new Command()
  .name("scaffold")
  .version("0.1.0")
  .description("AI-powered scaffolding for Solid Native")
  .option("-p, --provider <provider:string>", "AI provider (anthropic, openai)", {
    default: "anthropic",
  })
  .option("-m, --model <model:string>", "Model to use", {
    default: "claude-sonnet-4-5-20250929",
  })
  .arguments("<description:string>")
  .action(async (options, description) => {
    console.log("🤖 Generating component with AI...\n");

    try {
      const client = createAIClient({
        enabled: true,
        provider: options.provider as "anthropic" | "openai",
        model: options.model,
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
      const componentCode = await generator.generateComponent(description);

      console.log("✅ Generated component:\n");
      console.log(componentCode);
      console.log("\n💡 Tip: Save this to a .tsx file in your components directory");
    } catch (error) {
      console.error("❌ Error:", error instanceof Error ? error.message : String(error));
      Deno.exit(1);
    }
  });

if (import.meta.main) {
  await command.parse(Deno.args);
}

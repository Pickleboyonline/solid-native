/**
 * AI-powered code generator
 * Uses LLM to generate components, types, and documentation
 */

import type { AIClient } from "./client.ts";
import {
  type Component,
  ComponentSchema,
  generateComponentFile,
} from "./schemas/component.schema.ts";
import {
  COMPONENT_GENERATOR_SYSTEM_PROMPT,
  COMPONENT_GENERATOR_USER_TEMPLATE,
} from "./prompts/component-generator.ts";

export class AIGenerator {
  constructor(private client: AIClient) {}

  /**
   * Generate a Solid component from a natural language description
   */
  async generateComponent(description: string): Promise<string> {
    const response = await this.client.generate({
      messages: [
        { role: "system", content: COMPONENT_GENERATOR_SYSTEM_PROMPT },
        {
          role: "user",
          content: COMPONENT_GENERATOR_USER_TEMPLATE(description),
        },
      ],
      maxTokens: 4096,
      temperature: 0.7,
    });

    // Extract JSON from response (handle markdown code blocks)
    const jsonMatch = response.match(/```json\n([\s\S]*?)\n```/) ||
      response.match(/```\n([\s\S]*?)\n```/) ||
      [null, response];

    const jsonStr = jsonMatch[1] || response;

    try {
      const parsed = JSON.parse(jsonStr);
      const component = ComponentSchema.parse(parsed);
      return generateComponentFile(component);
    } catch (error) {
      throw new Error(
        `Failed to parse AI response: ${error instanceof Error ? error.message : String(error)}`,
      );
    }
  }

  /**
   * Generate TypeScript types from code
   */
  async generateTypes(code: string): Promise<string> {
    const response = await this.client.generate({
      messages: [
        {
          role: "system",
          content:
            "You are a TypeScript expert. Generate type definitions for the provided code.",
        },
        {
          role: "user",
          content: `Generate TypeScript type definitions for this code:\n\n${code}`,
        },
      ],
      maxTokens: 2048,
      temperature: 0.3,
    });

    return response.trim();
  }

  /**
   * Generate documentation for code
   */
  async generateDocumentation(code: string): Promise<string> {
    const response = await this.client.generate({
      messages: [
        {
          role: "system",
          content:
            "You are a technical documentation expert. Generate clear, comprehensive documentation.",
        },
        {
          role: "user",
          content: `Generate documentation for this code:\n\n${code}`,
        },
      ],
      maxTokens: 2048,
      temperature: 0.5,
    });

    return response.trim();
  }

  /**
   * Review code and provide suggestions
   */
  async reviewCode(code: string): Promise<string> {
    const response = await this.client.generate({
      messages: [
        {
          role: "system",
          content:
            "You are an expert code reviewer. Provide constructive feedback and suggestions.",
        },
        {
          role: "user",
          content: `Review this code and provide feedback:\n\n${code}`,
        },
      ],
      maxTokens: 2048,
      temperature: 0.4,
    });

    return response.trim();
  }
}

/**
 * AI Client - unified interface for LLM providers
 * Supports multiple AI providers with a consistent API
 */

import Anthropic from "@anthropic/sdk";
import OpenAI from "openai";
import type { AIConfig } from "../types/config.ts";

export interface Message {
  role: "system" | "user" | "assistant";
  content: string;
}

export interface GenerateOptions {
  messages: Message[];
  maxTokens?: number;
  temperature?: number;
  stream?: boolean;
}

export interface AIClient {
  generate(options: GenerateOptions): Promise<string>;
  generateStream(options: GenerateOptions): AsyncIterableIterator<string>;
}

export class AnthropicClient implements AIClient {
  private client: Anthropic;
  private model: string;

  constructor(apiKey: string, model: string) {
    this.client = new Anthropic({ apiKey });
    this.model = model;
  }

  async generate(options: GenerateOptions): Promise<string> {
    const systemMessages = options.messages.filter((m) => m.role === "system");
    const conversationMessages = options.messages.filter((m) =>
      m.role !== "system"
    );

    const response = await this.client.messages.create({
      model: this.model,
      max_tokens: options.maxTokens ?? 4096,
      temperature: options.temperature ?? 1.0,
      system: systemMessages.map((m) => m.content).join("\n\n"),
      messages: conversationMessages.map((m) => ({
        role: m.role as "user" | "assistant",
        content: m.content,
      })),
    });

    return response.content[0].type === "text"
      ? response.content[0].text
      : "";
  }

  async *generateStream(
    options: GenerateOptions,
  ): AsyncIterableIterator<string> {
    const systemMessages = options.messages.filter((m) => m.role === "system");
    const conversationMessages = options.messages.filter((m) =>
      m.role !== "system"
    );

    const stream = await this.client.messages.create({
      model: this.model,
      max_tokens: options.maxTokens ?? 4096,
      temperature: options.temperature ?? 1.0,
      system: systemMessages.map((m) => m.content).join("\n\n"),
      messages: conversationMessages.map((m) => ({
        role: m.role as "user" | "assistant",
        content: m.content,
      })),
      stream: true,
    });

    for await (const event of stream) {
      if (
        event.type === "content_block_delta" &&
        event.delta.type === "text_delta"
      ) {
        yield event.delta.text;
      }
    }
  }
}

export class OpenAIClient implements AIClient {
  private client: OpenAI;
  private model: string;

  constructor(apiKey: string, model: string) {
    this.client = new OpenAI({ apiKey });
    this.model = model;
  }

  async generate(options: GenerateOptions): Promise<string> {
    const response = await this.client.chat.completions.create({
      model: this.model,
      max_tokens: options.maxTokens ?? 4096,
      temperature: options.temperature ?? 1.0,
      messages: options.messages.map((m) => ({
        role: m.role,
        content: m.content,
      })),
    });

    return response.choices[0]?.message?.content ?? "";
  }

  async *generateStream(
    options: GenerateOptions,
  ): AsyncIterableIterator<string> {
    const stream = await this.client.chat.completions.create({
      model: this.model,
      max_tokens: options.maxTokens ?? 4096,
      temperature: options.temperature ?? 1.0,
      messages: options.messages.map((m) => ({
        role: m.role,
        content: m.content,
      })),
      stream: true,
    });

    for await (const chunk of stream) {
      const content = chunk.choices[0]?.delta?.content;
      if (content) {
        yield content;
      }
    }
  }
}

/**
 * Factory function to create the appropriate AI client
 */
export function createAIClient(config: AIConfig): AIClient {
  const apiKey = config.apiKey || Deno.env.get("ANTHROPIC_API_KEY") ||
    Deno.env.get("OPENAI_API_KEY");

  if (!apiKey) {
    throw new Error("AI API key not configured");
  }

  switch (config.provider) {
    case "anthropic":
      return new AnthropicClient(apiKey, config.model);
    case "openai":
      return new OpenAIClient(apiKey, config.model);
    case "local":
      throw new Error("Local AI provider not yet implemented");
    default:
      throw new Error(`Unknown AI provider: ${config.provider}`);
  }
}

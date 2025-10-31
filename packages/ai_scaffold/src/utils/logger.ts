/**
 * Logger utility
 * Structured logging with levels and formatting
 */

export type LogLevel = "debug" | "info" | "warn" | "error";

export interface LoggerConfig {
  level: LogLevel;
  prefix?: string;
  timestamp?: boolean;
}

const LOG_LEVELS: Record<LogLevel, number> = {
  debug: 0,
  info: 1,
  warn: 2,
  error: 3,
};

export class Logger {
  private config: Required<LoggerConfig>;

  constructor(config: Partial<LoggerConfig> = {}) {
    this.config = {
      level: "info",
      prefix: "",
      timestamp: true,
      ...config,
    };
  }

  private shouldLog(level: LogLevel): boolean {
    return LOG_LEVELS[level] >= LOG_LEVELS[this.config.level];
  }

  private format(level: LogLevel, ...args: unknown[]): string {
    const parts: string[] = [];

    if (this.config.timestamp) {
      parts.push(`[${new Date().toISOString()}]`);
    }

    parts.push(`[${level.toUpperCase()}]`);

    if (this.config.prefix) {
      parts.push(`[${this.config.prefix}]`);
    }

    const prefix = parts.join(" ");
    const message = args.map((arg) =>
      typeof arg === "object" ? JSON.stringify(arg, null, 2) : String(arg)
    ).join(" ");

    return `${prefix} ${message}`;
  }

  debug(...args: unknown[]): void {
    if (this.shouldLog("debug")) {
      console.debug(this.format("debug", ...args));
    }
  }

  info(...args: unknown[]): void {
    if (this.shouldLog("info")) {
      console.info(this.format("info", ...args));
    }
  }

  warn(...args: unknown[]): void {
    if (this.shouldLog("warn")) {
      console.warn(this.format("warn", ...args));
    }
  }

  error(...args: unknown[]): void {
    if (this.shouldLog("error")) {
      console.error(this.format("error", ...args));
    }
  }
}

// Default logger instance
export const logger = new Logger({
  level: Deno.env.get("LOG_LEVEL") as LogLevel || "info",
  prefix: "SolidNative",
});

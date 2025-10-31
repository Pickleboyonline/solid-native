/**
 * Component schema for AI-powered component generation
 * Defines the structure for generating Solid components via LLM
 */

import { z } from "zod";

/**
 * Style properties schema
 */
const StyleSchema = z.record(z.union([z.string(), z.number()]));

/**
 * Component prop schema
 */
export const ComponentPropSchema = z.object({
  name: z.string(),
  type: z.enum([
    "string",
    "number",
    "boolean",
    "function",
    "object",
    "array",
    "signal",
    "store",
  ]),
  required: z.boolean().default(true),
  description: z.string().optional(),
  defaultValue: z.any().optional(),
});

/**
 * Component schema
 */
export const ComponentSchema = z.object({
  name: z.string().describe("Component name in PascalCase"),
  description: z.string().describe("Brief description of component purpose"),
  props: z.array(ComponentPropSchema).describe("Component props"),
  imports: z.array(z.string()).describe("Required imports"),
  body: z.string().describe("Component JSX body"),
  exports: z.array(z.string()).optional().describe("Additional exports"),
  styles: z.record(StyleSchema).optional().describe("Component styles"),
});

export type ComponentProp = z.infer<typeof ComponentPropSchema>;
export type Component = z.infer<typeof ComponentSchema>;

/**
 * Generate TypeScript interface from props
 */
export function generatePropsInterface(
  name: string,
  props: ComponentProp[],
): string {
  if (props.length === 0) {
    return "";
  }

  const lines = props.map((prop) => {
    const optional = prop.required ? "" : "?";
    const propType = mapPropType(prop.type);
    const comment = prop.description ? `  /** ${prop.description} */\n` : "";

    return `${comment}  ${prop.name}${optional}: ${propType};`;
  });

  return `export interface ${name}Props {
${lines.join("\n")}
}`;
}

/**
 * Map schema prop type to TypeScript type
 */
function mapPropType(type: ComponentProp["type"]): string {
  switch (type) {
    case "string":
      return "string";
    case "number":
      return "number";
    case "boolean":
      return "boolean";
    case "function":
      return "() => void";
    case "object":
      return "Record<string, unknown>";
    case "array":
      return "unknown[]";
    case "signal":
      return "Accessor<unknown>";
    case "store":
      return "Store<unknown>";
    default:
      return "unknown";
  }
}

/**
 * Generate complete component file from schema
 */
export function generateComponentFile(component: Component): string {
  const hasProps = component.props.length > 0;
  const propsInterface = hasProps
    ? generatePropsInterface(component.name, component.props)
    : "";

  const imports = component.imports.length > 0
    ? component.imports.join("\n") + "\n\n"
    : "";

  const propsParam = hasProps ? `props: ${component.name}Props` : "";

  const additionalExports = component.exports && component.exports.length > 0
    ? "\n\n" + component.exports.join("\n")
    : "";

  return `/**
 * ${component.description}
 */

${imports}${propsInterface ? propsInterface + "\n\n" : ""}export function ${component.name}(${propsParam}) {
${component.body}
}${additionalExports}
`;
}

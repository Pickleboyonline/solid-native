use std::collections::HashMap;

/// JavaScript value types that can be passed across the FFI boundary
/// This enum represents the subset of JS types we support for properties
#[derive(Debug, Clone, PartialEq, uniffi::Enum)]
pub enum JSValue {
    /// Null value
    Null,
    /// Undefined value
    Undefined,
    /// Boolean value
    Boolean { value: bool },
    /// Number value (f64 in JavaScript)
    Number { value: f64 },
    /// String value
    String { value: String },
    /// Array of JSValues
    Array { values: Vec<JSValue> },
    /// Object represented as key-value pairs
    Object { properties: HashMap<String, JSValue> },
}

impl JSValue {
    /// Creates a null JSValue
    pub fn null() -> Self {
        JSValue::Null
    }

    /// Creates an undefined JSValue
    pub fn undefined() -> Self {
        JSValue::Undefined
    }

    /// Creates a boolean JSValue
    pub fn boolean(value: bool) -> Self {
        JSValue::Boolean { value }
    }

    /// Creates a number JSValue
    pub fn number(value: f64) -> Self {
        JSValue::Number { value }
    }

    /// Creates a string JSValue
    pub fn string(value: impl Into<String>) -> Self {
        JSValue::String {
            value: value.into(),
        }
    }

    /// Creates an array JSValue
    pub fn array(values: Vec<JSValue>) -> Self {
        JSValue::Array { values }
    }

    /// Creates an object JSValue from a HashMap
    pub fn object(properties: HashMap<String, JSValue>) -> Self {
        JSValue::Object { properties }
    }

    /// Creates an object JSValue from key-value pairs
    pub fn object_from_pairs(pairs: Vec<(String, JSValue)>) -> Self {
        let properties: HashMap<String, JSValue> = pairs.into_iter().collect();
        JSValue::Object { properties }
    }

    /// Checks if the value is null
    pub fn is_null(&self) -> bool {
        matches!(self, JSValue::Null)
    }

    /// Checks if the value is undefined
    pub fn is_undefined(&self) -> bool {
        matches!(self, JSValue::Undefined)
    }

    /// Checks if the value is a boolean
    pub fn is_boolean(&self) -> bool {
        matches!(self, JSValue::Boolean { .. })
    }

    /// Checks if the value is a number
    pub fn is_number(&self) -> bool {
        matches!(self, JSValue::Number { .. })
    }

    /// Checks if the value is a string
    pub fn is_string(&self) -> bool {
        matches!(self, JSValue::String { .. })
    }

    /// Checks if the value is an array
    pub fn is_array(&self) -> bool {
        matches!(self, JSValue::Array { .. })
    }

    /// Checks if the value is an object
    pub fn is_object(&self) -> bool {
        matches!(self, JSValue::Object { .. })
    }

    /// Gets the boolean value if this is a boolean
    pub fn as_boolean(&self) -> Option<bool> {
        match self {
            JSValue::Boolean { value } => Some(*value),
            _ => None,
        }
    }

    /// Gets the number value if this is a number
    pub fn as_number(&self) -> Option<f64> {
        match self {
            JSValue::Number { value } => Some(*value),
            _ => None,
        }
    }

    /// Gets the string value if this is a string
    pub fn as_string(&self) -> Option<&str> {
        match self {
            JSValue::String { value } => Some(value),
            _ => None,
        }
    }

    /// Gets the array values if this is an array
    pub fn as_array(&self) -> Option<&Vec<JSValue>> {
        match self {
            JSValue::Array { values } => Some(values),
            _ => None,
        }
    }

    /// Gets the object properties if this is an object
    pub fn as_object(&self) -> Option<&HashMap<String, JSValue>> {
        match self {
            JSValue::Object { properties } => Some(properties),
            _ => None,
        }
    }

    /// Parses a JSON string into a JSValue
    pub fn from_json(json: &str) -> Self {
        // Use serde_json to parse the JSON string
        match serde_json::from_str::<serde_json::Value>(json) {
            Ok(v) => Self::from_serde_value(&v),
            Err(_) => JSValue::Undefined,
        }
    }

    /// Converts a serde_json::Value to a JSValue
    fn from_serde_value(value: &serde_json::Value) -> Self {
        match value {
            serde_json::Value::Null => JSValue::Null,
            serde_json::Value::Bool(b) => JSValue::Boolean { value: *b },
            serde_json::Value::Number(n) => {
                JSValue::Number {
                    value: n.as_f64().unwrap_or(0.0),
                }
            }
            serde_json::Value::String(s) => JSValue::String { value: s.clone() },
            serde_json::Value::Array(arr) => {
                JSValue::Array {
                    values: arr.iter().map(Self::from_serde_value).collect(),
                }
            }
            serde_json::Value::Object(obj) => {
                let properties: HashMap<String, JSValue> = obj
                    .iter()
                    .map(|(k, v)| (k.clone(), Self::from_serde_value(v)))
                    .collect();
                JSValue::Object { properties }
            }
        }
    }

    /// Converts the JSValue to a simple string representation
    /// For simple display/serialization purposes
    pub fn to_simple_string(&self) -> String {
        match self {
            JSValue::Null => "null".to_string(),
            JSValue::Undefined => "undefined".to_string(),
            JSValue::Boolean { value } => value.to_string(),
            JSValue::Number { value } => value.to_string(),
            JSValue::String { value } => value.clone(),
            JSValue::Array { values } => format!("[{} items]", values.len()),
            JSValue::Object { properties } => format!("{{{}  properties}}", properties.len()),
        }
    }
}

// Implement Display for easy printing
impl std::fmt::Display for JSValue {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            JSValue::Null => write!(f, "null"),
            JSValue::Undefined => write!(f, "undefined"),
            JSValue::Boolean { value } => write!(f, "{}", value),
            JSValue::Number { value } => write!(f, "{}", value),
            JSValue::String { value } => write!(f, "\"{}\"", value),
            JSValue::Array { values } => {
                write!(f, "[")?;
                for (i, val) in values.iter().enumerate() {
                    if i > 0 {
                        write!(f, ", ")?;
                    }
                    write!(f, "{}", val)?;
                }
                write!(f, "]")
            }
            JSValue::Object { properties } => {
                write!(f, "{{")?;
                let mut items: Vec<_> = properties.iter().collect();
                items.sort_by_key(|(k, _)| *k);
                for (i, (key, val)) in items.iter().enumerate() {
                    if i > 0 {
                        write!(f, ", ")?;
                    }
                    write!(f, "\"{}\": {}", key, val)?;
                }
                write!(f, "}}")
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_jsvalue_null() {
        let val = JSValue::null();
        assert!(val.is_null());
        assert_eq!(val.to_simple_string(), "null");
    }

    #[test]
    fn test_jsvalue_boolean() {
        let val = JSValue::boolean(true);
        assert!(val.is_boolean());
        assert_eq!(val.as_boolean(), Some(true));
        assert_eq!(val.to_simple_string(), "true");
    }

    #[test]
    fn test_jsvalue_number() {
        let val = JSValue::number(42.5);
        assert!(val.is_number());
        assert_eq!(val.as_number(), Some(42.5));
        assert_eq!(val.to_simple_string(), "42.5");
    }

    #[test]
    fn test_jsvalue_string() {
        let val = JSValue::string("hello");
        assert!(val.is_string());
        assert_eq!(val.as_string(), Some("hello"));
        assert_eq!(val.to_simple_string(), "hello");
    }

    #[test]
    fn test_jsvalue_array() {
        let val = JSValue::array(vec![JSValue::number(1.0), JSValue::string("test")]);
        assert!(val.is_array());
        assert_eq!(val.as_array().unwrap().len(), 2);
    }

    #[test]
    fn test_jsvalue_object() {
        let val = JSValue::object_from_pairs(vec![
            ("name".to_string(), JSValue::string("John")),
            ("age".to_string(), JSValue::number(30.0)),
        ]);
        assert!(val.is_object());
        assert_eq!(val.as_object().unwrap().len(), 2);
    }

    #[test]
    fn test_jsvalue_display() {
        let val = JSValue::string("test");
        assert_eq!(format!("{}", val), "\"test\"");

        let val = JSValue::number(42.0);
        assert_eq!(format!("{}", val), "42");

        let val = JSValue::boolean(true);
        assert_eq!(format!("{}", val), "true");
    }
}

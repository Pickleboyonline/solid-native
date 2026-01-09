/// Error types for SNCore
#[derive(Debug, PartialEq, thiserror::Error, uniffi::Error)]
pub enum SNCoreError {
    #[error("Runtime error: {msg}")]
    RuntimeError { msg: String },

    #[error("Renderer error: {msg}")]
    RendererError { msg: String },

    #[error("JavaScript evaluation error: {msg}")]
    JsEvalError { msg: String },

    #[error("Context error: {msg}")]
    ContextError { msg: String },
}

impl SNCoreError {
    pub fn runtime(msg: impl Into<String>) -> Self {
        Self::RuntimeError { msg: msg.into() }
    }

    pub fn renderer(msg: impl Into<String>) -> Self {
        Self::RendererError { msg: msg.into() }
    }

    pub fn js_eval(msg: impl Into<String>) -> Self {
        Self::JsEvalError { msg: msg.into() }
    }

    pub fn context(msg: impl Into<String>) -> Self {
        Self::ContextError { msg: msg.into() }
    }
}

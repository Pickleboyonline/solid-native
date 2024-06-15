use std::{borrow::Borrow, collections::HashMap};

use anyhow::{bail, Context, Ok, Result};
use deno_core::v8::{self, Handle};
use taffy::prelude::*;
use thiserror::Error;

#[derive(Error, Debug)]
pub enum PaperMoonError {
    #[error("Root Node does not exist!")]
    RootNodeNotFound,
}
/// Placeholder to represent JSValue
pub struct Value {}

type Prop = HashMap<String, Value>;

/// Placeholder for foriegn function implimentation.
/// Recieves changes about props, etc...
struct ViewHostReciever {}

/// Owns all of what we need for the system.
/// effectivly "owns" our system
/// TODO: Do not maintain anything in the tree that we do not need.
pub struct Tree {
    taffy_tree: TaffyTree<()>,
    props: HashMap<NodeId, Prop>,
    view_host_reciever: HashMap<NodeId, ViewHostReciever>,
    root_node_id: Option<NodeId>,
}

impl Tree {
    pub fn new() -> Self {
        Tree {
            props: HashMap::new(),
            view_host_reciever: HashMap::new(),
            taffy_tree: TaffyTree::new(),
            root_node_id: None,
        }
    }

    /// Basically call this when you mutate the tree and need to update the tree.
    /// TODO: We need to notify of layout and props at same time
    /// Only call on insert/remove and update to style prop
    /// if style prop is updated we send the new style and the new layout to the reciever
    /// Not every style causes a dirty tree (for example, background color)
    /// Later on you could also just not call this func if there are no keys related to
    /// flex layout, since they dont update the layout
    fn update_layout_and_notify_reciever(
        &mut self,
        styled_node_id: Option<NodeId>,
        style_prop_entry: Option<(String, Value)>,
    ) -> Result<()> {
        /*
        It is easy for the case of insert and remove, but setting props is tricker.
        When props are set, layout can change but we have to calculate the new layout.
        Base Clase:
        Set Props/Styles => Call this func. => Calulcate layout => notify reciever.
        If layout change, notify of both. => Not necessary. We need to
        ---
        Basically, we take in the reciever node id and the new styles
         */

        // Calculate layout starting at root node.

        // Perform BFS to notify of new layout, if the node changed, notify of its style as well

        // if the node that was dirty did not change, update it as well
        Ok(())
    }

    pub fn get_root_node(&self) -> Result<NodeId> {
        match self.root_node_id {
            Some(id) => Ok(id),
            None => {
                bail!(PaperMoonError::RootNodeNotFound);
            }
        }
    }

    pub fn create_root_node(&mut self, view_host_reciever_type: &str) -> Result<NodeId> {
        let root_id = self.create_node(view_host_reciever_type)?;
        self.root_node_id = Some(root_id);
        Ok(root_id)
    }

    pub fn create_node(&mut self, view_host_reciever_type: &str) -> Result<NodeId> {
        // TODO: Lookup view host reciever type provided by host platform
        let node_id = self.taffy_tree.new_leaf(Style::default())?;
        self.props.insert(node_id, HashMap::new());
        Ok(node_id)
    }

    /// Instead of doing type, maybe just a trait as long as it convert to
    /// Some primative.
    /// TODO: Recacalutate the layout
    pub fn set_property(&mut self, node_id: NodeId, property_name: String, value: Option<Value>) {
        // Update prop in map
        // Since style comes in as one object, we can use that to just calculate the
        // if its a stle just use that
        if property_name == "style" {
            match &value {
                Some(jsvalue) => {
                    // TODO: Call the style thing
                    let style = compute_style_from_jsvalue(jsvalue);
                    self.taffy_tree
                        .set_style(node_id, style)
                        .expect("Could not set new style");
                }
                None => {
                    self.taffy_tree
                        .set_style(node_id, Style::default())
                        .expect("Could not set default style");
                }
            }
        }

        // self.taffy_tree.compute_layout(node, available_space).e;

        match self.props.get_mut(&node_id) {
            Some(prop_map) => {
                if let Some(value) = value {
                    prop_map.insert(property_name, value);
                }
            }
            None => {}
        }
    }

    pub fn is_text_node(&self, node_id: NodeId) -> bool {
        false
    }

    pub fn remove_node(&mut self, node_id: NodeId, child_id: NodeId) -> Result<()> {
        self.taffy_tree.remove_child(node_id, child_id)?;

        self.props.remove(&child_id);

        self.update_layout_and_notify_reciever(None, None)?;
        Ok(())
    }

    pub fn get_parent_node(&self, node_id: NodeId) {}

    pub fn get_first_child(&self, node_id: NodeId) {}

    pub fn get_next_sibling(&self, node_id: NodeId) {}
}

/// Need to return a style.
/// Start with the default
/// Collect keys
/// Pattern match and convert type
/// Set on tyle
/// Move style to caller
/// TODO: Flesh this out
fn compute_style_from_jsvalue(jsvalue: &Value) -> Style {
    Style::default()
}

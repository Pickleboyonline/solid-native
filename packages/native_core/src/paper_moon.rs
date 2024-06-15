use std::collections::HashMap;

use deno_core::Extension;
use taffy::prelude::*;

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
    viewHostReciever: HashMap<NodeId, ViewHostReciever>
}

impl Tree {
    pub fn new() -> Self {
        Tree {
            props: HashMap::new(),
            viewHostReciever: HashMap::new(),
            taffy_tree: TaffyTree::new()
        }
    }

    pub fn create_element(&mut self, view_host_reciever_type: &str) {
        // TODO: Lookup view host reciever type provided by host platform
        self.taffy_tree.new_leaf(Style::default()).unwrap();
        
    }

    pub fn set_property(&mut self, node_id: NodeId, property_name: String, value: Option<Value>) {
        // Update prop in map
        // Since style comes in as one object, we can use that to just calculate the 
        // self.taffy_tree.set_style(node, style)
    }

    pub fn is_text_node(&self, node_id: NodeId) -> bool {
        false
    }

    pub fn remove_node(&mut self, node_id: NodeId,child_idd: NodeId) {

    }

    pub fn get_parent_node(&self, node_id: NodeId) {

    }

    pub fn get_first_child(&self, node_id: NodeId) {

    }

    pub fn get_next_sibling(&self, node_id: NodeId) {

    }
}

fn compute_style_from_jsvalue() {
    
}
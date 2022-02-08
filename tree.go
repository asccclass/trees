package main

import(
)

type TreeNoe struct {
   Left		*TreeNode
   Right	*TreeNode
   Val		interface{}
}

type NodeQueue struct {
   Nodes	[]*TreeNode
}

func(q *NodeQueus) IsEmpty()(bool)  { 
   return len(q.Nodes) == 0 
}

// push a node in the queue
func(q *NodeQueue) Push(nodes ...*TreeNode) { 
   q.Nodes = append(q.Nodes, nodes...) 
}

// Pop nodes from the queue
func(q *NodeQueue) Pop()(*TreeNode) {
   if !q.IsEmpty() {
       n := q.Nodes[0]
       q.Nodes = q.Nodes[1:]
       return n
   }
   return nil
}

func CreateNodeQueue()(*NodeQueue) {
   return  &NodeQueue {
   }
}

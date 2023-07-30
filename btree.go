package main

type Node struct {
	PreNode *Node
	NextNode *Node
	ChildTreeNode *TreeNode
	CurTreeNode *TreeNoede
	Value   int
	Id  int
}

type TreeNode struct {
	ParentNode *Node
	LastNode  *Node
	FirstNode *Node
	Length  int
}

var m int 
m = 5
var RootTreeNode *TreeNode
var sqlNode *Node
// b+树插入
func InsertBtree(root Node, value int, id int) {
	if root == nil {
		treeNode := &TreeNode{Length:1}
		newNode := &Node{Value:value,Id: id,CurTreeNode:treeNode}
		treeNode.FirstNode = newNode
		treeNode.LastNode = newNode
		root = newNode
		RootTreeNode = treeNode
		sqlNode = root
		return 
	}
	if root.ChildTreeNode == nil {
		if root.Value >= vlaue {
			newNode := &Node{Value:value,Id: id,CurTreeNode:root.CurTreeNode}
			if root.PreNode == nil {
				newNode.NextNode = root
				root.PreNode = newNode
				newNode.CurTreeNode.FirstNode = newNode
				newNode.CurTreeNode.Length++
			}else {
				newNode.PreNode = root.PreNode
				root.PreNode.NextNode = newNode
				newNode.NextNode = root
				root.PreNode = newNode
				newNode.CurTreeNode.Length++
			}
			if newNode.CurTreeNode.Length > m {
				SpiltTreeNode(newNode)
			}
			
		}else {
			if root.NextNode==nil {
				newNode := &Node{Value:value,Id: id,,CurTreeNode:root.CurTreeNode}
			    newNode.PreNode = root
				root.NextNode = newNode
				newNode.CurTreeNode.LastNode = newNode
				newNode.CurTreeNode.Length++
				if newNode.CurTreeNode.Length > m {
					SpiltTreeNode(newNode)
				}
		    }else {
				InsertBtree(root.NextNode, value, id)
			} 
			
		}
		return
	}
	
	if root.ChildTreeNode != nil {
		if root.Value >= vlaue {
			InsertBtree(root.ChildTreeNode.FirstNode,value,id)

		}else  {
			if root.NextNode==nil {
				newNode := &Node{
					Value:value,
					Id: id,
					CurTreeNode:root.CurTreeNode,
				}
				root.ChildTreeNode.Length++
				root.ChildTreeNode.LastNode = newNode
				newNode.PreNode = root
				root.NextNode = newNode
				
				
			}else {
				InsertBtree(root.NextNode,value,id)
			}
		}
		
		return
	}

}


func SpiltTreeNode(curNode *node){
	if curNode.CurTreeNode.Length > m {
		tempNode := curNode.CurTreeNode.FirstNode
		last1Node := tempNode
		for i := 0; i < (m-1)/2; i++ {
			last1Node := last1Node.NextNode

		}
		leftTreeNode := &TreeNode{Length:(m+1)/2,FirstNode:tempNode,LastNode:last1Node}
		first2Node = last1Node.NextNode
		rightTreeNode :=curNode.CurTreeNode
		rightTreeNode.Length = rightTreeNode.Length - (m+1)/2
		rightTreeNode.FirstNode = first2Node
		UpTreeNode(leftTreeNode,rightTreeNode)
	}
	
}

func UpTreeNode(leftTreeNode *TreeNode, rightTreeNode, *TreeNode){
	if rightTreeNode.ParentNode == nil {
		treeNode := &TreeNode{Length:1}
		newNode := &Node{Value:rightTreeNode.LastNode.Value,CurTreeNode:treeNode}
		treeNode.FirstNode = newNode
		treeNode.LastNode = newNode
		rightTreeNode.ParentNode = newNode
		RootTreeNode = treeNode
	}
	if leftTreeNode.ParentNode == nil {
		newNode := &Node{Value:leftTreeNode.LastNode.Value,CurTreeNode:rightTreeNode.ParentNode.CurTreeNode}
		if rightTreeNode.ParentNode.PreNode == nil {
			rightTreeNode.ParentNode.PreNode = newNode
			newNode.NextNode = rightTreeNode.ParentNode
		}else {
			rightTreeNode.ParentNode.PreNode.NextNode = newNode
			newNode.PreNode = rightTreeNode.ParentNode.PreNode
			rightTreeNode.ParentNode.PreNode = newNode
			newNode.NextNode = rightTreeNode.ParentNode
		}
		rightTreeNode.ChildTreeNode.Length++
		SpiltTreeNode(newNode)
	}
}
func DeleteBtree(root Node, value int, id int) {
	if root.ChildTreeNode == nil {
		if value == root.Value {
			if root.CurTreeNode.FirstNode == root {
				root.CurTreeNode.FirstNode = root.NextNode
			}
			if root.CurTreeNode.LastNode == root {
				root.CurTreeNode.LastNode = root.PreNode
			}
			if root.PreNode != nil {
				root.PreNode.NextNode =root.NextNode
			}
			if root.NextNode != nil {
				root.NextNode.PreNode = root.PreNode
			}
			
			root.CurTreeNode.Length-- 
			if root.CurTreeNode.Length < (m+1)/2 {
				if root.CurTreeNode.ParentNode.NextNode != nil {
					Merge(root.CurTreeNode,root.CurTreeNode.ParentNode.NextNode.ChildTreeNode,1)
				}else {
			 	    Merge(root.CurTreeNode.ParentNode.PreNode.ChildTreeNode,root.CurTreeNode,2)
				}
				
			}

		}else if value > root.Value {
			DeleteBtree(root.NextNode,value,id)
		}
	}

	if value <= root.Value  {
		DeleteBtree(root.ChildTreeNode.FirstNode,value,id)
	}else {
		DeleteBtree(root.NextNode,value,id)
	}

}

func Merge(leftTreeNode *TreeNode, rightTreeNode *TreeNode, op int) {
	leftFirstNode, leftLastNode= leftTreeNode.FirstNode, leftTreeNode.LastNode
	rightFirstNode, rightLastNode= rightTreeNode.FirstNode, rightTreeNode.LastNode
	leftLastNode.Next = rightFirstNode
	rightFirstNode.Pre = leftLastNode
	if leftTreeNode.Length > (m+1)/2 {
		leftTreeNode.LastNode = leftLastNode.PreNode
		rightTreeNode.FirstNode = leftLastNode
		
		
	}else if rightTreeNode.Length > (m+1)/2 {
		leftTreeNode.LastNode = rightFirstNode
		rightTreeNode.FirstNode = rightFirstNode.Next
	}else {
		rightTreeNode.FirstNode = leftFirstNode
		rightTreeNode.Length += leftTreeNode.Length
		deleteSingle(leftTreeNode)


	}
	Replace(leftTreeNode.LastNode)
	Replace(rightTreeNode.LastNode)
	
}
func Replace(curNode *Node) {
	if curNode.CurTreeNode.LastNode == curNode{
		if curNode.Value != curNode.CurTreeNode.ParentNode.Value {
			urNode.CurTreeNode.ParentNode.Value = curNode.Value
			Replace(curNode.CurTreeNode.ParentNode)
		}
	}

}

func deleteSingle(curTreeNode *Node) {
	root := curTreeNode.ParentNode
	if root == nil {
		return
	}
	if root.PreNode != nil {
		root.PreNode.NextNode =root.NextNode
	}
	if root.NextNode != nil {
		root.NextNode.PreNode = root.PreNode
	}
	root.CurTreeNode.Length--
	if root.CurTreeNode.Length < (m+1)/2 {
		if root.CurTreeNode.ParentNode.NextNode != nil {
			Merge(root.CurTreeNode,root.CurTreeNode.ParentNode.NextNode.ChildTreeNode,1)
		}else {
			Merge(root.CurTreeNode.ParentNode.PreNode.ChildTreeNode,root.CurTreeNode,2)
		}
		
	}
	root = nil
	curTreeNode = nil
	
	
}
func main(){

}

package main

import (
	"fmt"
)

type Node struct {
	PreNode       *Node
	NextNode      *Node
	ChildTreeNode *TreeNode
	CurTreeNode   *TreeNode
	Value         int
	Id            int
}

type TreeNode struct {
	ParentNode *Node
	LastNode   *Node
	FirstNode  *Node
	Length     int
}

var (
	m = 5
)
var BtreeNodeSlice = make([][]*TreeNode, 100)
var RootTreeNode *TreeNode
var rootNode *Node
var sqlNode *Node

// b+树插入
func InsertBtree(root *Node, value int, id int) {
	fmt.Printf("====llvalue = %d,",value)
	if root == nil {
		treeNode := &TreeNode{Length: 1}
		newNode := &Node{Value: value, Id: id, CurTreeNode: treeNode}
		treeNode.FirstNode = newNode
		treeNode.LastNode = newNode
		rootNode = newNode
		RootTreeNode = treeNode
		sqlNode = rootNode
		return
	}

	if root.ChildTreeNode == nil {
		if root.Value >= value {
			newNode := &Node{Value: value, Id: id, CurTreeNode: root.CurTreeNode}
		//	fmt.Printf("llvalue = %d, root.Value= %d,root.CurTreeNode=%+v\n", value, root.Value,root.CurTreeNode)
			if root == root.CurTreeNode.FirstNode {
				fmt.Printf("rrrrrvalue = %d, root.Value= %d curTreeNode=%+v\n", value, root.Value, root.CurTreeNode)
				newNode.NextNode = root
				root.PreNode = newNode
				newNode.CurTreeNode.FirstNode = newNode
				newNode.CurTreeNode.Length++
			} else {

				newNode.PreNode = root.PreNode
				root.PreNode.NextNode = newNode
				newNode.NextNode = root
				root.PreNode = newNode
				newNode.CurTreeNode.Length++
			}
			fmt.Printf(" newNode=%p, newNode=%+v,,newNode.CurTreeNode=%p ,newNode.CurTreeNode=%+v\n", newNode,newNode,newNode.CurTreeNode,newNode.CurTreeNode)
			if newNode.CurTreeNode.Length > m {
				SpiltTreeNode(newNode)
			}
			fmt.Printf("==value = %d, root.Value= %d curTreeNode=%p curTreeNode=%+v\n", value, root.Value,root.CurTreeNode, root.CurTreeNode)
		} else {

			if root == root.CurTreeNode.LastNode {
				newNode := &Node{Value: value, Id: id, CurTreeNode: root.CurTreeNode}
				newNode.PreNode = root
				root.NextNode = newNode
				newNode.CurTreeNode.LastNode = newNode
				newNode.CurTreeNode.Length++
				Replace(newNode)
				if newNode.CurTreeNode.Length > m {
					SpiltTreeNode(newNode)
				}
			} else {
				InsertBtree(root.NextNode, value, id)
			}

		}
		return
	}

	if root.ChildTreeNode != nil {
		if root.Value >= value {
			InsertBtree(root.ChildTreeNode.FirstNode, value, id)

		} else {
			if root == root.CurTreeNode.LastNode {
				InsertBtree(root.ChildTreeNode.FirstNode, value, id)
			} else {
				InsertBtree(root.NextNode, value, id)
			}

		}

		return
	}

}

func SpiltTreeNode(curNode *Node) {
	if curNode.CurTreeNode.Length > m {
		rightTreeNode := curNode.CurTreeNode
		tempNode := curNode.CurTreeNode.FirstNode
		last1Node := tempNode
		leftTreeNode := &TreeNode{}
		for i := 0; i < (m-1)/2; i++ {
			last1Node.CurTreeNode = leftTreeNode
			last1Node = last1Node.NextNode
		}
		last1Node.CurTreeNode = leftTreeNode
		leftTreeNode.Length = (m + 1) / 2
		leftTreeNode.FirstNode = tempNode
		leftTreeNode.LastNode = last1Node
		first2Node := last1Node.NextNode
		
		rightTreeNode.Length = rightTreeNode.Length - (m+1)/2
		rightTreeNode.FirstNode = first2Node
	//	fmt.Printf("leftTreeNode= %p,leftTreeNode=%+v,rightTreeNode= %p,rightTreeNode=%+v\n",leftTreeNode,leftTreeNode,rightTreeNode,rightTreeNode)
		UpTreeNode(leftTreeNode, rightTreeNode)
	}

}

func UpTreeNode(leftTreeNode *TreeNode, rightTreeNode *TreeNode) {
	//fmt.Printf("y111leftTreeNode=%+v\n",leftTreeNode)
	if rightTreeNode.ParentNode == nil {
		treeNode := &TreeNode{Length: 1}
		newNode := &Node{Value: rightTreeNode.LastNode.Value, CurTreeNode: treeNode, ChildTreeNode: rightTreeNode}
		treeNode.FirstNode = newNode
		treeNode.LastNode = newNode
		rightTreeNode.ParentNode = newNode
		RootTreeNode = treeNode
		fmt.Printf("yy left newNode = %+v, newNode.CurTreeNode=%+v\n",newNode, newNode.CurTreeNode)

	}
	fmt.Printf("y111leftTreeNode=%+v\n",leftTreeNode)
	if leftTreeNode.ParentNode == nil {
		//fmt.Printf("eawwfg\n")
		fmt.Printf("leftTreeNode.LastNode.Value=%p,rightTreeNode.ParentNode.CurTreeNode=%p,  leftTreeNode=%p\n",leftTreeNode.LastNode.Value,rightTreeNode.ParentNode.CurTreeNode,  leftTreeNode)
		newNode := &Node{Value: leftTreeNode.LastNode.Value, CurTreeNode: rightTreeNode.ParentNode.CurTreeNode, ChildTreeNode: leftTreeNode}
		fmt.Printf("rightTreeNode.ParentNode = %p, rightTreeNode.ParentNode=%+v\n",rightTreeNode.ParentNode, rightTreeNode.ParentNode)

		if rightTreeNode.ParentNode == rightTreeNode.ParentNode.CurTreeNode.FirstNode {
			rightTreeNode.ParentNode.PreNode = newNode
			newNode.NextNode = rightTreeNode.ParentNode
			rightTreeNode.ParentNode.CurTreeNode.FirstNode = newNode
		} else {

			rightTreeNode.ParentNode.PreNode.NextNode = newNode
			newNode.PreNode = rightTreeNode.ParentNode.PreNode
			rightTreeNode.ParentNode.PreNode = newNode
			newNode.NextNode = rightTreeNode.ParentNode
		}
		leftTreeNode.ParentNode = newNode
		fmt.Printf("yy right newNode = %+v, newNode.CurTreeNode=%+v\n",newNode, newNode.CurTreeNode)
		newNode.CurTreeNode.Length++
		SpiltTreeNode(newNode)
	}
}
func DeleteBtree(root *Node, value int, id int) {
	if root.ChildTreeNode == nil {
		if value == root.Value {
			if root.CurTreeNode.FirstNode == root {
				root.CurTreeNode.FirstNode = root.NextNode
			}
			if root.CurTreeNode.LastNode == root {
				root.CurTreeNode.LastNode = root.PreNode
			}
			if root.PreNode != nil {
				root.PreNode.NextNode = root.NextNode
			}
			if root.NextNode != nil {
				root.NextNode.PreNode = root.PreNode
			}

			root.CurTreeNode.Length--
			if root.CurTreeNode.Length < (m+1)/2 {
				if root.CurTreeNode.ParentNode.NextNode != nil {
					Merge(root.CurTreeNode, root.CurTreeNode.ParentNode.NextNode.ChildTreeNode)
				} else {
					Merge(root.CurTreeNode.ParentNode.PreNode.ChildTreeNode, root.CurTreeNode)
				}

			}

		} else if value > root.Value {
			DeleteBtree(root.NextNode, value, id)
		}
	}

	if value <= root.Value {
		DeleteBtree(root.ChildTreeNode.FirstNode, value, id)
	} else {
		DeleteBtree(root.NextNode, value, id)
	}

}

func Merge(leftTreeNode *TreeNode, rightTreeNode *TreeNode) {
	leftFirstNode, leftLastNode := leftTreeNode.FirstNode, leftTreeNode.LastNode
	rightFirstNode := rightTreeNode.FirstNode
	leftLastNode.NextNode = rightFirstNode
	rightFirstNode.PreNode = leftLastNode
	if leftTreeNode.Length > (m+1)/2 {
		leftTreeNode.LastNode = leftLastNode.PreNode
		rightTreeNode.FirstNode = leftLastNode

	} else if rightTreeNode.Length > (m+1)/2 {
		leftTreeNode.LastNode = rightFirstNode
		rightTreeNode.FirstNode = rightFirstNode.NextNode
	} else {
		rightTreeNode.FirstNode = leftFirstNode
		rightTreeNode.Length += leftTreeNode.Length
		deleteSingle(leftTreeNode)

	}
	Replace(leftTreeNode.LastNode)
	Replace(rightTreeNode.LastNode)

}
func Replace(curNode *Node) {
	if curNode.CurTreeNode.ParentNode == nil {
		return
	}
	if curNode.CurTreeNode.LastNode == curNode {
		if curNode.Value != curNode.CurTreeNode.ParentNode.Value {
			curNode.CurTreeNode.ParentNode.Value = curNode.Value
			Replace(curNode.CurTreeNode.ParentNode)
		}
	}

}

func deleteSingle(curTreeNode *TreeNode) {
	root := curTreeNode.ParentNode
	if root == nil {
		return
	}
	if root.PreNode != nil {
		root.PreNode.NextNode = root.NextNode
	}
	if root.NextNode != nil {
		root.NextNode.PreNode = root.PreNode
	}
	root.CurTreeNode.Length--
	if root.CurTreeNode.Length < (m+1)/2 {
		if root.CurTreeNode.ParentNode.NextNode != nil {
			Merge(root.CurTreeNode, root.CurTreeNode.ParentNode.NextNode.ChildTreeNode)
		} else {
			Merge(root.CurTreeNode.ParentNode.PreNode.ChildTreeNode, root.CurTreeNode)
		}

	}
	root = nil
	curTreeNode = nil

}

func FindBtree(root *Node, value int) *Node {
	if root.ChildTreeNode == nil {
		if value <= root.Value {
			return root

		} else if value > root.Value {
			return FindBtree(root.NextNode, value)
		}
		return nil
	}
	if value <= root.Value {
		return FindBtree(root.ChildTreeNode.FirstNode, value)
	} else {
		return FindBtree(root.NextNode, value)
	}

}

func GetBtreeSlice(rootTree *TreeNode, level int) int {
	if rootTree == nil {
		return level

	}
	ans := 0
	BtreeNodeSlice[level] = append(BtreeNodeSlice[level], rootTree)
	nodei := rootTree.FirstNode
	for nodei = rootTree.FirstNode; nodei != rootTree.LastNode; nodei = nodei.NextNode {
		ans = GetBtreeSlice(nodei.ChildTreeNode, level+1)
	}
	ans = GetBtreeSlice(nodei.ChildTreeNode, level+1)
	return ans
}
func PrintBtree() {
	BtreeNodeSlice = make([][]*TreeNode, 100)
	tt := GetBtreeSlice(RootTreeNode, 0)
	for i := 0; i < tt; i++ {
		for j := 0; j < len(BtreeNodeSlice[i]); j++ {
			printSignleBtree(BtreeNodeSlice[i][j])
		}
	//	fmt.Println("==============")
	}

	for i := 0; i < tt; i++ {
		for j := 0; j < len(BtreeNodeSlice[i]); j++ {
			printSignleBtreeValue(BtreeNodeSlice[i][j])
		}
		fmt.Println("")
	}
	
}

func printSignleBtree(rootTree *TreeNode) {
	// nodei := rootTree.FirstNode
	// for nodei = rootTree.FirstNode; nodei != rootTree.LastNode; nodei = nodei.NextNode {
	// 	fmt.Printf("%d-+ treeNode:%p node:%p\n", nodei.Value, nodei.CurTreeNode, nodei)
	// }
	// fmt.Printf("%d-+ treeNode:%p node:%p\n", nodei.Value, nodei.CurTreeNode, nodei)
}

func printSignleBtreeValue(rootTree *TreeNode) {
	nodei := rootTree.FirstNode
	for nodei = rootTree.FirstNode; nodei != rootTree.LastNode; nodei = nodei.NextNode {
		fmt.Printf("%d-", nodei.Value)
	}
	fmt.Printf("%d  ", nodei.Value)
}
func main() {
	InsertBtree(rootNode, 50, 50)
	for i := 2; i <= 8; i++ {
		for j := i * 10; j > i * 10 -10; j--{
			temp := RootTreeNode.FirstNode
			InsertBtree(temp, j, j)
			PrintBtree()
		}

	}

	// for i := 40; i >= 10; i-- {
	// 	temp := RootTreeNode.FirstNode
	// 		InsertBtree(temp, i, i)
	// 		PrintBtree()
	// 	}
	PrintBtree()
	
	// for i := 6; i >= 2; i-- {
	// 	InsertBtree(RootTreeNode.FirstNode, i, i)
	// }
	// PrintBtree()
	// fmt.Printf("//////\n")
	// InsertBtree(RootTreeNode.FirstNode, 1, 1)
	
}

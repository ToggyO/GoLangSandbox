package binary_tree

import (
	"hello/packages/data_structures/models"
)

type TreeTraverser[T interface{}] struct {
	tree *BinaryTree[T]
}

func NewTreeTraverser[T interface{}](tree *BinaryTree[T]) *TreeTraverser[T] {
	return &TreeTraverser[T]{tree: tree}
}

func (t *TreeTraverser[T]) Postorder(action func(node *models.TreeNode[T])) {
	t.postorder(t.tree.root, action)
}

func (t *TreeTraverser[T]) postorder(node *models.TreeNode[T], action func(node *models.TreeNode[T])) {
	if node == nil {
		return
	}

	action(node)

	t.postorder(node.GetLeftNode(), action)
	t.postorder(node.GetRightNode(), action)
	println(node.GetValue())
}

//public class Program
//{
//	public static void Main()
//	{
//		BinaryTree tree = new BinaryTree();
//		tree.root = new Node(1);
//		tree.root.left = new Node(2);
//		tree.root.right = new Node(3);
//		tree.root.left.left = new Node(4);
//		tree.root.left.right = new Node(5);
//
//		Console.WriteLine("Preorder traversal "
//		+ "of binary tree is ");
//		tree.printPreorder();
//
//		Console.WriteLine("\nInorder traversal "
//		+ "of binary tree is ");
//		tree.printInorder();
//
//		Console.WriteLine("\nPostorder traversal "
//		+ "of binary tree is ");
//		tree.printPostorder();
//	}
//
//	public class Node {
//		public int key;
//		public Node left, right;
//
//		public Node(int item)
//		{
//			key = item;
//			left = right = null;
//		}
//	}
//
//	public class BinaryTree
//	{
//		// Root of Binary Tree
//		public Node root;
//
//		public BinaryTree() { root = null; }
//
//		/* Given a binary tree, print
//		   its nodes according to the
//		   "bottom-up" postorder traversal. */
//		void printPostorder(Node node)
//		{
//			if (node == null)
//			return;
//
//			// first recur on left subtree
//			printPostorder(node.left);
//
//			// then recur on right subtree
//			printPostorder(node.right);
//
//			// now deal with the node
//			Console.Write(node.key + " ");
//		}
//
//		/* Given a binary tree, print
//	    its nodes in inorder*/
//		void printInorder(Node node)
//		{
//			if (node == null)
//				return;
//
//			/* first recur on left child */
//			printInorder(node.left);
//
//			/* then print the data of node */
//			Console.Write(node.key + " ");
//
//			/* now recur on right child */
//			printInorder(node.right);
//		}
//
//		/* Given a binary tree, print
//	   its nodes in preorder*/
//		void printPreorder(Node node)
//		{
//			if (node == null)
//				return;
//
//			/* first print data of node */
//			Console.Write(node.key + " ");
//
//			/* then recur on left subtree */
//			printPreorder(node.left);
//
//			/* now recur on right subtree */
//			printPreorder(node.right);
//		}
//
//		// Wrappers over above recursive functions
//		public void printPostorder() { printPostorder(root); }
//		public void printInorder() { printInorder(root); }
//		public void printPreorder() { printPreorder(root); }
//	}
//}

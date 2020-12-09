package main

/*
Given a binary tree

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

Follow up:

You may only use constant extra space.
Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.

Example 1:

Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

Constraints:

The number of nodes in the given tree is less than 6000.
-100 <= node.val <= 100
*/

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) != 0 {
		newQueue := make([]*Node, 0)

		preIndex := -1
		for i, v := range queue {
			if v != nil {
				preIndex = i
				break
			}
		}
		if preIndex == -1 {
			break
		}

		pre := queue[preIndex]
		newQueue = append(newQueue, pre.Left)
		newQueue = append(newQueue, pre.Right)

		for _, v := range queue[preIndex+1:] {
			if v == nil {
				continue
			}
			pre.Next = v
			pre = v
			newQueue = append(newQueue, pre.Left)
			newQueue = append(newQueue, pre.Right)
		}
		queue = newQueue
	}
	return root
}

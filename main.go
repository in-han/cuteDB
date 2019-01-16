package main

/**
* Insertion Algorithm
1. It will begin from root and value will always be inserted into a leaf node
2. Insert Function Begin
3. If current node is leaf node, then return pick current node, Current Node Insertion Algorithm
    1. This section gives 3 outputs, 1 middle element and 2 child nodes or null,null,null
    2. Insert into the current node
    3. If its full, then sort it and make two new child nodes without the middle node ( NODE CREATION WILL ONLY TAKE PLACE HERE)
    4. take out the middle element along with the two child nodes,  Leaf Splitting no children Algorithm:
        1. Pick middle element by using length of array/2, lets say its index i
        2. Club all elements from 0 to i-1, and i+1 to len(array) and create new seperate nodes by inserting these 2 arrays into the respective keys[] of respective nodes
        3. Since the current node is a leaf node, we do not need to worry about its children and we can leave them to be null for both
        4. return middle,leftNode,rightNode
    5. If its not full, then return null,null,null
4. If this is not a leaf node, then find out the proper child node, Child Node Searching Algorithm:
    1. Input : Value to be inserted, the current Node. Output : Pointer to the childnode
    2. Since the list of values/elements is sorted, perform a binary or linear search to find the first element greater than the value to be inserted, if such an element is found, return pointer at position i, else return last pointer ( ie. the last pointer)
5. After getting the pointer to that element, call insert function Step 2 on that node RECURSIVELY ONLY HERE
6. If we get output from child node insert function Step 2, then this means that we have to insert the middle element received and accomodate the 2 pointers in the current node as well  discarding the old pointer ( NODE DESTRUCTION WILL ONLY TAKE PLACE HERE )
    1. If we got null as output then do nothing, else
    2. Insert into current Node, Popped up element and two child pointers insertion algorithm, Popped Up Joining Algorithm:
        1. Insert element and sort the array
        2. Now we need to discard 1 child pointer and insert 2 child pointers, Child Pointer Manipulation Algorithm :
        3. Find index of inserted element in array, lets say that it is i
        4. Now in the child pointer array, insert the left and right pointers at ith and i+1 th index
    3. If its full, sort it and make two new child nodes, Leaf Splitting with children Algorithm:
        1. Pick middle element by using length of array/2, lets say its index i (Same as 3.4.1)
        2. Club all elements from 0 to i-1, and i+1 to len(lkeys array) and create new seperate nodes by inserting these 2 arrays into the respective keys[] of respective nodes (Same as 3.4.2)
        3. For children[], split the current node's children array into 2 parts, part1 will be from 0 to i, and part 2 will be from i+1 to len(children array), and insert them into leftNode children, and rightNode children
        4. If current node is not the root node return middle,leftNode,rightNode
        5. else if current node == rootNode, Root Node Splitting Algorithm:
            1. Create a new node with elements array as keys[0] = middle
            2. children[0]=leftNode and children[1]=rightNode
            3. Set btree.root=new node
            4. return null,null,null

*/

const maxLeafSize = 2
const maxchildren = maxLeafSize + 1

type Btree struct {
	root *Node
}

type Node struct {
	keys     []int64
	children []*Node
}

func (btree *Btree) isRootNode(node *Node) bool {
	return btree.root == node
}

//Todo: Sorting algorithm
func (n *Node) sort() {

}

func (n *Node) isLeaf() bool {
	return true
}

func (n *Node) addElement(element int64) int {

	// n.keys = append(n.keys, value)
	// n.sort()
	return 0
}

func (n *Node) hasOverFlown() bool {
	return true
}

func (n *Node) getElements() []int64 {
	return n.keys
}

func (n *Node) getElementAtIndex(index int) int64 {
	return n.keys[index]
}

func (n *Node) getChildAtIndex(index int) *Node {
	return n.children[index]
}

func (n *Node) setChildAtIndex(index int, childNode *Node) {
	n.children[index] = childNode
}

func (n *Node) getLastChildNode() *Node {
	return n.children[len(n.children)-1]
}

func (n *Node) getChildNodes() []*Node {
	return n.children
}

func (n *Node) SplitNodeWithNoChildren() (int64, *Node, *Node) {
	/**
		LEAF SPLITTING WITHOUT CHILDREN ALGORITHM
				If its full, then  make two new child nodes without the middle node ( NODE CREATION WILL ONLY TAKE PLACE HERE)
	    		Take out the middle element along with the two child nodes,  Leaf Splitting no children Algorithm:
	        	1. Pick middle element by using length of array/2, lets say its index i
	        	2. Club all elements from 0 to i-1, and i+1 to len(array) and create new seperate nodes by inserting these 2 arrays into the respective keys[] of respective nodes
	        	3. Since the current node is a leaf node, we do not need to worry about its children and we can leave them to be null for both
	        	4. return middle,leftNode,rightNode
	*/
	elements := n.getElements()
	midIndex := len(elements) / 2
	middle := elements[midIndex]

	// Now lets split elements array into 2 as we are splitting this node
	elements1 := elements[0:midIndex]
	elements2 := elements[midIndex+1 : len(elements)]

	// Now lets construct new Nodes from these 2 element arrays
	leftNode := NewNodeWithoutChildren(elements1)
	rightNode := NewNodeWithoutChildren(elements2)

	return middle, leftNode, rightNode
}

func (n *Node) SplitNodeWithChildren() (int64, *Node, *Node) {
	/**
		NON-LEAF NODE SPLITTING ALGORITHM WITH CHILDREN MANIPULATION
		If its full, sort it and make two new child nodes, Leaf Splitting with children Algorithm:
	        1. Pick middle element by using length of array/2, lets say its index i (Same as 3.4.1)
			2. Club all elements from 0 to i-1, and i+1 to len(lkeys array) and create new seperate nodes
			   by inserting these 2 arrays into the respective keys[] of respective nodes (Same as 3.4.2)
			3. For children[], split the current node's children array into 2 parts, part1 will be
			   from 0 to i, and part 2 will be from i+1 to len(children array), and insert them into
			   leftNode children, and rightNode children

	*/
	elements := n.getElements()
	midIndex := len(elements) / 2
	middle := elements[midIndex]

	// Now lets split elements array into 2 as we are splitting this node
	elements1 := elements[0:midIndex]
	elements2 := elements[midIndex+1 : len(elements)]

	// Lets split the children
	children := n.getChildNodes()
	children1 := children[0 : midIndex+1]
	children2 := children[midIndex+1 : len(children)]

	// Now lets construct new Nodes from these 2 element arrays
	leftNode := NewNodeWithChildren(elements1, children1)
	rightNode := NewNodeWithChildren(elements2, children2)

	return middle, leftNode, rightNode
}

func (n *Node) AddPoppedUpElementIntoCurrentNodeAndUpdateWithNewChildren(element int64, leftNode *Node, rightNode *Node) {
	/**
		POPPED UP JOINING ALGORITHM
			Insert into current Node, Popped up element and two child pointers insertion algorithm, Popped Up Joining Algorithm:
	        1. Insert element and sort the array
	        2. Now we need to discard 1 child pointer and insert 2 child pointers, Child Pointer Manipulation Algorithm :
	        3. Find index of inserted element in array, lets say that it is i
	        4. Now in the child pointer array, insert the left and right pointers at ith and i+1 th index

	*/

	//CHILD POINTER MANIPULATION ALGORITHM
	insertionIndex := n.addElement(element)
	n.setChildAtIndex(insertionIndex, leftNode)
	n.setChildAtIndex(insertionIndex+1, rightNode)
}

// create a
func NewNodeWithoutChildren(elements []int64) *Node {
	return &Node{keys: elements}
}

func NewNodeWithChildren(elements []int64, children []*Node) *Node {
	return &Node{keys: elements, children: children}
}

func NewRootNodeWithSingleElementAndTwoChildren(element int64, child1 *Node, child2 *Node) *Node {

	return nil
}

func (n *Node) GetInsertionChildNodeForElement(element int64) *Node {
	/** CHILD NODE SEARCHING ALGORITHM
		If this is not a leaf node, then find out the proper child node, Child Node Searching Algorithm:
	    1. Input : Value to be inserted, the current Node. Output : Pointer to the childnode
		2. Since the list of values/elements is sorted, perform a binary or linear search to find the
		   first element greater than the value to be inserted, if such an element is found, return pointer at position i, else return last pointer ( ie. the last pointer)
	*/

	for i := 0; i < len(n.getElements()); i++ {
		if element < n.getElementAtIndex(i) {
			return n.getChildAtIndex(i)
		}
	}

	// This means that no element is found with value greater than the element to be inserted
	// so we need to return the last child node
	return n.getLastChildNode()
}

func (n *Node) insert(value int64, btree *Btree) (int64, *Node, *Node) {

	if n.isLeaf() {
		n.addElement(value)
		if !n.hasOverFlown() {
			return -1, nil, nil
		} else {
			// Split the node and return to parent function with pooped up element and left,right nodes
			return n.SplitNodeWithNoChildren()
		}
	} else {
		// Get the child Node for insertion
		childNodeToBeInserted := n.GetInsertionChildNodeForElement(value)
		poppedMiddleElement, leftNode, rightNode := childNodeToBeInserted.insert(value, btree)

		if poppedMiddleElement == -1 {
			// this means element has been inserted into the child and hence we do nothing
			return poppedMiddleElement, leftNode, rightNode
		} else {
			// Insert popped up element into current node along with updating the child pointers
			// with new left and right nodes returned
			n.AddPoppedUpElementIntoCurrentNodeAndUpdateWithNewChildren(poppedMiddleElement, leftNode, rightNode)

			if !n.hasOverFlown() {
				// this means that element has been easily inserted into current parent Node
				// without overflowing
				return -1, nil, nil
			} else {
				// this means that the current parent node has overflown, we need to split this up
				// and move the popped up element upwards if this is not the root
				poppedMiddleElement, leftNode, rightNode := n.SplitNodeWithChildren()

				/**
					If current node is not the root node return middle,leftNode,rightNode
				    else if current node == rootNode, Root Node Splitting Algorithm:
				            1. Create a new node with elements array as keys[0] = middle
				            2. children[0]=leftNode and children[1]=rightNode
				            3. Set btree.root=new node
				            4. return null,null,null
				*/

				if !btree.isRootNode(n) {
					return poppedMiddleElement, leftNode, rightNode
				} else {
					newRootNode := NewNodeWithChildren
				}

			}

		}
	}

	return 0, nil, nil
}

func main() {
}
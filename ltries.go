/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package ltries

// Trie object.
type trieObject struct {
	root *nodeObject
}

// Node object.
type nodeObject struct {
	charCode int
	index    int
	left     *nodeObject
	right    *nodeObject
	medium   *nodeObject
}

// This function insert new value in trie.
func (trie *trieObject) Insert(word string, index int) {
	trie.root = put(trie.root, word, index, 0)
}

// This function put value in trie.
func put(node *nodeObject, word string, index int, digit int) *nodeObject {
	var code = charCode(string(word[digit]))
	if node == nil {
		node = &nodeObject{charCode: code, index: -1}
	}
	if code > node.charCode {
		// Go right.
		node.right = put(node.right, word, index, digit)
	} else if code < node.charCode {
		// Go left.
		node.left = put(node.left, word, index, digit)
	} else if digit < len(word)-1 {
		// Go medium.
		node.medium = put(node.medium, word, index, digit+1)
	} else {
		// Save index.
		node.index = index
	}
	return node
}

// This function return index of given word.
func (trie trieObject) Get(word string) int {
	if len(word) == 0 {
		return -1
	}
	node := find(trie.root, word, 0)
	if node != nil {
		return node.index
	}
	return -1
}

// This function find given word index.
func find(node *nodeObject, word string, digit int) *nodeObject {
	var code = charCode(string(word[digit]))
	if node == nil {
		return nil
	}
	if code > node.charCode {
		// Go right.
		return find(node.right, word, digit)
	} else if code < node.charCode {
		// Go left.
		return find(node.left, word, digit)
	} else if digit < len(word)-1 {
		// Go medium.
		return find(node.medium, word, digit+1)
	} else {
		return node
	}
}

// This function remove given word from trie.
func (trie *trieObject) Remove(word string) {
	trie.root = delete(trie.root, word, 0)
}

// This function delete node.
func delete(node *nodeObject, word string, digit int) *nodeObject {
	var code = charCode(string(word[digit]))
	if node == nil {
		return nil
	}
	if code > node.charCode {
		// Go right.
		node.right = delete(node.right, word, digit)
	} else if code < node.charCode {
		// Go left.
		node.left = delete(node.left, word, digit)
	} else if digit < len(word)-1 {
		// Go medium.
		node.medium = delete(node.medium, word, digit+1)
	} else if node.medium == nil && node.left == nil && node.right == nil {
		// Remove node.
		return nil
	}
	return node
}

// This function return character code.
func charCode(char string) int {
	return int([]rune(char)[0])
}

// This function create new trie.
func Create() trieObject {
	return trieObject{}
}

// TreeNode - узел дерева
type TreeNode struct {
    value interface{}
    left  *TreeNode
    right *TreeNode
}

// FullBinaryTree - полное бинарное дерево
type FullBinaryTree struct {
    root *TreeNode
    size int
}

// NewFullBinaryTree создает новое полное бинарное дерево
func NewFullBinaryTree() *FullBinaryTree {
    return &FullBinaryTree{nil, 0}
}

// Add добавляет элемент в дерево
func (t *FullBinaryTree) Add(value interface{}) {
    newNode := &TreeNode{value: value}
    
    if t.root == nil {
        t.root = newNode
        t.size++
        return
    }
    
    // Используем очередь для поиска места для нового узла
    queue := []*TreeNode{t.root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.left == nil {
            node.left = newNode
            t.size++
            return
        } else if node.right == nil {
            node.right = newNode
            t.size++
            return
        } else {
            queue = append(queue, node.left, node.right)
        }
    }
}

// Find ищет элемент в дереве
func (t *FullBinaryTree) Find(value interface{}) (*TreeNode, bool) {
    if t.root == nil {
        return nil, false
    }
    
    queue := []*TreeNode{t.root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.value == value {
            return node, true
        }
        
        if node.left != nil {
            queue = append(queue, node.left)
        }
        if node.right != nil {
            queue = append(queue, node.right)
        }
    }
    
    return nil, false
}

// Remove удаляет элемент из дерева с сохранением структуры
func (t *FullBinaryTree) Remove(value interface{}) bool {
    if t.root == nil {
        return false
    }
    
    // Находим узел для удаления и самый глубокий правый узел
    var nodeToRemove, deepestNode *TreeNode
    var parentOfDeepest *TreeNode
    
    queue := []*TreeNode{t.root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.value == value {
            nodeToRemove = node
        }
        
        // Обновляем самый глубокий узел
        deepestNode = node
        
        if node.left != nil {
            parentOfDeepest = node
            queue = append(queue, node.left)
        }
        if node.right != nil {
            parentOfDeepest = node
            queue = append(queue, node.right)
        }
    }
    
    if nodeToRemove == nil {
        return false
    }
    
    // Заменяем значение удаляемого узла значением самого глубокого узла
    nodeToRemove.value = deepestNode.value
    
    // Удаляем самый глубокий узел
    if parentOfDeepest != nil {
        if parentOfDeepest.right == deepestNode {
            parentOfDeepest.right = nil
        } else {
            parentOfDeepest.left = nil
        }
    } else {
        t.root = nil
    }
    
    t.size--
    return true
}

// InOrder обход дерева (левый-корень-правый)
func (t *FullBinaryTree) InOrder() []interface{} {
    result := make([]interface{}, 0)
    t.inOrderTraversal(t.root, &result)
    return result
}

func (t *FullBinaryTree) inOrderTraversal(node *TreeNode, result *[]interface{}) {
    if node != nil {
        t.inOrderTraversal(node.left, result)
        *result = append(*result, node.value)
        t.inOrderTraversal(node.right, result)
    }
}

// PreOrder обход дерева (корень-левый-правый)
func (t *FullBinaryTree) PreOrder() []interface{} {
    result := make([]interface{}, 0)
    t.preOrderTraversal(t.root, &result)
    return result
}

func (t *FullBinaryTree) preOrderTraversal(node *TreeNode, result *[]interface{}) {
    if node != nil {
        *result = append(*result, node.value)
        t.preOrderTraversal(node.left, result)
        t.preOrderTraversal(node.right, result)
    }
}

// PostOrder обход дерева (левый-правый-корень)
func (t *FullBinaryTree) PostOrder() []interface{} {
    result := make([]interface{}, 0)
    t.postOrderTraversal(t.root, &result)
    return result
}

func (t *FullBinaryTree) postOrderTraversal(node *TreeNode, result *[]interface{}) {
    if node != nil {
        t.postOrderTraversal(node.left, result)
        t.postOrderTraversal(node.right, result)
        *result = append(*result, node.value)
    }
}

// LevelOrder обход дерева по уровням
func (t *FullBinaryTree) LevelOrder() []interface{} {
    result := make([]interface{}, 0)
    
    if t.root == nil {
        return result
    }
    
    queue := []*TreeNode{t.root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        result = append(result, node.value)
        
        if node.left != nil {
            queue = append(queue, node.left)
        }
        if node.right != nil {
            queue = append(queue, node.right)
        }
    }
    
    return result
}

// Size возвращает количество элементов в дереве
func (t *FullBinaryTree) Size() int {
    return t.size
}

// String возвращает строковое представление дерева
func (t *FullBinaryTree) String() string {
    return fmt.Sprintf("InOrder: %v", t.InOrder())
}

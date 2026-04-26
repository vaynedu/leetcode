package binarytree

import (
	"reflect"
	"testing"
)

func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	// 用切片模拟队列，层层构建
	nodes := make([]*TreeNode, 0, len(vals))
	for _, v := range vals {
		if v == nil {
			nodes = append(nodes, nil)
		} else {
			nodes = append(nodes, &TreeNode{Val: v.(int)})
		}
	}
	// 构建父子关系
	k := 1
	for i := 0; i < len(nodes) && k < len(nodes); i++ {
		if nodes[i] == nil {
			continue
		}
		nodes[i].Left = nodes[k]
		k++
		if k < len(nodes) {
			nodes[i].Right = nodes[k]
			k++
		}
	}
	return nodes[0]
}

//       1
//    2    3
//  4  5  6  7
var tree1 = buildTree([]interface{}{1, 2, 3, 4, 5, 6, 7})

func TestPreorderRec(t *testing.T) {
	var result []int
	Preorder(tree1, &result)
	expected := []int{1, 2, 4, 5, 3, 6, 7}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Preorder: got %v, want %v", result, expected)
	}
}

func TestInorderRec(t *testing.T) {
	var result []int
	Inorder(tree1, &result)
	expected := []int{4, 2, 5, 1, 6, 3, 7}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Inorder: got %v, want %v", result, expected)
	}
}

func TestPostorderRec(t *testing.T) {
	var result []int
	Postorder(tree1, &result)
	expected := []int{4, 5, 2, 6, 7, 3, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Postorder: got %v, want %v", result, expected)
	}
}

func TestPreorderIter(t *testing.T) {
	result := PreorderIter(tree1)
	expected := []int{1, 2, 4, 5, 3, 6, 7}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PreorderIter: got %v, want %v", result, expected)
	}
}

func TestInorderIter(t *testing.T) {
	result := InorderIter(tree1)
	expected := []int{4, 2, 5, 1, 6, 3, 7}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InorderIter: got %v, want %v", result, expected)
	}
}

func TestPostorderIter(t *testing.T) {
	result := PostorderIter(tree1)
	expected := []int{4, 5, 2, 6, 7, 3, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PostorderIter: got %v, want %v", result, expected)
	}
}

func TestEmptyTree(t *testing.T) {
	var result []int
	Preorder(nil, &result)
	if len(result) != 0 {
		t.Errorf("expected empty, got %v", result)
	}
	if PreorderIter(nil) != nil {
		t.Error("PreorderIter(nil) should return nil")
	}
}

package student_test

import (
	"reflect"
	"testing"

	"github.com/01-edu/z01"

	solutions "./solutions"
	student "./student"
)

func parentListRotRight(root *student.TreeNode) string {
	if root == nil {
		return ""
	}

	var parent string

	if root.Parent == nil {
		parent = "nil"
	} else {
		parent = root.Parent.Data
	}

	r := "Node: " + root.Data + " Parent: " + parent + "\n"
	r += parentListRotRight(root.Left) + parentListRotRight(root.Right)
	return r
}

func CompareTreesRotRight(t *testing.T, fn interface{}, root, rootA *solutions.TreeNode, rootAS *student.TreeNode) {
	parentSol := solutions.ParentList(rootA)
	parentStu := parentListRotRight(rootAS)
	solTree := solutions.FormatTree(root)
	if parentSol != parentStu {
		t.Fatalf("Tree:\n%s\nExpected\n%s instead of\n%s\n", solTree, parentSol, parentStu)
	}
}

func FormatTree_rotright(root *student.TreeNode) string {
	if root == nil {
		return ""
	}
	res := root.Data + "\n"
	res += formatSubTree_rotright(root, "")
	return res
}

func formatSubTree_rotright(root *student.TreeNode, prefix string) string {
	if root == nil {
		return ""
	}

	var res string

	hasLeft := root.Left != nil
	hasRight := root.Right != nil

	if !hasLeft && !hasRight {
		return res
	}

	res += prefix
	if hasLeft && hasRight {
		res += "├── "
	}

	if !hasLeft && hasRight {
		res += "└── "
	}

	if hasRight {
		printStrand := (hasLeft && hasRight && (root.Right.Right != nil || root.Right.Left != nil))
		newPrefix := prefix
		if printStrand {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		res += root.Right.Data + "\n"
		res += formatSubTree_rotright(root.Right, newPrefix)
	}

	if hasLeft {
		if hasRight {
			res += prefix
		}
		res += "└── " + root.Left.Data + "\n"
		res += formatSubTree_rotright(root.Left, prefix+"    ")
	}
	return res
}

func errorMessage_rotright(t *testing.T, fn interface{}, rootOr, root *solutions.TreeNode, rootS *student.TreeNode) {
	t.Fatalf("%s(\n%s\n)\n == \n%s instead of \n%s\n",
		z01.NameOfFunc(fn),
		solutions.FormatTree(rootOr),
		FormatTree_rotright(rootS),
		solutions.FormatTree(root),
	)
}

func CompareNode_rotright(t *testing.T, fn interface{}, rootOr, root *solutions.TreeNode, rootS *student.TreeNode) {
	solTree := solutions.FormatTree(root)
	stuTree := FormatTree_rotright(rootS)

	if solTree != stuTree {
		errorMessage_rotright(t, fn, rootOr, root, rootS)
	}
}

func CompareReturn_rotright(t *testing.T, fn1, fn2, rootOr, arg1, arg2 interface{}) {
	arar1 := []interface{}{arg1}
	arar2 := []interface{}{arg2}

	out1 := z01.Monitor(fn1, arar1)
	out2 := z01.Monitor(fn2, arar2)

	for i, v := range out1.Results {
		switch str := v.(type) {
		case *solutions.TreeNode:
			CompareNode_rotright(t, fn1, rootOr.(*solutions.TreeNode), str, out2.Results[i].(*student.TreeNode))
		default:
			if !reflect.DeepEqual(str, out2.Results[i]) {
				t.Fatalf("%s(%s) == %s instead of %s\n",
					z01.NameOfFunc(fn1),
					z01.Format(arg1),
					z01.Format(out2.Results...),
					z01.Format(out1.Results...),
				)
			}
		}
	}
}

func TestBTreeRotateRight(t *testing.T) {
	root := &solutions.TreeNode{Data: "04"}
	rootOr := &solutions.TreeNode{Data: "04"}
	rootS := &student.TreeNode{Data: "04"}

	ins := []string{"01", "07", "05", "12", "02", "03", "10"}

	for _, v := range ins {
		root = solutions.BTreeInsertData(root, v)
		rootS = student.BTreeInsertData(rootS, v)
		rootOr = solutions.BTreeInsertData(rootOr, v)
	}
	fn := interface{}(solutions.BTreeRotateRight)
	CompareReturn_rotright(t, fn, student.BTreeRotateRight, rootOr, root, rootS)
	CompareTreesRotRight(t, fn, rootOr, root, rootS)
}
package binary_tree

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

var comparer ComparerFunc[int] = func(a, b int) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	}

	return -1
}

func TestBinaryTree(t *testing.T) {
	tree := setup(t)

	t.Run("Insert into tree", func(t *testing.T) {
		require.NoError(t, errors.New("KEK"), "SHPEK")
	})

	t.Run("Found in tree", func(t *testing.T) {
		values := []int{5, 35, 30, 20, 10, 15, 25}

		var result bool
		for _, v := range values {
			result = tree.Find(v)
			if !result {
				break
			}
		}

		require.True(t, result, "There is no provided element in the tree")
	})

	t.Run("Not found in tree", func(t *testing.T) {
		values := []int{0, 3, 8, 11, 19, 24, 29, 31, 40}

		var result bool
		for _, v := range values {
			result = tree.Find(v)
			if result {
				break
			}
		}

		require.False(t, result, "The provided element must not be in the tree")
	})

	t.Run("Remove left", func(t *testing.T) {
		isRemoved := tree.Remove(10)
		//require.Thrue(t, isRemoved)
		//tree.Traverse()

		isRemoved = tree.Remove(30)
		require.True(t, isRemoved)
	})
}

func setup(t *testing.T) *BinaryTree[int] {
	t.Helper()

	tree := NewBinaryTree[int](comparer)

	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(15)
	tree.Insert(35)
	tree.Insert(5)
	tree.Insert(25)
	tree.Insert(1)
	tree.Insert(7)
	tree.Insert(13)
	tree.Insert(18)
	tree.Insert(22)
	tree.Insert(27)
	tree.Insert(31)
	tree.Insert(40)

	//tree.Insert(18)
	//tree.Insert(20)
	//tree.Insert(25)
	//tree.Insert(22)
	//tree.Insert(30)
	//tree.Insert(40)
	//tree.Insert(35)
	//tree.Insert(42)
	//tree.Insert(50)
	//tree.Insert(16)
	//tree.Insert(21)

	require.Equal(t, 3, tree.root.GetSubtreeHeight())

	//    18
	// 16    20
	//          25
	//        22  30
	//      21      40
	//            35  42
	//                   50

	//		       		            20
	//			    10     	     	                  30
	//        5           15  	         	 25               35
	//   1       7      13    18          22     27         31     40

	//		       		            20
	//			    5 (3) 	     	                  30
	//        1          7 (2) 	         	 25               35
	//                      15          22     27         31     40
	//                   13    18
	return tree
}

//          7
//     <>       15
//           13   18

//                  15
//          7              18
//     <>       13
//

//                5
//      1                  15
//                  7           18
//              <>      13

// пересчет дерева при инсерте
// пересчет дерева при удалении <------
// балансировки

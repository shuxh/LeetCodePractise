/*
题目要求：
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

解题思路：
1. 增加上下左右四个辅助移动的边界条件。
2. 每个方向迭代完成后，更新辅助边界条件；比如到达最右边的时候，代表当前行已经遍历完成，更新startRow++;
   到达最下端的时候，代表当前列遍历完成；更新endCol--;
3. 当辅助上下或左右辅助边界条件大小不满足时，停止。
*/

package spiralOrder

func SpiralOrder(matrix [][]int) []int {
	var res = make([]int, 0)

	if matrix == nil || len(matrix) == 0 {
		return res
	}

	nrows := len(matrix)
	ncols := len(matrix[0])

	startRow := 0       // 辅助起始行号
	endRow := nrows - 1 // 辅助终止行号

	startCol := 0       // 辅助起始列号
	endCol := ncols - 1 // 辅助终止列号

	for {
		// 从左至右遍历行
		for i := startCol; i <= endCol; i++ {
			res = append(res, matrix[startRow][i])
		}
		// 更新辅助起始行号
		startRow++
		if startRow > endRow {
			break
		}

		// 从上至下遍历列
		for i := startRow; i <= endRow; i++ {
			res = append(res, matrix[i][endCol])
		}
		// 更新辅助终止列号
		endCol--
		if endCol < startCol {
			break
		}

		// 从右向左遍历
		for i := endCol; i >= startCol; i-- {
			res = append(res, matrix[endRow][i])
		}
		// 更新辅助终止行号
		endRow--
		if endRow < startRow {
			break
		}
		// 从下至上遍历
		for i := endRow; i >= startRow; i-- {
			res = append(res, matrix[i][startCol])
		}
		// 更新辅助起始的列号
		startCol++

		if startCol > endCol {
			break
		}
	}

	return res
}

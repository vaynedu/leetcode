package backtrack

//// 伪码
//var res [][]int
//func backTrack(路径，选择列表){
//	if 满足结束条件{ // 终止条件
//		res = append(res, 路径)  // 存放结果
//		return
//	}
//	for _,选择 := range 选择列表{ // 选择：本层集合中元素（树中节点孩子的数量就是集合的大小）
//		做选择  // 处理节点
//		backTrack(路径，选择列表)  // 递归
//		撤销选择  // 回溯，撤销处理结果
//	}
//}

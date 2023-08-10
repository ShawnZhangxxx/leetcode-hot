/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 10:50:18
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 10:59:52
 * @Description:
 */

package main

func leastInterval(tasks []byte, n int) int {
	taskCount := make([]int, 26)
	maxCount := 0
	maxCountTasks := 0
	for _, task := range tasks {
		taskCount[task-'A']++
		if taskCount[task-'A'] > maxCount {
			maxCount = taskCount[task-'A']
			maxCountTasks = 1
		} else if taskCount[task-'A'] == maxCount {
			maxCountTasks++
		}
	}

	// 计算最大任务数和空闲时间数
	interval := (maxCount-1)*(n+1) + maxCountTasks
	// 取任务数和时间数两者中的较大值
	if len(tasks) > interval {
		return len(tasks)
	}

	return interval
}

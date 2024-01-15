package class047;

这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。

示例 1：

输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]
示例 2：

输入：bookings = [[1,2,10],[2,2,15]], n = 2
输出：[10,25]
解释：
航班编号        1   2
预订记录 1 ：   10  10
预订记录 2 ：       15
总座位数：      10  25
因此，answer = [10,25]

// 测试链接 : https://leetcode.cn/problems/corporate-flight-bookings/ 1109
public class Code01_CorporateFlightBookings {

	// bookings
	// [1,5,6]
	// [2,9,3]
	// ...
	//不支持变操作边查询,要查询的用线段树结构
	public static int[] corpFlightBookings(int[][] bookings, int n) {
		int[] cnt = new int[n + 2]; //要修改r+1 位置,省的考虑边界问题 ,从1开始
		// 设置差分数组，每一个操作对应两个设置
		for (int[] book : bookings) {
			cnt[book[0]] += book[2];
			cnt[book[1] + 1] -= book[2];
		}
		// 加工前缀和 通过差分数组得到前缀和
		for (int i = 1; i < cnt.length; i++) {
			cnt[i] += cnt[i - 1];
		}
		int[] ans = new int[n];
		for (int i = 0; i < n; i++) {
			ans[i] = cnt[i + 1];
		}
		return ans;
	}

}
/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main




func main() {

	//位图相比于hash表比较省空间,条件是连续的范围,初始化就得好

	//位图 创建 set = new int[(n+31)/32] 个数组,最多3个,32位一个整数

	//加入一个数 add set[num/32] |= 1 << (num % 32) 将这位置为1 |= 或进去将改位修改

	//删除一个数 remove set[n/32] &= ~(1 << num%32) 与上 1101111 推荐这个方法不用判断之前是否是0或者1 ,异或也能做不推荐

	//改变一个数 reverse flip 原来1->0 0->1   set[n/32] ^= (1 << (num % 32))

	//查询一个数 contains  (set[n/32] >> (num % 32) & 1) == 1




}



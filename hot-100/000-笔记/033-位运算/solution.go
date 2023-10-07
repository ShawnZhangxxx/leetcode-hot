/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main




func main() {


	//二进制实现加法
	//先是^运算 ,无进位相加
	//再是进行&运算,之后左移一位得到进位信息
	//同样 循环无进位相加,再得到进位信息,直到进位信息为0才停止.
	//add(int a,int b) {  位运算实现加法
	//	int ans = 0
	//	while(b != 0 ){
	//		ans = a ^ b        无进位加法
	//		b = (a & b) << 1   进位信息
	//		a = ans
	//	}
	//	return ans
	//}

	//减法 a-b = a + (~b + 1)


	//乘法 和十进制的乘法一样 扣掉1位b,如果最右边是1,加一个a左移1位
	//multiple(int a ,int b){
	//	int ans = 0
	//	while(b != 0 ){
	//		if(b & 1 != 0){
	//			ans = add(ans , a);
	//		}
	//		a <<= 1     //可能越界,但是计算机不管,越界是程序员控制的
	//		b >>>= 1     b = b >>> 1
	//	}
	//	return ans
	//}

	//除法 由最高位不断/2^30 右移30位,余数直接扔掉了
	//280 = 25 * 2^3 + 25 * 2^1 + 25 * 2^0
	//280 / 25 = 1 0 1 1
	//           3   1 0  位置

	//div(int a ,int b){
	//整数最小值还得做特殊处理 不能解决相反数, 当a为最小值, 要这么算才能不越界(a+b)/b -1
	//	int x = a < 0 ? neg(a) : a;
	//	int y = b < 0 ? neg(b) : b;
	//	int ans = 0
	//	for (int i = 30; i>= 0; i= minus(i,1)){
	//		if( (x>>i) >= y){
	//			ans != (1 << i)
	//			x = minus(x,y <<i)
	//		}
	//	}
	//	return a<0 ^ b < 0 ? neg(ans):ans
	//}

}



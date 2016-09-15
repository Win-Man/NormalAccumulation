[继续(3n+1)猜想 ](https://www.patest.cn/contests/pat-b-practise/1005)

**解题思路：**
用一个数组标记进行猜想验证时能产生的数，然后对输入的的数进行遍历，如果能在标记数组中被找到，说明能被覆盖，不是关键数。

**AC code:**

``` java
import java.util.Arrays;
import java.util.Scanner;


public class Main {

	public static void main(String[] args){
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();
		int[] arr = new int[n];
		int[] flag = new int[105];
		for(int i = 0;i < n;i++){
			arr[i] = sc.nextInt();
		}
		for(int i = 0;i < n;i++){
			int tmp = arr[i];
			while(true){
				tmp = tmp%2 == 0? tmp/2:(3*tmp+1)/2;
				if(tmp == 1){
					break;
				}
				/**
				 * 1<n<=100；所以能覆盖的数如果超过105可以不用考虑
				 */
				if(tmp < 105){
					flag[tmp] = 1;
				}
			}
			
		}
		Arrays.sort(arr);
		int count = 0;
		for(int i = n-1;i >= 0;i--){
			/**
			 * 判断这个数能否被覆盖
			 */
			if(flag[arr[i]] != 1){
				if(count == 0){
					System.out.print(arr[i]);
					count++;
				}else{
					System.out.print(" "+arr[i]);
				}
			}
		}
		System.out.println();
	}
}
```
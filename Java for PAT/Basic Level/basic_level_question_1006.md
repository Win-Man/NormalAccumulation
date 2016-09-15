[换个格式输出整数](https://www.patest.cn/contests/pat-b-practise/1006)

**解题思路：**


**AC code:**

``` java
import java.util.Scanner;


public class Main {
	
	public static void main(String[] args){
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();
		int ge = n % 10;
		int shi = (n / 10) % 10;
		int bai = n / 100;
		for(int i = 0;i < bai; i++){
			System.out.print('B');
		}
		for(int i = 0;i < shi; i++){
			System.out.print('S');
		}
		for(int i = 1;i <= ge; i++){
			System.out.print(i);
		}
		System.out.println();
	}

}
```
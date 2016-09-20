[素数对猜想](https://www.patest.cn/contests/pat-b-practise/1007)

**解题思路：**


**AC code:**

``` java
import java.util.Scanner;

public class Main {
	
	public static boolean isPrimer(int n){
		if(n <= 1){
			return false;
		}
		for(int i = 2;i*i < n+1;i++){
			if(n % i == 0){
				return false;
			}
		}
		return true;
	}
	
	public static void main(String[] args){
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();
		int[] flag = new int[100005];
		for(int i =0;i < flag.length;i++){
			if(isPrimer(i)){
				flag[i] = 1;
			}else{
				flag[i] = 0;
			}
		}
		int count = 0;
		for(int i = 3;i <= n-2;i++){
			if(flag[i] == 1 && flag[i+2] == 1){
				count ++;
			}
		}
		System.out.println(count);
	}

}

```
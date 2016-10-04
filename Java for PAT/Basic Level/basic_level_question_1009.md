[说反话](https://www.patest.cn/contests/pat-b-practise/1009)

**解题思路：**


**AC code:**

``` java
import java.util.*;

public class Main {
	
	public static void main(String[] args){
		Scanner sc = new Scanner(System.in);
		String str = null;
		str = sc.nextLine();
		String[] arr = str.split(" ");
		if(arr.length >= 1){
			System.out.print(arr[arr.length-1]);
		}
		for(int i = arr.length-2;i >= 0;i--){
			System.out.print(" "+arr[i]);
		}
		System.out.println();
	}

}

```
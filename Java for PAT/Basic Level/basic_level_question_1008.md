[数组元素循环右移问题](https://www.patest.cn/contests/pat-b-practise/1008)

**解题思路：**


**AC code:**

``` java
public class Main {
	
	public static void main(String[] args){
		Scanner sc = new Scanner(System.in);
		int n = sc.nextInt();
		int m = sc.nextInt();
		int[] aa = new int[n];
		for(int i = 0;i < n;i++){
			aa[i] = sc.nextInt();
		}
		int[] a = new int[n];
		for(int i =0;i < n;i++ ){
			a[(i+m)%n] = aa[i];
		}
		
		System.out.print(a[0]);
		for(int i = 1; i < n;i++){
			System.out.print(" "+a[i]);
		}
		System.out.println();
	}

}

```
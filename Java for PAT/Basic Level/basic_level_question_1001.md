[题目链接](https://www.patest.cn/contests/pat-b-practise/1001)

**AC code:**

``` java
import java.util.Scanner;

/**
 * Created by sg on 2016/9/12.
 */
public class Main {
    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int count = 0;
        while(n != 1){
            count++;
            if(n % 2 == 0){
                n = n >> 1;
            }else{
                n = ((n * 3) + 1) >> 1;
            }
        }
        System.out.println(count);
    }
}
```
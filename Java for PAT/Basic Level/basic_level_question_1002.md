[写出这个数 ](https://www.patest.cn/contests/pat-b-practise/1002)

**AC code:**

``` java
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

/**
 * Created by sg on 2016/9/13.
 */
public class Main {
    public static void main(String[] args){
        String[] pinyi = {"ling","yi","er","san","si","wu","liu","qi","ba","jiu"};
        Scanner sc = new Scanner(System.in);
        String n = sc.next();
        int strSum = 0;
        for(int i = 0;i < n.length();i++){
            strSum += (n.charAt(i)-'0');
        }
        List<String> result = new ArrayList<String>();
        while(strSum > 0){
            result.add(pinyi[strSum % 10]);
            strSum /= 10;
        }
        for(int i =0;i < result.size();i++){
            System.out.print(result.get(result.size()-1-i));
            if(i+1 != result.size()){
                System.out.print(" ");
            }
        }
        System.out.println();
    }
}
```
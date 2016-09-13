[我要通过！](https://www.patest.cn/contests/pat-b-practise/1003)

**解题思路：**
这道题最主要是的是对第三条规则的理解
>  如果 aPbTc 是正确的，那么 aPbATca 也是正确的，其中 a, b, c 均或者是空字符串，或者是仅由字母 A 组成的字符串。

如果 aPbTc正确 按照第二条规则，a=c b=A,所以cPATc正确，则cPAATcc正确，则cPAAATccc正确，则cPAAAATcccc正确····可以看到P与T之间的A的个数等于T之后的c的个数，所以得出的规则是P之前的字符长度乘以P、T之间A的个数等于T之后字符长度

**AC code:**

``` java
import java.util.Scanner;
/**
 * Created by sg on 2016/9/13.
 */
public class Main {

    public static boolean checkPAT(String str){
        boolean result = true;
        int pPos = -1;
        int tPos = -1;
        for(int i =0;i < str.length();i++){
            if(str.charAt(i) != 'P' && str.charAt(i) != 'A' && str.charAt(i) != 'T'){
                return false;
            }
            if(str.charAt(i) == 'P'){
                if(pPos != -1){
                    return false;
                }else{
                    pPos = i;
                }
            }
            if(str.charAt(i) == 'T'){
                if(tPos != -1){
                    return false;
                }else{
                    tPos = i;
                }
            }
        }
        if(pPos == -1 || tPos == -1){
            return false;
        }
        if(tPos - pPos <= 1){
            return false;
        }
        if(tPos <= pPos){
            result = false;
        }
        if((tPos - pPos - 1)*pPos != str.length() - tPos - 1){
            return false;
        }
        return result;
    }

    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        for(int i = 0;i < n;i++){
            String str = sc.next();
            if(checkPAT(str)){
                System.out.println("YES");
            }else{
                System.out.println("NO");
            }
        }
    }
}
```
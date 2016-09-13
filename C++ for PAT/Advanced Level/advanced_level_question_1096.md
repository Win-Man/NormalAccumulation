[Consecutive Factors](https://www.patest.cn/contests/pat-a-practise/1096)


**解题思路：**
不解的是用i*i<=n代替i

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<math.h>
using namespace std;
int main(){
    for (int n; scanf("%d", &n) != EOF;){
        int maxCount = 0;
        int start = 0;
        int haha = sqrt((double)n);
        for (int i = 2; i<haha; i++){
            int num = n;
            for (int j = i; j < n; j++){
                if (num%j == 0){
                    num /= j;
                }
                else{
                    if (j - i > maxCount){
                        maxCount = j - i;
                        start = i;
                    }
                    break;
                }
            }
        }
        if (maxCount == 0){
            printf("1\n%d\n", n);
        }
        else{
            printf("%d\n", maxCount);
            printf("%d", start);
            for (int i = 1; i < maxCount; i++){
                printf("*%d", start + i);
            }
            printf("\n");
        }
    }
    return 0;
}
```
[Find Coins](https://www.patest.cn/contests/pat-a-practise/1048)


**解题思路：**
真的是PAT虐我千百遍，我待PAT如初恋，很简单的一道题，开一个面值的数组，已面值为数组下标，存放该面值的硬币的数量，然后从i=1开始遍历，看m-i的面值的硬币是否存在，如果存在，直接输出就好了，需要考虑输出两个面值一样的硬币的情况。 
！！！重点来了！！！ 
题目中说

>The second line contains N face values of the coins, which are all positive numbers no more than 500.

每个硬币的面额不超过500，那我就开了一个505的数组，我想这应该够了吧，结果测试点3死活过不了，数组开城1001才过。PAT你也真是够了，能老老实实按照题目意思来吗。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
using namespace std;
int main(){
    for (int n, m; scanf("%d%d", &n, &m) != EOF;){
        int coins[1001] = { 0 };
        for (int i = 0; i < n; i++){
            int temp;
            scanf("%d", &temp);
            coins[temp]++;
        }
        int flag = 0;
        for (int i = 1; i < m; i++){
            if (coins[i]>0){
                coins[i]--;
                if (coins[m - i]>0){
                    flag = 1;
                    printf("%d %d\n", i, m-i);
                    break;
                }
                coins[i]++;
            }
        }
        if (!flag){
            printf("No Solution\n");
        }
    }
}
```
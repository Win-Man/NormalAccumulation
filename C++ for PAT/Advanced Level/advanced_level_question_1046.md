[Shortest Distance](https://www.patest.cn/contests/pat-a-practise/1046)


**解题思路：**
如果只是简单的记录各个点之间的距离，每次询问的时候进行计算的话，最后一个测试点会运行超时，于是就改成数组中记录的是，从一号点到当前点的总距离，查询的时候直接将距离相减就是正向的两个点之间的距离，反向的距离用总的环的距离减去正向两个点之间的距离就可，输出较小的距离。

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<stdio.h>
using namespace std;
int main(){
    for (int n; scanf("%d", &n) != EOF;){
        int *arr = new int[n+1];
        int total = 0;
        for (int i = 1; i <= n; i++){
            arr[i] = total;
            int temp;
            scanf("%d", &temp);
            total += temp;
        }
        int m;
        scanf("%d", &m);
        while (m--){
            int begin;
            int end;
            scanf("%d%d", &begin, &end);
            if (begin > end){
                begin = begin^end;
                end = begin^end;
                begin = begin^end;
            }
            int distance = arr[end] - arr[begin];
            if (distance > total - distance){
                printf("%d\n", total - distance);
            }
            else{
                printf("%d\n", distance);
            }
        }
    }
}
```
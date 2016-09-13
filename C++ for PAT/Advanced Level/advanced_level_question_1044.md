[Shopping in Mars](https://www.patest.cn/contests/pat-a-practise/1044)


**解题思路：**
看到题目很顺手就写下来了，一提交，两个运行超时，一个内存超限，果然题目没那么简单，想了各种优化，没什么想法，网上找了一下都说是要二分，不想用二分，于是继续找，终于找到一篇文章是从n*n的复杂度上进行优化，用到了几个点是我没考虑到的。 
1.当i~j的区间上的和等于m之后，i+1~j区间上的情况就不用考虑了 
2.对于找不到等于m的情况，在找到某个区间i~j之和大于m之后，就可以跳出j的循环，从i+1为起点的区间再开始寻找，因为之后的区间i~j+1……..i~n累加的和更大，不用考虑。 
再加上c的输入输出的效率优化已经足够了。

自己写的程序一开始有一个情况没考虑到，导致优化之后最后一个点一直错误，就是在没有找到区间和等于m的情况之前就发现了区间和大于m的情况，使对最后输出的结果需要筛选一下。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<algorithm>
using namespace std;
struct pairs{
    int begin;
    int end;
    int num;
};
bool cmp(const pairs p1, const pairs p2){
    if (p1.num < p2.num){
        return true;
    }
    else if (p1.num == p2.num){
        if (p1.begin < p2.begin){
            return true;
        }
        else{
            return false;
        }
    }
    else{
        return false;
    }
}
int main(){
    for (int n, m; scanf("%d%d", &n, &m) != EOF;){
        int *arr = new int[n + 1];
        int i;
        int j;
        arr[0] = 0;
        int total = 0;
        for (i = 1; i <= n; i++){
            int temp;
            scanf("%d", &temp);
            total += temp;
            arr[i] = total;
        }
        vector<pairs>vec;
        int flag = 0;
        int jmp = 0;
        for ( i = 1; i <= n; i++){
            if (flag && i <= jmp){
                j = jmp+1;
            }
            else{
                jmp = 0;
                j = i;
            }
            for (; j <= n; j++){
                if (arr[j] - arr[i - 1] == m){
                    pairs p;
                    p.begin = i;
                    p.end = j;
                    p.num = arr[j] - arr[i - 1];
                    vec.push_back(p);
                    flag = 1;
                    jmp = j;
                    break;
                }
                else if (arr[j] - arr[i - 1] > m&& !flag){
                    pairs p;
                    p.begin = i;
                    p.end = j;
                    p.num = arr[j] - arr[i - 1];
                    vec.push_back(p);
                    //测试点3是 >m的情况
                    break;
                }
            }
        }
        sort(vec.begin(), vec.end(),cmp);
        if (vec.size()){
            int check = vec[0].num;
            for (i = 0; i < vec.size(); i++){
                if (vec[i].num == check){
                    printf("%d-%d\n", vec[i].begin, vec[i].end);
                }
                else{
                    break;
                }
            }
        }
    }
    return 0;
}
```
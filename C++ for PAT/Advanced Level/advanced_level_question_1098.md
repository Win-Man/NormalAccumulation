[Insertion or Heap Sort](https://www.patest.cn/contests/pat-a-practise/1098)


**解题思路：**
与插入与归并那题一样，先模拟插入排序，找中间过程有没有一样的，如果一直找不到就表示是堆排序的，未排序的部分中第一个一定是未排序的序列中最大的，堆是一棵完全二叉树，所以用数组就能完成的堆的创建，每次将未排序序列中的第一个元素与最后一个元素进行交换，然后对剩余的未排序序列进行堆调整就好。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<algorithm>
using namespace std;
int main(){
    for (int n; scanf("%d", &n) != EOF;){
        vector<int>in(n);
        vector<int>comp(n);
        for (int i = 0; i < n; i++){
            scanf("%d", &in[i]);
        }
        for (int i = 0; i < n; i++){
            scanf("%d", &comp[i]);
        }
        vector<int>change(in);
        //判断是否是快速排序
        int flag = 1;
        for (int i = 2; i < n; i++){
            sort(change.begin(), change.begin() + i);
            flag = 1;
            for (int j = 0; j < n; j++){
                if (change[j] != comp[j]){
                    flag = 0;
                    break;
                }
            }
            if (flag){
                sort(change.begin(), change.begin() + i + 1);
                break;
            }
        }
        if (flag){
            printf("Insertion Sort\n");
            printf("%d", change[0]);
            for (int i = 1; i < n; i++){
                printf(" %d", change[i]);
            }
            printf("\n");
        }
        else{
            //说明是堆排序的
            int index = 0;
            for (int i = 1; i < n; i++){
                if (comp[i] > comp[0]){
                    index = i;
                    break;
                }
            }
            //将堆顶的数据与最后一个数据进行交换
            comp[0] = comp[0] ^ comp[index - 1];
            comp[index - 1] = comp[0] ^ comp[index - 1];
            comp[0] = comp[0] ^ comp[index - 1];
            //调整堆
            for (int i = 0; i < index - 1; i++){
                if (i * 2 + 1 < index - 1 && i * 2 + 2 == index - 1){
                    if (comp[i] < comp[i * 2 + 1]){
                        comp[i] = comp[i] ^ comp[i * 2 + 1];
                        comp[i * 2 + 1] = comp[i] ^ comp[i * 2 + 1];
                        comp[i] = comp[i] ^ comp[i * 2 + 1];
                    }
                    break;
                }
                if (i * 2 + 1 == index - 1){
                    break;
                }
                int left = comp[i * 2 + 1];
                int right = comp[i * 2 + 2];
                if (comp[i] < left || comp[i] < right){
                    if (right > left){
                        comp[i * 2 + 2] = comp[i];
                        comp[i] = right;
                    }
                    else{
                        comp[i * 2 + 1] = comp[i];
                        comp[i] = left;
                    }
                }
            }
            printf("Heap Sort\n");
            printf("%d", comp[0]);
            for (int i = 1; i < n; i++){
                printf(" %d", comp[i]);
            }
            printf("\n");
        }
    }
}
```
[Quick Sort](https://www.patest.cn/contests/pat-a-practise/1101)


**解题思路：**
正向遍历数组一遍，用一个数来记录左边最大值，对于所有满足大于左边最大的数进行标记，再反向遍历数组一边。用一个数来记录右边最小值，对于所有满足小于右边最小的数进行二次标记，两次都有标记的数据就是满足要求的数据。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<string.h>
#include<string>
#include<vector>
#include<algorithm>
using namespace std;
int main(){
    for(int n;scanf("%d",&n)!=EOF;){
        vector<int>vec;
        vector<int>res;
        int *arr = new int[n];
        memset(arr,0,n*sizeof(int));
        int left_max = 0;
        int right_min = 1000000000;
        for(int i = 0;i < n;i++){
            int temp;
            scanf("%d",&temp);
            vec.push_back(temp);
        }
        for(int i = 0;i < vec.size();i++){
            if(vec[i] > left_max){
                arr[i]++;
                left_max = vec[i];
            }
        }
        for(int i = vec.size()-1;i >= 0;i--){
            if(vec[i] <= right_min){
                arr[i]++;
                right_min = vec[i];
            }
        }
        for(int i = 0; i < n;i++){
            if(arr[i] == 2){
                res.push_back(vec[i]);
            }
        }
        sort(res.begin(),res.end());
        printf("%d\n",res.size());
        if(res.size()){
            printf("%d",res[0]);
            for(int i = 1;i < res.size();i++){
                printf(" %d",res[i]);
            }
            printf("\n");
        }else{
            printf("\n");
        }
        delete[] arr;
    }
    return 0;
}
```
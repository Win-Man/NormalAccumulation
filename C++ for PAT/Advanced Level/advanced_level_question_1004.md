[Counting Leaves](https://www.patest.cn/contests/pat-a-practise/1004)

**解题思路：**
在输入的时候就标记哪些节点是非叶子节点，然后层次遍历就好

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
using namespace std;
struct node{
    int size;
    vector<int>children;
};
int main(){
    for (int n, k; scanf("%d %d", &n, &k) != EOF;){
        node arr[105];
        //记录节点是否是叶子节点
        int flag[105] = {0};
        for (int i = 0; i < k; i++){
            int temp,size;
            scanf("%d %d", &temp,&size);
            arr[temp].size = size;
            flag[temp] = 1;
            while (size--){
                int child;
                scanf("%d", &child);
                arr[temp].children.push_back(child);
            }
        }
        int count = 1;
        int every_count = 0;
        if (flag[1]){
            printf("0");
        }
        else{
            printf("1");
        }
        vector<node>vec;
        for (int i = 0; i < arr[1].size; i++){
            if (!flag[arr[1].children[i]]){
                every_count++;
            }
            vec.push_back(arr[arr[1].children[i]]);
            count++;
        }
        if (vec.size()){
            printf(" %d", every_count);
        }
        while (count != n){
            vector<node>haha;
            every_count = 0;
            for (int i = 0; i < vec.size(); i++){
                for (int j = 0; j < vec[i].size; j++){
                    if (!flag[vec[i].children[j]]){
                        every_count++;
                    }
                    haha.push_back(arr[vec[i].children[j]]);
                    count++;
                }
            }
            vec = vector<node>(haha);
            printf(" %d", every_count);
        }
        printf("\n");
    }
    return 0;
}
```
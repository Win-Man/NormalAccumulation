[The Largest Generation](https://www.patest.cn/contests/pat-a-practise/1094)


**解题思路：**
创建一个节点结构，每个节点中存放该节点的孩子数以及孩子的编号，然后从根节点开始，一层一层的统计每层的孩子的总数，当当前层的孩子总数比之前记录的总数大的时候，更新最大值。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<queue>
using namespace std;
struct node{
    int size;
    vector<int>children;
};
int main(){
    for (int n, m; scanf("%d%d", &n,&m) != EOF;){
        node *arr = new node[n+1];
        for (int i = 0; i < m; i++){
            int num, size, temp;
            scanf("%d", &num);
            scanf("%d", &size);
            arr[num].size = size;
            while (size--){
                scanf("%d", &temp);
                arr[num].children.push_back(temp);
            }
        }
        int count = 1;
        int level = 1;
        int ger_sum = 1;
        int ger_level = 1;
        vector<node>vec;
        vec.push_back(arr[1]);
        while (1){
            level++;
            int sum = 0;
            vector<node>temp;
            for (int i = 0; i < vec.size(); i++){
                sum += vec[i].size;
                for (int j = 0; j < vec[i].children.size(); j++){
                    temp.push_back(arr[vec[i].children[j]]);
                    count++;
                }
            }
            if (sum > ger_sum){
                ger_sum = sum;
                ger_level = level;
            }
            vec = vector<node>(temp);
            if (count == n){
                break;
            }
        }
        printf("%d %d\n", ger_sum, ger_level);
    }
    return 0;
}
```
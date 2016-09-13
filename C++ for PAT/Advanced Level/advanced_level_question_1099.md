[Build A Binary Search Tree](https://www.patest.cn/contests/pat-a-practise/1099)


**解题思路：**
首先根据输入构件一棵二叉树，对输入的数据进行从小到大排序，然后对二叉树进行先序遍历，将节点的编号与排序之后的数据的排序号进行一一对应，最后层次遍历输出。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<algorithm>
#include<queue>
using namespace std;
struct node{
    int num;
    int left;
    int right;
};
node arr[105];
int ans[105];
int coun = 1;
void sortTree(int root){
    if (root == -1){
        return;
    }
    int left = arr[root].left;
    int right = arr[root].right;
    sortTree(left);
    ans[root] = coun;
    coun++;
    sortTree(right);
}
int main(){
    for (int n; scanf("%d", &n) != EOF;){
        vector<int> nums(n);
        coun = 1;
        for (int i = 0; i < n; i++){
            scanf("%d%d", &arr[i].left, &arr[i].right);
            arr[i].num = i;
        }
        for (int i = 0; i < n; i++){
            scanf("%d", &nums[i]);
        }
        sort(nums.begin(), nums.end());
        sortTree(0);
        queue<int>que;
        printf("%d", nums[ans[0] - 1]);
        if (arr[0].left != -1){
            que.push(arr[0].left);
        }
        if (arr[0].right != -1){
            que.push(arr[0].right);
        }
        while (!que.empty()){
            int temp = que.front();
            que.pop();
            int left = arr[temp].left;
            int right = arr[temp].right;
            if (left != -1){
                que.push(left);
            }
            if (right != -1){
                que.push(right);
            }
            printf(" %d", nums[ans[temp]-1]);
        }
        printf("\n");
    }
    return 0;
}
```
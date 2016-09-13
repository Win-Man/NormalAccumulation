[Complete Binary Search Tree](https://www.patest.cn/contests/pat-a-practise/1064)


**解题思路：**
题意是需要将一串数字按照完全二叉搜索树的样子排列，然后再输出树的层次遍历。因为是完全二叉树，所以用一个一维数组就能表示一棵树了，因为二叉搜索树的中序遍历是将数字从小到大输出，所以先通过中序遍历，将完全二叉搜索树构造出来，然后将数组输出就行。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<string.h>
#include<algorithm>
using namespace std;
int tree[3000];
vector<int> nums;
int pos = 0;
void buildTree(int root){
    if (tree[root] == -1){
        return;
    }
    buildTree(root * 2 + 1);
    tree[root] = nums[pos];
    pos++;
    buildTree(root * 2 + 2);
}
int main(){
    for (int n; scanf("%d", &n) != EOF;){
        memset(tree, -1, sizeof(tree));
        nums = vector<int>(n);
        for (int i = 0; i < n; i++){
            scanf("%d", &nums[i]);
            tree[i] = 1;
        }
        sort(nums.begin(), nums.end());
        pos = 0;
        buildTree(0);
        printf("%d",tree[0]);
        for (int i = 1; i < n; i++){
            printf(" %d", tree[i]);
        }
        printf("\n");
    }
    return 0;
}
```
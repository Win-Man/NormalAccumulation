[Invert a Binary Tree](https://www.patest.cn/contests/pat-a-practise/1102)


**解题思路：**
就是简单的二叉树遍历

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<queue>
#include<vector>
using namespace std;
struct node{
    char left;
    char right;
};
node tree[15];
int cou = 0,n;
void levelShow(int root){
    queue<int>temp;
    temp.push(root);
    printf("%d", root);
    temp.pop();
    if (tree[root].right != '-'){
        temp.push(tree[root].right - '0');
    }
    if (tree[root].left != '-'){
        temp.push(tree[root].left - '0');
    }
    while (!temp.empty()){
        int show = temp.front();
        temp.pop();
        if (tree[show].right != '-'){
            temp.push(tree[show].right - '0');
        }
        if (tree[show].left != '-'){
            temp.push(tree[show].left - '0');
        }
        printf(" %d", show);
    }
}
void invertedShow(int root){
    if (tree[root].right != '-'){
        int right = tree[root].right - '0';
        invertedShow(right);
    }
    printf("%d", root);
    cou++;
    if (cou != n){
        printf(" ");
    }
    if (cou == n){
        printf("\n");
    }
    if (tree[root].left != '-'){
        int left = tree[root].left - '0';
        invertedShow(left);
    }
}
int main(){
    for (; scanf("%d", &n) != EOF;){
        int flag[15] = { 0 };
        for (int i = 0; i < n; i++){
            char left, right;
            scanf("\n%c %c", &left, &right);
            tree[i].left = left;
            tree[i].right = right;
            if (left != '-'){
                flag[left - '0']++;
            }
            if (right != '-'){
                flag[right - '0']++;
            }
        }
        //寻找根节点
        int root = 0;
        for (int i = 0; i < n; i++){
            if (flag[i]){
                continue;
            }
            root = i;
            break;
        }
        levelShow(root);
        printf("\n");
        cou = 0;
        invertedShow(root);
    }
    return 0;
}
```
[Tree Traversals Again](https://www.patest.cn/contests/pat-a-practise/1086)


**解题思路：**
与上一题一样，也是先根据输入构造二叉树然后进行后序遍历进行输出。只是这道题没有明确的给出前序遍历和中序遍历，中序遍历根据题目的输出模仿栈的压入弹出就可以得到，前序遍历的话需要自己想一下，其实就是压栈的顺序就是前序遍历。 
所以样例的前序遍历是：1 2 3 4 5 6，中序遍历是：3 2 4 1 6 5

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<stack>
using namespace std;
struct node{
    int left;
    int right;
};
node tree[35];
int index = 0;
int pre[35];
int in[35];
void post(int *pre, int * in, int pBegin, int pEnd, int iBegin, int iEnd){
    if (pBegin == pEnd || iBegin == iEnd){
        return;
    }
    int iMid = 0;
    for (int i = iBegin; i < iEnd; i++){
        if (in[i] == pre[pBegin]){
            iMid = i;
            break;
        }
    }
    int pMid = iMid - iBegin + pBegin + 1;
    if (pMid > pBegin + 1){
        tree[pBegin].left = pBegin + 1;
    }
    else{
        tree[pBegin].left = -1;
    }
    if (pEnd > pMid){
        tree[pBegin].right = pMid;
    }
    else{
        tree[pBegin].right = -1;
    }
    post(pre, in, pBegin + 1, pMid, iBegin, iMid);
    post(pre, in, pMid, pEnd, iMid + 1, iEnd);
}
void postShow(int root){
    if (root == -1){
        return;
    }
    postShow(tree[root].left);

    postShow(tree[root].right);
    if (index == 0){
        printf("%d", pre[root]);
        index++;
    }
    else{
        printf(" %d", pre[root]);
    }

}
int main(){
    for (int n; scanf("%d", &n) != EOF;){
        stack<int>temp;
        int pIndex = 0;
        int iIndex = 0;
        for (int i = 0; i < n * 2; i++){
            char op[6];
            scanf("%s", op);
            if (op[1] == 'u'){
                scanf("%d", &pre[pIndex]);
                temp.push(pre[pIndex]);
                pIndex++;
            }
            else{
                in[iIndex] = temp.top();
                temp.pop();
                iIndex++;
            }
        }
        index = 0;
        post(pre, in, 0, n, 0, n);
        postShow(0);
        printf("\n");
    }
    return 0;
}
```
[Tree Traversals Again](https://www.patest.cn/contests/pat-a-practise/1020)


**解题思路：**
根据题目给的后序遍历和中序遍历构建树，然后对数进行层次遍历输出结果。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<queue>
using namespace std;
struct node{
    int left;
    int right;
};
node que[35];
int index = 0;
void level(int *post, int *in, int pBegin, int pEnd, int iBegin, int iEnd){
    if (pEnd == pBegin || iBegin == iEnd){
        return;
    }
    //if (pEnd - pBegin == 1){
    //  //printf("%d ", post[pBegin]);
    //  return;
    //}

    int iMid;
    for (int i = iBegin; i < iEnd; i++){
        if (in[i] == post[pEnd - 1]){
            iMid = i;
            break;
        }
    }
    int pMid = pBegin + iMid - iBegin;
    //printf("%d ", post[pEnd - 1]);
    //que[index] = post[pEnd - 1];
    //index++;
    if (pMid > pBegin){
        que[pEnd-1].left = pMid-1;
    }
    else{
        que[pEnd - 1].left = -1;
    }
    if (pEnd - 1 > pMid){
        que[pEnd-1].right = pEnd-2;
    }
    else{
        que[pEnd - 1].right = -1;
    }
    level(post, in, pBegin, pMid, iBegin, iMid);

    level(post, in, pMid, pEnd - 1, iMid + 1, iEnd);

}
int main(){
    for (int n; scanf("%d", &n) != EOF;){
        int *post = new int[n];
        int *in = new int[n];
        for (int i = 0; i < n; i++){
            scanf("%d", &post[i]);
        }
        for (int i = 0; i < n; i++){
            scanf("%d", &in[i]);
        }
        level(post, in, 0, n, 0, n);
        queue<int>out;
        printf("%d", post[n - 1]);
        if (que[n-1].left != -1)
            out.push(que[n-1].left);
        if (que[n-1].right != -1)
            out.push(que[n - 1].right);
        while (!out.empty()){
            int temp = out.front();
            out.pop();
            printf(" %d", post[temp]);
            if (que[temp].left != -1){
                out.push(que[temp].left);
            }
            if (que[temp].right != -1){
                out.push(que[temp].right);
            }
        }
        printf("\n");
    }
    return 0;
}
```
[Insert or Merge](https://www.patest.cn/contests/pat-a-practise/1089)


**解题思路：**
对输入的数组进行插入排序或者是归并排序，中间过程中对数据进行比较，如果与评判数组一致则记录下，下一次排序的结果，若到最终都不一致，则进行另一种排序方式。我是先进行归并排序，在中间过程中寻找是否一致，若没有找到，则进行插入排序。 
做这道题的时候傻逼了，一看到归并排序直接上手就写，写了个二分的归并，看结果一直不对，才发现题目中的归并并不是二分的。 
编号为2的测试点是个坑，是经过别人提醒才知道的，这道题的插入排序默认一个数字算是属于有序的序列，所以要从两个数字楷开始进行插入排序。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<algorithm>
#include<vector>
using namespace std;
vector<int> inVec(105);
vector<int> judgeVec(105);
vector<int> outArr(105);
int n= 0;
//type = 1 merger type = 2 insert
int type = 0;
bool isMerger(vector<int>tempVec){
    int step = 2;
    while (step <= n){
        int flag = 0;
        int begin = 0;
        int end = step;
        while (end < n){
            sort(tempVec.begin() + begin, tempVec.begin() + end);
            begin = end;
            end += step;
        }
        sort(tempVec.begin() + begin, tempVec.begin() + n );

        step *= 2;

        if (type == 1){
            for (int i = 0; i < n; i++){
                outArr[i] = tempVec[i];
            }
            return true;
        }

        for (int i = 0; i < n; i++){
            if (tempVec[i] != judgeVec[i]){
                flag = 1;
                break;
            }
        }
        if (!flag){
            type = 1;
        }
    }
    return false;
}
void insert(vector<int>tempVec){
    //两个数字开始
    for (int i = 2; i <= n; i++){
        int flag = 0;
        sort(tempVec.begin(), tempVec.begin() + i);
        if (type == 2){
            for (int i = 0; i < n; i++){
                outArr[i] = tempVec[i];
            }
            return;
        }
        for (int i = 0; i < n; i++){
            if (tempVec[i] != judgeVec[i]){
                flag = 1;
            }
        }
        if (!flag){
            type = 2;
        }
    }
}
int main(){
    for (; scanf("%d", &n) != EOF;){
        type = 0;
        int temp[100] = { 0 };
        for (int i = 0; i < n; i++){
            scanf("%d", &inVec[i]);
        }
        for (int i = 0; i < n; i++){
            scanf("%d", &judgeVec[i]);
        }
        vector<int>tempVec(inVec);
        if (isMerger(tempVec)){
            cout << "Merge Sort" << endl;
        }
        else{
            cout << "Insertion Sort" << endl;
            insert(tempVec);
        }
        cout << outArr[0];
        for (int i = 1; i < n; i++){
            cout << " " << outArr[i];
        }
        cout << endl;
    }
    return 0;
}
```
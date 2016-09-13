[Lowest Price in Supply Chain](https://www.patest.cn/contests/pat-a-practise/1079)


**解题思路：**
简单的广搜，当搜索到某一层有零售商的时候就停止搜索，但是要记录那一层所有零售商的数量。 
要测试root直接为零售商的结果，一开始没有想到，导致测试点1一直运行超时，因为死循环了 

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<math.h>
using namespace std;
int main(){
    int n;
    double price;
    double rate;
    for (; scanf("%d%lf%lf", &n, &price, &rate) != EOF;){
        vector<int>* arr = new vector<int>[n];
        int *arr_check = new int[n];
        for (int i = 0; i < n; i++){
            int size;
            scanf("%d", &size);
            arr_check[i] = size;
            arr[i] = vector<int>(size);
            for (int j = 0; j < size;j++){
                //int temp;
                //scanf("%d", &temp);
                //arr[i].push_back(temp);
                scanf("%d", &arr[i][j]);
            }
        }
        int level = 0;
        int flag = 0;
        int count = 0;
        vector<int>vec;
        vec.push_back(0);
        //考虑root直接为零售商的情况
        if (arr[0].size() == 0){
            flag = 1;
            count++;
        }
        while (!flag){
            vector<int>temp;
            for (int i = 0; i < vec.size(); i++){
                for (int j = 0; j < arr[vec[i]].size(); j++){
                    if (arr_check[arr[vec[i]][j]] == 0){
                        flag = 1;
                        count++;
                    }
                    if (arr_check[arr[vec[i]][j]] != 0 && !flag)
                        temp.push_back(arr[vec[i]][j]);
                }
            }
            level++;
            vec = vector<int>(temp);
        }
        printf("%.4lf %d\n",price*pow(rate/100.0+1,level*1.0),count);
    }
    return 0;
}
```
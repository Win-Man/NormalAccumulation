[Higest Price in Supply Chain](https://www.patest.cn/contests/pat-a-practise/1090)


**解题思路：**
根据输入构建一棵树，然后找一下树的最大深度

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<stdio.h>
#include<math.h>
using namespace std;
int main(){
    int n;
    double price;
    double rate;
    for (; scanf("%d%lf%lf", &n, &price, &rate) != EOF;){
        vector<int> *vec = new vector<int>[n];
        int root;
        for (int i = 0; i < n; i++){
            int temp;
            scanf("%d", &temp);
            if (temp != -1){
                vec[temp].push_back(i);
            }
            else{
                root = i;
            }
        }
        int level = 0;
        int count = 1;
        vector<int>haha;
        haha.push_back(root);
        while (count != n){
            vector<int>temp;
            for (int i = 0; i < haha.size(); i++){
                for (int j = 0; j < vec[haha[i]].size(); j++){
                    temp.push_back(vec[haha[i]][j]);
                    count++;
                }
            }
            haha = vector<int>(temp);
            level++;
        }
        printf("%.2lf %d\n", price*pow(rate / 100.0 + 1.0, level*1.0), haha.size());
    }
    return 0;
}
```
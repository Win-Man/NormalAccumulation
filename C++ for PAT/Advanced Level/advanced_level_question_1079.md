[Total Sales of Supply Chain](https://www.patest.cn/contests/pat-a-practise/1079)


**解题思路：**
因为同类型的题目写过了，所以写起来很快，一遍过。

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
        int *pro_amount = new int[n];
        vector<int> *nodes = new vector<int>[n];
        for (int i = 0; i < n; i++){
            int size;
            scanf("%d", &size);
            if (size){
                nodes[i] = vector<int>(size);
                for (int j = 0; j < size;j++){
                    scanf("%d", &nodes[i][j]);
                }
                pro_amount[i] = 0;
            }
            else{
                nodes[i] = vector<int>(size);
                scanf("%d", &pro_amount[i]);
            }
        }
        int level = 0;
        double sum = 0.0;
        vector<int>vec;
        vec.push_back(0);
        while (vec.size()){
            vector<int>temp;
            for (int i = 0; i < vec.size(); i++){
                //是零售商
                if (pro_amount[vec[i]]){
                    sum += price*pow(rate / 100.0 + 1, level)*pro_amount[vec[i]];
                }
                else{
                    for (int j = 0; j < nodes[vec[i]].size(); j++){
                        temp.push_back(nodes[vec[i]][j]);
                    }
                }
            }
            level++;
            vec = vector<int>(temp);
        }
        printf("%.1lf\n", sum);
    }
    return 0;
}
```
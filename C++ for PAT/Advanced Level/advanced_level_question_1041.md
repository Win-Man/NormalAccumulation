[Be Unique](https://www.patest.cn/contests/pat-a-practise/1041)


**解题思路：**

**AC code:**

``` c++
#include<iostream>  
#include<stdio.h>  
#include<vector>  
#include<string.h>  
using namespace std;  
int main(){  
    int count[10005];  
    int n = 0;  
    for (; scanf("%d", &n) != EOF;){  
        memset(count, 0, sizeof(int)* 10005);  
        vector<int>numbers;  
        for (int i = 0; i < n; i++){  
            int in;  
            scanf("%d", &in);  
            numbers.push_back(in);  
            count[in]++;  
        }  
        bool flag = 0;  
        int index = 0;  
        for (index = 0; index < numbers.size(); index++){  
            if (count[numbers[index]] == 1){  
                flag = 1;  
                break;  
            }  
        }  
        if (flag){  
            cout << numbers[index] << endl;  
        }  
        else{  
            cout << "None" << endl;  
        }  
    }  
    return 0;  
}  
```
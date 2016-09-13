[General Palindromic Number](https://www.patest.cn/contests/pat-a-practise/1019)


**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
using namespace std;
vector<int> getNum(int n, int d){
    vector<int>vec;
    while (n){
        vec.push_back(n%d);
        n /= d;
    }
    return vec;
}
bool isPalin(vector<int>vec){
    bool flag = true;
    for (int i = 0; i < vec.size() / 2; i++){
        if (vec[i] != vec[vec.size() - 1 - i]){
            flag = false;
            break;
        }
    }
    return flag;
}
int main(){
    for (int n; scanf("%d", &n) != EOF;){
        int d;
        scanf("%d", &d);
        if (n == 0){
            printf("Yes\n0\n");
            continue;
        }
        vector<int>vec = getNum(n, d);
        if (isPalin(vec)){
            printf("Yes\n");
        }
        else{
            printf("No\n");
        }
        printf("%d", vec[vec.size() - 1]);
        for (int i = 1; i < vec.size(); i++){
            printf(" %d", vec[vec.size() - 1 - i]);
        }
        printf("\n");
    }
    return 0;
}
```
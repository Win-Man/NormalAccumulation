[Sun of Number Segments](https://www.patest.cn/contests/pat-a-practise/1104)


**解题思路：**


**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<string.h>
#include<iomanip>
using namespace std;
double arr[100005];
int main(){
  int n;
  for (; scanf("%d", &n) != EOF;){
    memset(arr, 0, sizeof(double)* 100005);
    for (int i = 0; i < n; i++){
      scanf("%lf", &arr[i]);
    }
    double sum = 0.0;
    for (int i = 0; i < n; i++){
      sum += (arr[i])*(n - i)*(i + 1);
    }
    cout << fixed << setprecision(2) << sum << endl;
  }
  return 0;
}   
```
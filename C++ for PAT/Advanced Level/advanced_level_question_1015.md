[Reversible Primes](https://www.patest.cn/contests/pat-a-practise/1015)

**解题思路：**
看题太快，题目要求原先的数和反转之后的数都是素数才输出Yes，一直以为是反转后的数是素数就好了，还有判断素数居然写错了，坑爹。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<string>
#include<math.h>
using namespace std;
int reverseNum(int num, int d){
    string reNum = "";
    while (num){
        reNum += (num%d) + '0';
        num /= d;
    }
    int ret = 0;
    for (int i = 0; i < reNum.length(); i++){
        ret += (reNum[reNum.length() - i - 1] - '0')*(int)pow(d*1.0, i*1.0);
    }
    return ret;
}
bool isPrime(int n){
    if (n < 2){
        return false;
    }
    for (int i = 2; i*i <= n; i++){
        if (n%i == 0){
            return false;
        }
    }
    return true;
}
int priFlag[100005];
int main(){
    for (int i = 0; i < 100005; i++){
        if (isPrime(i)){
            priFlag[i] = 1;
        }
        else{
            priFlag[i] = 0;
        }
    }
    for (int n; scanf("%d", &n) != EOF && n >= 0;){
        int d;
        scanf("%d", &d);
        int renum = reverseNum(n, d);
        if (priFlag[renum] && priFlag[n]){
            printf("Yes\n");
        }
        else{
            printf("No\n");
        }
    }
    return 0;
}
```
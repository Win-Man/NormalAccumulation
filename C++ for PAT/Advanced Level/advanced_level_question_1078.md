[Hashing](https://www.patest.cn/contests/pat-a-practise/1052)


**解题思路：**
坑爹的二次探测，一开始以为是一道很简单的题，结果还是被坑了，完全没有注意到冲撞处理策略，所以最后一个点一直过不了，后来看了别人的博客说要经过二次探测检测之后还是不能插入才输出“-”，不是一开始不能插入就直接输出“-”。too young 
二次冲撞检测就是，在i位置无法插入的情况下，进行i+1^2,i+2^2,i+3^2….依次递增的位置检测，如果检测到可以插入那就插入，一直找不到才输出“-”，网上别人定的范围是到30000^2，自己发现到10000^2也可以过，关于这个范围的指定不是很清楚为什么。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<string.h>
using namespace std;
bool isPrime(int n){
    if (n < 2){
        return false;
    }
    for (int i = 2; i*i <= n+1; i++){
        if (n%i == 0){
            return false;
        }
    }
    return true;
}
int main(){
    //10007
    int primes[10050];
    int per_prime = 0;
    for (int i = 10049; i >= 0; i--){
        if (isPrime(i)){
            per_prime = i;
            primes[i] = i;
        }
        else{
            primes[i] = per_prime;
        }
    }
    for (int m, n; scanf("%d%d", &m, &n) != EOF;){
        int size = primes[m];
        int *flag = new int[size];
        memset(flag, 0, sizeof(int)*size);
        int test = 0;
        while (n > 1 ){
            int temp; 
            test = 0;
            scanf("%d", &temp);
            for (int i = 0; i <= 10000; i++){
                if (!flag[(temp + i*i) % size]){
                    flag[(temp + i*i) % size] = 1;
                    test = 1;
                    printf("%d ", (temp + i*i) % size);
                    break;
                }
            }
            if (!test){
                printf("- ");
            }
            n--;
        }
        if (n){
            int temp;
            test = 0;
            scanf("%d", &temp);
            for (int i = 0; i <= 10000; i++){
                if (!flag[(temp + i*i) % size]){
                    flag[(temp + i*i) % size] = 1;
                    test = 1;
                    printf("%d\n", (temp + i*i) % size);
                    break;
                }
            }
            if (!test){
                printf("-\n");
            }
        }
    }
    return 0;
}
```
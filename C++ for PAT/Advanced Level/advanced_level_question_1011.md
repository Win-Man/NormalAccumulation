[World Cup Betting](https://www.patest.cn/contests/pat-a-practise/1011)


**解题思路：**
简单的题目，但是题目中给的示例输出是错误的，输出应该是 
T T W 37.97

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
using namespace std;
int main(){
    double game[3] = { 0 };
    for (; scanf("%lf %lf %lf", &game[0], &game[1], &game[2]) != EOF;){
        double odds1 = 0, odds2 = 0, odds3 = 0;
        char game1, game2, game3;
        if (game[0] >= game[1] && game[0] >= game[2]){
            odds1 = game[0];
            game1 = 'W';
        }
        if (game[1] >= game[0] && game[1] >= game[2]){
            odds1 = game[1];
            game1 = 'T';
        }
        if (game[2] >= game[0] && game[2] >= game[1]){
            odds1 = game[2];
            game1 = 'L';
        }

        scanf("%lf %lf %lf", &game[0], &game[1], &game[2]);
        if (game[0] >= game[1] && game[0] >= game[2]){
            odds2 = game[0];
            game2 = 'W';
        }
        if (game[1] >= game[0] && game[1] >= game[2]){
            odds2 = game[1];
            game2 = 'T';
        }
        if (game[2] >= game[0] && game[2] >= game[1]){
            odds2 = game[2];
            game2 = 'L';
        }

        scanf("%lf %lf %lf", &game[0], &game[1], &game[2]);
        if (game[0] >= game[1] && game[0] >= game[2]){
            odds3 = game[0];
            game3 = 'W';
        }
        if (game[1] >= game[0] && game[1] >= game[2]){
            odds3 = game[1];
            game3 = 'T';
        }
        if (game[2] >= game[0] && game[2] >= game[1]){
            odds3 = game[2];
            game3 = 'L';
        }

        double ans = (odds1*odds2*odds3*0.65 - 1) * 2;
        printf("%c %c %c %.2lf\n", game1, game2, game3, ans);
    }
    return 0;
}
```
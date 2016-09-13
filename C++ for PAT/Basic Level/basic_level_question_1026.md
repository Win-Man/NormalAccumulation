[程序运行时间](https://www.patest.cn/contests/pat-b-practise/1026)

``` c++
#include<iostream>
#include<stdio.h>
using namespace std;
int main(){
	for(int c1,c2;scanf("%d%d",&c1,&c2) != EOF;){
		int second = (int)((c2 - c1)/100.0+0.5);
		int clock=0,minute=0;
		clock = second / 3600;
		second %= 3600;
		minute = second / 60;
		second %= 60;
		printf("%02d:%02d:%02d\n",clock,minute,second);
	}
	return 0;
}
```
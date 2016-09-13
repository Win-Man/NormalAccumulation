[A+B in Hogwarts ](https://www.patest.cn/contests/pat-a-practise/1058)

``` c++
#include<iostream>
#include<stdio.h>
using namespace std;
int main(){
	//g s k 17 29
	int g1 = 0,g2 = 0,s1 = 0,s2 = 0,k1 = 0,k2 = 0;
	for(;scanf("%d.%d.%d %d.%d.%d",&g1,&s1,&k1,&g2,&s2,&k2)!=EOF;){
		int flag = 0;
		int res_k = (k1+k2)%29;
		flag = (k1+k2)/29;
		int res_s = (s1+s2+flag)%17;
		flag = (s1+s2+flag)/17;
		int res_g = g1+g2+flag;
		printf("%d.%d.%d\n",res_g,res_s,res_k);
	}
	return 0;
}

```
[ÔÚ»ô¸ñÎÖ´ÄÕÒÁãÇ®](https://www.patest.cn/contests/pat-b-practise/1037)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<stdio.h>
using namespace std;
int main(){
	int g1=0,g2=0,s1=0,s2=0,k1=0,k2=0;
	//14.1.28 10.16.27
	for(;scanf("%d.%d.%d %d.%d.%d",&g1,&s1,&k1,&g2,&s2,&k2)!=EOF;){
		int res_g = 0;
		int res_s = 0;
		int res_k = 0;
		int res = 0;
		int p = g1*29*17+s1*29+k1;
		int a = g2*29*17+s2*29+k2;
		if(a >= p){
			res = a-p;
			res_g = (res/(29*17));
			res_s = (res%(29*17))/29;
			res_k = (res%(29*17))%29;
			printf("%d.%d.%d\n",res_g,res_s,res_k);
		}else{
			res = p-a;
			res_g = (res/(29*17));
			res_s = (res%(29*17))/29;
			res_k = (res%(29*17))%29;
			printf("-%d.%d.%d\n",res_g,res_s,res_k);
		}
	}
	return 0;
}
```
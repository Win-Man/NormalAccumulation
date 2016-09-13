[The Dominant Color ](https://www.patest.cn/contests/pat-a-practise/1054)

``` c++
#include<iostream>
#include<map>
#include<stdio.h>
using namespace std;
int main(){
	for(int n,m;scanf("%d%d",&n,&m)!=EOF;){
		map<int,int>sence;
		for(int i = 0;i < n; i++){
			for(int j = 0;j < m;j++){
				int temp;
				scanf("%d",&temp);
				sence[temp]++;
			}
		}
		map<int,int>::iterator iter ;
		int max = 0;
		int max_num = 0;
		for(iter = sence.begin();iter != sence.end();iter++){
			if(iter->second > max){
				max = iter->second;
				max_num = iter->first;
			}
		}
		printf("%d\n",max_num);
	}
	return 0;
}
```
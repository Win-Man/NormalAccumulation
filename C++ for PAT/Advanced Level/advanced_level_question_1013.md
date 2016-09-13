[Battle Over Cities ](https://www.patest.cn/contests/pat-a-practise/1013)

``` c++
#include<iostream>
#include<stdio.h>
#include<string.h>
#include<vector>
using namespace std;
int father[1005];
int flag[1005];
void init(){
	memset(flag,0,sizeof(flag));
	for(int i = 0;i < 1005;i++){
		father[i] = i;
	}
}
int getfather(int x){
	if(x != father[x]){
		father[x] = getfather(father[x]);
	}
	return father[x];
}
void join(int x,int y){
	int i = getfather(x);
	int j = getfather(y);
	if(i != j){
		father[i] = j;
	}
}
struct way{
	int x,y;
};
int main(){
	for(int n,m,k;scanf("%d%d%d",&n,&m,&k)!= EOF;){
		vector<way>ways;
		for(int i = 0;i < m;i++){
			way temp;
			scanf("%d%d",&temp.x,&temp.y);
			ways.push_back(temp);
		}
		for(int i = 0;i < k;i++){
			int q;
			scanf("%d",&q);
			init();
			for(int j = 0;j < ways.size();j++){
				if(ways[j].x != q && ways[j].y != q){
					join(ways[j].x,ways[j].y);
				}
			}
			for(int j = 1;j <= n;j++){
				flag[getfather(j)] = 1;
			}
			int count = 0;
			for(int j = 1;j <= n;j++){
				if(j == q){
					continue;
				}
				if(flag[j]){
					count++;
				}
			}
			printf("%d\n",count-1);
		}
	}
	return 0;
}
```
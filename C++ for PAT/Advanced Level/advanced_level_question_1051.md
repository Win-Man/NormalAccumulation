[Pop Sequence ](https://www.patest.cn/contests/pat-a-practise/1051)

**AC code:**

``` c++
#include<iostream>
#include<stack>
#include<stdio.h>
using namespace std;
int main(){
	for(int m,n,k;scanf("%d%d%d",&m,&n,&k)!= EOF;){
		while(k--){
			//定义一个空栈 add是模拟栈压入的时候的数 flag用来标记是否可行
			stack<int>sta;
			int flag = 1;
			int add = 1;
			for(int i = 0; i < n;i++){
				int num;
				scanf("%d",&num);
				if(flag){
					//当栈为空栈或者说是栈的元素还不等于读取的数字的时候，一直对栈进行push操作，中间判断栈容量是否超过
					while(sta.size() == 0 || sta.top() != num){
						sta.push(add);
						if(sta.size() > m){
							flag = 0;
							break;
						}
						add++;
					}
					if(flag && sta.size() >= 1 && sta.top() == num){
						sta.pop();
					}
				}
			}
			if(flag){
				printf("YES\n");
			}else{
				printf("NO\n");
			}
		}
	}
	return 0;
}
```
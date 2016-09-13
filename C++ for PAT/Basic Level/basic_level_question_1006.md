[换个格式输出整数 ](https://www.patest.cn/contests/pat-b-practise/1006)

**AC code:**

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	for(int n;cin>>n;){
		if(n<10){
			for(int i = 1;i <= n;i++){
				cout<<i;
			}
			cout<<endl;
		}else if(n<100){
			cout<<string(n/10,'S');
			for(int i = 1;i <= n%10;i++){
				cout<<i;
			}
			cout<<endl;
		}else
		{
			cout<<string(n/100,'B');
			cout<<string(n/10%10,'S');
			for(int i = 1;i <= n%10;i++){
				cout<<i;
			}
			cout<<endl;
		}
	}
	return 0;
}
```
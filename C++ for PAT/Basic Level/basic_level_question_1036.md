[跟奥巴马一起编程](https://www.patest.cn/contests/pat-b-practise/1036)

**AC code:**

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	int n;
	char ch;
	for(;cin>>n>>ch;){
		int m = (int)(n*1.0/2.0+0.5);
		cout<<string(n,ch)<<endl;
		for(int i = 0;i < m-2;i++){
			cout<<ch<<string(n-2,' ')<<ch<<endl;
		}
		cout<<string(n,ch)<<endl;
	}
	return 0;
}
```
[Hello World for U ](https://www.patest.cn/contests/pat-a-practise/1031)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<math.h>
using namespace std;
int main(){
	for(string str;cin>>str;){
		int n1,n2;
		n1 = (str.length()+2)/3;
		n2 = str.length()-n1*2 + 2;
		int j;
		for(j = 0;j < n1-1;j++){
			cout<<str[j]<<string(n2-2,' ')<<str[str.length()-1-j]<<endl;
		}
		for(;j < n1+n2-1;j++){
			cout<<str[j];
		}
		cout<<endl;
	}
	return 0;
}
```
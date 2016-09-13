[Colors in Mars ](https://www.patest.cn/contests/pat-a-practise/1027)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<map>
using namespace std;
int main(){
	string str_13 = "0123456789ABC";
	for(int num1,num2,num3;cin>>num1>>num2>>num3;){
		int high = num1 / 13;
		int low = num1 % 13;
		cout<<"#"<<str_13[high]<<str_13[low];
		high = num2 / 13;
		low = num2 % 13;
		cout<<str_13[high]<<str_13[low];
		high = num3 / 13;
		low = num3 % 13;
		cout<<str_13[high]<<str_13[low]<<endl;
	}
	return 0;
}
```
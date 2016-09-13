[D进制的A+B ](https://www.patest.cn/contests/pat-b-practise/1022)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
	for(int a,b,d;cin>>a>>b>>d;){
		int sum = a + b;
		vector<short>res;
		while(sum){
			res.push_back(sum%d);
			sum/=d;
		}
		if(res.size()){
			if(res[res.size()-1]){
				cout<<res[res.size()-1];
			}
			for(int i = 0;i<res.size()-1;i++){
				cout<<res[res.size()-2-i];
			}
			cout<<endl;
		}else{
			cout<<"0"<<endl;
		}
		
	}
	return 0;
}
/*考虑输入为0的情况*/
```
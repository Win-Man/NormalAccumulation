[写出这个数 ](https://www.patest.cn/contests/pat-b-practise/1002)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
	string pinyin[10]={"ling","yi","er","san","si","wu","liu","qi","ba","jiu"};
	string num;
	cin>>num;
	int sum = 0;
	vector<string>out;
	for(int i =0;i<num.length();i++){
		sum+= num[i]-'0';
	}
	while(sum){
		out.push_back(pinyin[sum%10]);
		sum/=10;
	}
	if(out.size() == 1){
		cout<<out[0]<<endl;
	}else
	{
		cout<<out[out.size()-1];
		for(int i =1 ;i<out.size();i++){
			cout<<" "<<out[out.size()-i-1];
		}
		cout<<endl;
	}
	return 0;
}
```
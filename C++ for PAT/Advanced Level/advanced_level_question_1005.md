[Spell It Right ](https://www.patest.cn/contests/pat-a-practise/1005)

``` c++
#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
	for(string str;cin>>str;){
		int sum = 0;
		string list[10]={"zero","one","two","three","four","five","six","seven","eight","nine"};
		vector<int>res;
		for(int i = 0;i < str.length();i++){
			sum += (str[i]-'0');
		}
		while(sum){
			res.push_back(sum%10);
			sum /= 10;
		}
		if(res.size()){
			cout<<list[res[res.size()-1]];
			for(int i = res.size()-2;i > -1;i--){
				cout<<" "<<list[res[i]];
			}
			cout<<endl;
		}else{
			cout<<"zero"<<endl;
		}
	}
	return 0;
}
```
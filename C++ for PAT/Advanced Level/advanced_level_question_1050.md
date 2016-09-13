[String Subtraction ](https://www.patest.cn/contests/pat-a-practise/1050)

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	for(string s1;getline(cin,s1);){
		string s2;
		getline(cin,s2);
		int flag[257]={0};
		for(int i = 0;i < s2.length();i++){
			flag[s2[i]] = 1;
		}
		for(int i = 0; i < s1.length();i++){
			if(!flag[s1[i]]){
				cout<<s1[i];
			}
		}
		cout<<endl;
	}
	return 0;
}
```
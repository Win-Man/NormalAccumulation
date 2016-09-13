[ÓÐ¼¸¸öPAT](https://www.patest.cn/contests/pat-b-practise/1040)

``` c++
#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
	for(string str;cin>>str;){
		int all_t = 0;
		for(int i = 0;i < str.length();i++){
			if(str[i] == 'T'){
				all_t ++;
			}
		}
		int sum = 0;
		int count_P = 0;
		int count_T = 0;
		for(int i = 0;i < str.length();i++){
			if(str[i] == 'P'){
				count_P++;
			}else if(str[i] == 'T'){
				count_T++;
			}else{
				sum+=(count_P*(all_t-count_T))%1000000007;
				sum = sum % 1000000007;
			}
		}
		cout<<sum<<endl;
	}
	return 0;
}
```
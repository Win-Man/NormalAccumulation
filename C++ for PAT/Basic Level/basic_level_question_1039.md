[µ½µ×Âò²»Âò](https://www.patest.cn/contests/pat-b-practise/1039)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<stdio.h>
using namespace std;
int main(){
	for(string sale,want;cin>>sale>>want;){
		int sale_number[10]={0};
		int sale_word_lower[26]={0};
		int sale_word_upper[26]={0};
		int want_number[10]={0};
		int want_word_lower[26]={0};
		int want_word_upper[26]={0};
		for(int i = 0; i < sale.length();i++){
			if(sale[i] >= '0' && sale[i] <= '9'){
				sale_number[sale[i] - '0']++;
			}else if(sale[i] >= 'a' && sale[i] <= 'z'){
				sale_word_lower[sale[i] - 'a']++;
			}else{
				sale_word_upper[sale[i] - 'A']++;
			}
		}
		for(int i = 0; i < want.length();i++){
			if(want[i] >= '0' && want[i] <= '9'){
				want_number[want[i] - '0']++;
			}else if(want[i] >= 'a' && want[i] <= 'z'){
				want_word_lower[want[i] - 'a']++;
			}else{
				want_word_upper[want[i] - 'A']++;
			}
		}
		int flag = 0;
		for(int i = 0;i < 10;i++){
			if(want_number[i] > sale_number[i]){
				flag  = flag + sale_number[i] - want_number[i];
			}
		}
		for(int i =0;i < 26;i++){
			if(want_word_lower[i] > sale_word_lower[i]){
				flag = flag + sale_word_lower[i]- want_word_lower[i];
			}
		}
		for(int i =0;i < 26;i++){
			if(want_word_upper[i] > sale_word_upper[i]){
				flag = flag + sale_word_upper[i] - want_word_upper[i];
			}
		}
		if(flag < 0){
			flag = -flag;
			cout<<"No "<<flag<<endl;
		}else{
			cout<<"Yes "<<sale.length()-want.length()<<endl;
		}
	}
	return 0;
}
```
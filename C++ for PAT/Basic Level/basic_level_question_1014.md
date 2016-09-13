[福尔摩斯的约会 ](https://www.patest.cn/contests/pat-b-practise/1014)

**AC code:**

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	string week[7]={"MON","TUE","WED","THU","FRI","SAT","SUN"};
	for(string str1,str2,str3,str4;cin>>str1>>str2>>str3>>str4;){
		int index = 0;
		for(int i = 0;i<str1.length()&&i<str2.length();i++){
			if(str1[i] >= 'A' && str1[i] <= 'G'){
				if(str1[i] == str2[i]){
					index = i;
					cout<<week[str1[i] - 'A']<<" ";
					break;
				}
			}
		}
		for(int i = index+1;i<str1.length()&& i<str2.length();i++){
			if((str1[i] >= '0' && str1[i] <= '9')||(str1[i] >= 'A' && str1[i] <= 'N')){
				if(str1[i] == str2[i]){
					if(str1[i] >= '0' && str1[i] <= '9'){
						cout<<"0"<<str1[i]<<":";
						break;
					}else{
						cout<<str1[i] - 'A' + 10<<":";
						break;
					}
				}
			}
		}
		if(str3.length() < str4.length()){
			for(int i = 0;i<str3.length();i++){
				if((str3[i] >= 'A' && str3[i] <= 'Z')|| (str3[i] >= 'a' && str3[i] <= 'z')){
					if(str3[i] == str4[i]){
						cout<<i/10<<i%10<<endl;
						break;
					}
				}
			}
		}else{
			for(int i = 0;i < str4.length();i++){
				if((str3[i] >= 'A' && str3[i] <= 'Z')|| (str3[i] >= 'a' && str3[i] <= 'z')){
					if(str3[i] == str4[i]){
						cout<<i/10<<i%10<<endl;
						break;
					}
				}
			}
		}
	}
	return 0;
}
```
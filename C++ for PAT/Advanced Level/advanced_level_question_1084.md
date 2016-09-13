[Broken Keyboard ](https://www.patest.cn/contests/pat-a-practise/1084)

**AC code:**

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	for(string str1,str2;cin>>str1>>str2;){
		int word[26]={0};
		int number[10]={0};
		int flag1 = 0;
		for(int i = 0;i < str1.length();i++){
			int flag = 0;
			if(str1[i] >= 'a' && str1[i] <= 'z'){
				if(word[str1[i] - 'a']){
					continue;
				}else{
					for(int j = 0;j < str2.length();j++){
						if(str2[j] == str1[i]){
							flag = 1;
							break;
						}
					}
					if(!flag){
						word[str1[i] - 'a']++;
						cout<<(char)(str1[i]-'a'+'A');
					}
				}
			}else if(str1[i] >= 'A' && str1[i] <= 'Z'){
				if(word[str1[i]-'A']){
					continue;
				}else{
					for(int j = 0;j < str2.length();j++){
						if(str2[j] == str1[i]){
							flag = 1;
							break;
						}
					}
					if(!flag){
						word[str1[i] - 'A']++;
						cout<<str1[i];
					}
				}
			}else if(str1[i] >= '0' && str1[i] <= '9'){
				if(number[str1[i]-'0']){
					continue;
				}else{
					for(int j = 0;j< str2.length();j++){
						if(str2[j] == str1[i]){
							flag = 1;
							break;
						}
					}
					if(!flag){
						number[str1[i] - '0']++;
						cout<<str1[i];
					}
				}
			}else if(str1[i] == '_'){
				if(flag1){
					continue;
				}else{
					for(int j = 0;j< str2.length();j++){
						if(str2[j] == str1[i]){
							flag = 1;
							break;
						}
					}
					if(!flag){
						flag1++;
						cout<<str1[i];
					}
				}
			}
		}
		cout<<endl;
	}
	return 0;
}
```
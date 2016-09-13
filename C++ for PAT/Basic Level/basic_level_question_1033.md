[旧键盘打字](https://www.patest.cn/contests/pat-b-practise/1033)

``` c++
#include<iostream>
#include<string>
#include<stdio.h>
#include<vector>
using namespace std;
int main(){
	for(string broken,output;getline(cin,broken)&&getline(cin,output);){
		int word[26]={0};
		int number[10]={0};
		int sign[5]={0};//以及下划线“_”（代表空格）、“,”、“.”、“-”、“+”（代表上档键）
		for(int i =0;i<broken.length();i++){
			if(broken[i] >= '0'&& broken[i] <= '9'){
				number[broken[i] - '0']++;
			}else if(broken[i] >= 'A'&& broken[i] <= 'Z'){
				word[broken[i] - 'A']++;
			}else{
				switch (broken[i])
				{
				case '_':
					sign[0]++;
					break;
				case ',':
					sign[1]++;
					break;
				case '.':
					sign[2]++;
					break;
				case '-':
					sign[3]++;
					break;
				case '+':
					sign[4]++;
					break;
				default:
					break;
				}
			}
		}
		for(int i =0;i < output.length();i++){
			if(output[i] >= '0' && output[i] <= '9'){
				if(!number[output[i] - '0']){
					cout<<output[i];
				}
			}else if(output[i] >= 'a' && output[i] <= 'z'){
				if(!word[output[i] - 'a']){
					cout<<output[i];
				}
			}else if(output[i] >= 'A' && output[i] <= 'Z'){
				if(!word[output[i] - 'A'] && !sign[4]){
					cout<<output[i];
				}
			}else{
				switch (output[i])
				{
				case '_':
					if(!sign[0]){
						cout<<output[i];
					}
					break;
				case ',':
					if(!sign[1]){
						cout<<output[i];
					}
					break;
				case '.':
					if(!sign[2]){
						cout<<output[i];
					}
					break;
				case '-':
					if(!sign[3]){
						cout<<output[i];
					}
					break;
				case '+':
					if(!sign[4]){
						cout<<output[i];
					}
					break;
				default:
					break;
				}
			}
		}
		cout<<endl;
	}
	return 0;
}
/*考虑第一行输入为空行的情况*/
```
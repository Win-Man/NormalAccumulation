[Êä³öPATest](https://www.patest.cn/contests/pat-b-practise/1043)

``` c++
#include<iostream>
#include<string>
#include<stdio.h>
using namespace std;
int main(){
	for(string str;cin>>str;){
		//PATest
		int arr[6]={0};
		for(int i = 0;i < str.length();i++){
			switch (str[i])
			{
			case 'P':
				arr[0]++;
				break;
			case 'A':
				arr[1]++;
				break;
			case 'T':
				arr[2]++;
				break;
			case 'e':
				arr[3]++;
				break;
			case 's':
				arr[4]++;
				break;
			case 't':
				arr[5]++;
				break;
			default:
				break;
			}
		}
		while(1){
			int flag = 1;
			if(arr[0]){
				arr[0]--;
				flag = 0;
				cout<<"P";
			}
			if(arr[1]){
				arr[1]--;
				flag = 0;
				cout<<"A";
			}
			if(arr[2]){
				arr[2]--;
				flag = 0;
				cout<<"T";
			}
			if(arr[3]){
				arr[3]--;
				flag = 0;
				cout<<"e";
			}
			if(arr[4]){
				arr[4]--;
				flag = 0;
				cout<<"s";
			}
			if(arr[5]){
				arr[5]--;
				flag = 0;
				cout<<"t";
			}
			if(flag){
				break;
			}
		}
		cout<<endl;
	}
	return 0;
}
```
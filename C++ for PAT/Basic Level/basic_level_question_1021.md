[个位数统计 ](https://www.patest.cn/contests/pat-b-practise/1021)

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	for(string str;cin>>str;){
		int arr[10]={0};
		for(int i =0;i < str.length();i++){
			arr[str[i] - '0']++;
		}
		for(int i = 0;i < 10;i++){
			if(arr[i]){
				cout<<i<<":"<<arr[i]<<endl;
			}
		}
	}
	return 0;
}
```
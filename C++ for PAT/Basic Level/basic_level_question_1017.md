[A³ıÒÔB ](https://www.patest.cn/contests/pat-b-practise/1017)

``` c++
#include<iostream>
#include<string>
#include<math.h>
using namespace std;
int main(){
	string str;
	int num;
	for(;cin>>str>>num;){
		if(str.length() == 1){
			cout<<(str[0] - '0')/num<<" "<<(str[0] - '0')%num<<endl;
		}else if(str.length() == 2){
			cout<<((str[0] - '0')*10 + str[1] - '0')/num<<" "
				<<((str[0] - '0')*10 + str[1] - '0')%num<<endl;
		}else{
			int remainder = ((str[0] - '0')*10 + str[1] - '0')%num;
			cout<<((str[0] - '0')*10 + str[1] - '0')/num;
			for(int i = 2;i < str.length();i++){
				int temp = str[i] - '0' + remainder*10;
				cout<<temp/num;
				remainder = temp % num;
			}
			cout<<" "<<remainder<<endl;
		}
	}
	return 0 ;
}
```
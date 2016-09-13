[×Ö·ûÍ³¼Æ](https://www.patest.cn/contests/pat-b-practise/1042)

``` c++
#include<iostream>
#include<string>
#include<map>
using namespace std;
int main(){
	for(string str;getline(cin,str);){
		map<char,int>arr;
		for(int i = 0;i < str.length();i++){
			if(str[i] >= 'a' && str[i] <= 'z'){
				arr[str[i]]++;
			}else if(str[i] >= 'A' && str[i] <= 'Z'){
				arr[(char)(str[i]-'A'+'a')]++;
			}
		}
		map<char,int>::iterator iter = arr.begin();
		char ch = iter->first;
		int times = iter->second;
		for(;iter!= arr.end();iter++){
			if(iter->second > times){
				ch = iter->first;
				times = iter->second;
			}
		}
		cout<<ch<<" "<<times<<endl;
	}
	return 0;
}
```
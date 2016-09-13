[我要通过！](https://www.patest.cn/contests/pat-b-practise/1003)

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	int n;
	cin>>n;
	string str;
	for(int i =0;i<n;i++){
		cin>>str;
		int index_P=0,index_T=0,count_P=0,count_T=0,count_A=0;
		for(int i = 0;i<str.length();i++){
			if(str[i] == 'P'){
				index_P = i;
				count_P++;

			}else if(str[i] == 'T'){
				index_T = i;
				count_T++;

			}else if(str[i] == 'A'){
				count_A++;
			}
		}
		if(count_A + count_P + count_T != str.length() || count_P > 1
			|| count_T > 1 || index_T - index_P <=1 ||(index_T-index_P-1)*index_P != str.length()-index_T -1){
			cout<<"NO"<<endl;
		}else
		{
			cout<<"YES"<<endl;
		}
	}
	return 0;
}
```
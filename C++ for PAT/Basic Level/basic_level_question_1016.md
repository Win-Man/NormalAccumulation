[²¿·ÖA+B ](https://www.patest.cn/contests/pat-b-practise/1016)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<math.h>
using namespace std;
//string add(string str1,string str2){
//	reverse(str1.begin(),str1.end());
//	reverse(str2.begin(),str2.end());
//	if(str1.length() < str2.length()){
//		int carry = 0;
//		for(int i = 0;i < str1.length();i++){
//			if()
//		}
//	}
//}
int main(){
	string str_A,str_B;
	char a,b;
	for(;cin>>str_A>>a>>str_B>>b;){
		int count_A = 0,count_B = 0;
		for(int i =0;i < str_A.length();i++){
			if(str_A[i] == a){
				count_A++;
			}
		}
		for(int j =0;j < str_B.length();j++){
			if(str_B[j] == b){
				count_B++;
			}
		}
		long long int sum = 0;
		for(int i = 0;i<count_A;i++){
			sum+= pow(10,i)*(a - '0');
		}
		for(int i = 0;i < count_B;i++){
			sum+= pow(10,i)*(b - '0');
		}
		cout<<sum<<endl;
	}
	return 0 ;
}
```
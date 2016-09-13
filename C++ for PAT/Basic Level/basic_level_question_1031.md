[查验身份证](https://www.patest.cn/contests/pat-b-practise/1031)

``` c++
#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
	for(int n;cin>>n;){
		int arr[17]={7,9,10,5,8,4,2,1,6,3,7,9,10,5,8,4,2};
		char match[11]={'1','0','X','9' ,'8','7','6','5','4','3','2'};
		int all_pass = 1;
		for(int i =0;i < n;i++){
			int flag = 1;
			string str;
			int sum = 0;
			cin>>str;
			for(int j = 0;j < str.length()-1;j++){
				if(str[j] >= '0' && str[j] <= '9'){
					sum+= arr[j]*(str[j] - '0');
				}else{
					flag = 0;
					break;
				}
			}
			if(match[sum%11] != str[str.length()-1]){
				flag = 0;
			}
			if(!flag){
				all_pass = 0;
				cout<<str<<endl;
			}
		}
		if(all_pass){
			cout<<"All passed"<<endl;
		}
	}
	return 0;
}
```
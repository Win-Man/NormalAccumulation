[A+B and C (64bit) ](https://www.patest.cn/contests/pat-a-practise/1065)

``` c++
#include<iostream>
using namespace std;
int main(){
	for(int n;cin>>n;){
		for(int i = 1;i <= n;i++){
			long long int a,b,c;
			cin>>a>>b>>c;
			long long int sum = a+b;
			if(a > 0 && b > 0 && sum <= 0){
				cout<<"Case #"<<i<<": true"<<endl;;
			}else if(a < 0 && b < 0 && sum >= 0){
				cout<<"Case #"<<i<<": false"<<endl;
			}else if(sum > c){
				cout<<"Case #"<<i<<": true"<<endl;;
			}else
			{
				cout<<"Case #"<<i<<": false"<<endl;
			}
		}
	}
	return 0;
}
```
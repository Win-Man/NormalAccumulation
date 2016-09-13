[A+BºÍC ](https://www.patest.cn/contests/pat-b-practise/1011)

``` c++
#include<iostream>
using namespace std;
int main(){
	for(int n;cin>>n;){
		for(int i = 1;i <= n;i++){
			long long int a,b,c;
			cin>>a>>b>>c;
			if((a + b) > c){
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
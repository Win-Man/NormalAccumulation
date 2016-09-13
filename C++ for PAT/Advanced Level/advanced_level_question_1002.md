[A+B for Polynomials ](https://www.patest.cn/contests/pat-a-practise/1002)

**AC code:**

``` c++
#include<iostream>
#include<map>
#include<iomanip>
using namespace std;
int main(){
	for(int n1,n2;cin>>n1;){
		double ploy1[1005]={0.0};
		double ploy2[1005]={0.0};
		double res[1005]={0.0};
		for(int i = 0;i < n1;i++){
			int m;
			cin>>m;
			cin>>ploy1[m];
		}
		cin>>n2;
		for(int i = 0;i < n2;i++){
			int m;
			cin>>m;
			cin>>ploy2[m];
		}
		int count = 0;
		for(int i = 0;i < 1005;i++){
			if(ploy1[i] + ploy2[i] != 0.0){
				count++;
			}
			res[i] = ploy1[i] + ploy2[i];
		}
		cout<<count;
		for(int i = 1004;i >= 0;i--){
			if(res[i] != 0.0){
				cout<<" "<<i<<" "<<fixed<<setprecision(1)<<res[i];
			}
		}
		cout<<endl;
	}
	return 0;
}
```
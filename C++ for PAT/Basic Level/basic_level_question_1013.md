[ÊıËØÊı ](https://www.patest.cn/contests/pat-b-practise/1013)

``` c++
#include<iostream>
#include<math.h>
#include<vector>
using namespace std;
bool isPrimer(int n){
	for(int i=2;i<(int)sqrt(n)+1;i++){
		if(n%i == 0){
			return false;
		}
	}
	return true;
}
int main(){
	vector<int>vec;
	for(int i =2;i < 110000;i++){
		if(isPrimer(i)){
			vec.push_back(i);
		}
	}
	for(int n,m;cin>>n>>m;){
		int count = 0;
		for(int i = n-1;i<m;i++){
			count++;
			cout<<vec[i];
			if(i+1 == m){
				cout<<endl;
			}else{
				if(count == 10){
					cout<<endl;
					count = 0;
				}else{
					cout<<" ";
				}	
			}
			
		}
	}
	return 0;
}
```
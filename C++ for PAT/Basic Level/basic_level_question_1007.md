[ËØÊı¶Ô²ÂÏë ](https://www.patest.cn/contests/pat-b-practise/1007)

**AC code:**

``` c++
#include<iostream>
#include<math.h>
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
	int arr[100005]={0,0,0,0,0,1};
	int pre_primer = 5;
	for(int i = 6;i<100005;i++){
		if(isPrimer(i) ){
			if(i - pre_primer == 2){
				arr[i] = arr[i-1]+1;
			}else
			{
				arr[i] = arr[i-1];
			}
			pre_primer = i;
		}else 
		{
			arr[i] = arr[i-1];
		}
	}
	for(int n;cin>>n;){
		cout<<arr[n]<<endl;
	}
	return 0;
}
```
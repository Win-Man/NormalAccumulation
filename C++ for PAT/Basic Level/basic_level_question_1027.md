[´òÓ¡É³Â©](https://www.patest.cn/contests/pat-b-practise/1027)

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	int n;
	char ch;
	for(;cin>>n>>ch;){
		int sum = 1;
		int max = 0;
		for(int i = 3;;i =  i + 2){
			sum += i*2;
			if(sum > n){
				max = i - 2;
				sum -= i*2;
				break;
			}
		}
		for(int i=0;i<max;i=i+2){
			cout<<string(i/2,' ')<<string(max-i,ch)<<endl;
		}
		for(int i =3;i<=max;i=i+2){
			cout<<string((max-i)/2,' ')<<string(i,ch)<<endl;
		}
		cout<<n-sum<<endl;
	}
	return 0;
}
```
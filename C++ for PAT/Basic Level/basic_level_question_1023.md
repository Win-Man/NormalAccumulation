[组个最小数 ](https://www.patest.cn/contests/pat-b-practise/1023)

``` c++
#include<iostream>
#include<string>
#include<vector>
using namespace std;
int main(){
		int arr[10]={0};
		for(int i = 0;i < 10;i++){
			int n;
			cin>>n;
			arr[i]=n;
		}
		for(int i =1;i < 10;i++){
			if(arr[i]){
				cout<<i;
				arr[i]--;
				break;
			}
		}
		for(int i =0;i < 10;i++){
			cout<<string(arr[i],i + '0');
		}
		cout<<endl;
	return 0;
}
```
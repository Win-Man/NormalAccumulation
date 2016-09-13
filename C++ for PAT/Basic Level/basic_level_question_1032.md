[挖掘机技术哪家强](https://www.patest.cn/contests/pat-b-practise/1032)

``` c++
#include<iostream>
#include<string>
#include<vector>
using namespace std;

int main(){
	for(int n;cin>>n;){
		int max = 1;
		int school[100005]={0};
		for(int i = 0;i < n;i++){
			int id,score;
			cin>>id>>score;
			if(id > max){
				max = id;
			}
			school[id]+= score;
		}
		int res_index = 1;
		int sum = school[1];
		for(int i =1;i<=max;i++){
			if(school[i] > sum){
				sum = school[i];
				res_index = i;
			}
		}
		cout<<res_index<<" "<<school[res_index]<<endl;
	}
	return 0;
}
```
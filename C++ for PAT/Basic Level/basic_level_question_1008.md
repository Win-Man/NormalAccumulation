[数组元素循环右移问题 ](https://www.patest.cn/contests/pat-b-practise/1008)

**AC code:**

``` c++
#include<iostream>
#include<vector>
using namespace std;
int main(){
	for(int n,m;cin>>n>>m;){
		vector<int>vec;
		int index = 0;
		for(int i = 0;i < n ;i ++){
			int temp;
			cin>>temp;
			vec.push_back(temp);
			if((i+m)%n == 0){
				index = i;
			}
		}
		cout<<vec[index];
		for(int i = 1;i < vec.size();i++){
			cout<<" "<<vec[(index+i)%n];
		}
		cout<<endl;
	}
	return 0;
}
```
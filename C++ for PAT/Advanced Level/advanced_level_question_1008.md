[Elevator ](https://www.patest.cn/contests/pat-a-practise/1008)

**AC code:**

``` c++
#include<iostream>
#include<vector>
using namespace std;
int main(){
	for(int n;cin>>n;){
		vector<int>vec;
		for(int i = 0;i < n;i++){
			int m;
			cin>>m;
			vec.push_back(m);
		}
		int sum = vec[0]*6 + 5;
		if(vec.size() > 1){
			for(int i = 1;i < vec.size();i++){
				if(vec[i] > vec[i-1]){
					sum+= ((vec[i] - vec[i-1])*6 + 5);
				}else if( vec[i] < vec[i-1]){
					 sum+= ((vec[i-1] - vec[i])*4 + 5);
				}else{
					sum += 5;
				}
			}
		}
		cout<<sum<<endl;
	}
	return 0;
}
/*当前后两个数停留在同一层的时候还需要增加5秒的停留时间*/
```
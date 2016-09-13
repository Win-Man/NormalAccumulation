[¼ÌÐø(3n+1)²ÂÏë ](https://www.patest.cn/contests/pat-b-practise/1005)

``` c++
#include<iostream>
#include<string>
#include<sstream>
#include<stdio.h>
#include<algorithm>
#include<vector>
using namespace std;
int main(){
	for(int n;cin>>n;){
		vector<int>vec;
		vector<int>res;
		int arr[10000]={0};
		for(int i =0;i < n;i++){
			int temp;
			cin>>temp;
			vec.push_back(temp);
		}
		for(int i = 0;i < vec.size();i++){
			int m = vec[i];
			while(1){
				if(m & 1){
					m = (m*3+1)>>1;
					arr[m]++;
				}else{
					m = m>>1;
					arr[m]++;
				}
				if(m == 1){
					break;
				}
			}
		}
		for(int i =0;i < vec.size();i++){
			if(arr[vec[i]]){
				continue;
			}else{
				res.push_back(vec[i]);
			}
		}
		sort(res.begin(),res.end());
		if(res.size()){
			cout<<res[res.size()-1];
			for(int i =res.size()-2;i > -1;i--){
				cout<<" "<<res[i];
			}
			cout<<endl;
		}else{
			cout<<endl;
		}
	}
	return 0;
}
```
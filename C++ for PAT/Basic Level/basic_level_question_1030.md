[完美数列](https://www.patest.cn/contests/pat-b-practise/1030)

``` c++
#include<iostream>
#include<stdio.h>
#include<string.h>
#include<string>
#include<vector>
#include<algorithm>
using namespace std;
int main(){
	for(long long int n,p;cin>>n>>p;){
		vector<long long int>vec;
		for(int i = 0 ;i < n;i++){
			long long int temp;
			scanf("%d",&temp);
			vec.push_back(temp);
		}
		sort(vec.begin(),vec.end());
		long long int max = 0;//the number of list
		long long int max_index = 0;//the index of max number in the list
		for(long long int i = 0;i < vec.size();i++){
			for(int j = max_index;j < vec.size();j++){
				if(vec[j] <= vec[i]*p){
					if(j == vec.size()-1){
						max_index = j;
						if(max_index - i + 1 > max){
							max = max_index - i + 1;
						}
						break;
					}
					continue;
				}else{
					max_index = j-1;
					if(max_index - i + 1 > max){
						max = max_index - i + 1;
					}
					break;
				}
			}
			if(max_index == vec.size()-1){
				break;
			}
		}
		cout<<max<<endl;
	}
	return 0;
}
//要注意数据的范围，m*p可能超过int的范围，之前最后一个点不过，改为long long后过
```
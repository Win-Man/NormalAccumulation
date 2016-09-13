[Median ](https://www.patest.cn/contests/pat-a-practise/1029)

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<stdio.h>
using namespace std;
int main(){
	for(int n1;cin>>n1;){
		vector<long long>vec1(n1);
		for(int i = 0;i < n1;i++){
			scanf("%lld",&vec1[i]);
		}
		int n2;
		cin>>n2;
		vector<long long>vec2(n2);
		for(int i = 0; i < n2;i++){
			scanf("%lld",&vec2[i]);
		}
		int index1 = 0,index2 = 0,count = 0;
		int num;
		if((n1+n2)%2 == 0){
			num = (n1+n2)/2;
		}else{
			num = (n1+n2+1)/2;
		}
		while(1){
			if(index1 >= n1){
				index2++;
				count++;
				if(count == num){
					cout<<vec2[index2-1]<<endl;
					break;
				}
			}else if(index2 >= n2){
				index1++;
				count++;
				if(count == num){
					cout<<vec1[index1-1]<<endl;
					break;
				}
			}else{
				if(vec1[index1] < vec2[index2]){
					index1++;
					count++;
					if(count == num){
						cout<<vec1[index1-1]<<endl;
						break;
					}
				}else{
					index2++;
					count++;
					if(count == num){
						cout<<vec2[index2-1]<<endl;
						break;
					}
				}
			}
		}
		if((n1+n2) == 0){
			cout<<endl;
		}
	}
	return 0;
}
```
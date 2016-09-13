[¿ìËÙÅÅĞò](https://www.patest.cn/contests/pat-b-practise/1045)

``` c++
#include<iostream>
#include<stdio.h>
#include<string.h>
#include<string>
#include<vector>
#include<algorithm>
using namespace std;
int main(){
	for(int n;scanf("%d",&n)!=EOF;){
		vector<int>vec;
		vector<int>res;
		int *arr = new int[n];
		memset(arr,0,n*sizeof(int));
		int left_max = 0;
		int right_min = 1000000000;
		for(int i = 0;i < n;i++){
			int temp;
			scanf("%d",&temp);
			vec.push_back(temp);
		}
		for(int i = 0;i < vec.size();i++){
			if(vec[i] > left_max){
				arr[i]++;
				left_max = vec[i];
			}
		}
		for(int i = vec.size()-1;i >= 0;i--){
			if(vec[i] <= right_min){
				arr[i]++;
				right_min = vec[i];
			}
		}
		for(int i = 0; i < n;i++){
			if(arr[i] == 2){
				res.push_back(vec[i]);
			}
		}
		sort(res.begin(),res.end());
		printf("%d\n",res.size());
		if(res.size()){
			printf("%d",res[0]);
			for(int i = 1;i < res.size();i++){
				printf(" %d",res[i]);
			}
			printf("\n");
		}else{
			printf("\n");
		}
		delete[] arr;
	}
	return 0;
}
```
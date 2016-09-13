[Êý×ÖºÚ¶´ ](https://www.patest.cn/contests/pat-b-practise/1019)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<math.h>
#include<algorithm>
#include<stdio.h>
#include<vector>
using namespace std;
int main(){
	for(int n;cin>>n;){
		int num = n;
		while(1){
			vector<int>vec;
			int num1=0,num2=0;
			vec.push_back(num/1000);
			vec.push_back(num%1000/100);
			vec.push_back(num%100/10);
			vec.push_back(num%10);
			sort(vec.begin(),vec.end());
			num1 = vec[3] * 1000 + vec[2] * 100 + vec[1] * 10 + vec[0];
			num2 = vec[0] * 1000 + vec[1] * 100 + vec[2] * 10 + vec[3]; 
			num = num1 - num2;
			if(num1 == num2){
				printf("%d - %d = 0000\n",n,n);
				break;
			}else{
				printf("%d%d%d%d - %d%d%d%d = %d%d%d%d\n",num1/1000,num1%1000/100,num1%100/10,num1%10,
					num2/1000,num2%1000/100,num2%100/10,num2%10,num/1000,num%1000/100,num%100/10,num%10);
			}
			if(num == 6174){
				break;
			}
		}
	}
	return 0;
}
```
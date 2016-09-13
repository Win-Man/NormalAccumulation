[Magic Coupon ](https://www.patest.cn/contests/pat-a-practise/1037)

``` c++
#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;
int main(){
	for(int nc;cin>>nc;){
		vector<int>coupons_positive;
		vector<int>coupons_negative;
		vector<int>products_positive;
		vector<int>products_negative;
		for(int i = 0;i < nc;i++){
			int temp;
			cin>>temp;
			if(temp >= 0){
				coupons_positive.push_back(temp);
			}else{
				coupons_negative.push_back(temp);
			}
		}
		int np;
		cin>>np;
		for(int i = 0;i < np;i++){
			int temp;
			cin>>temp;
			if(temp >= 0){
				products_positive.push_back(temp);
			}else{
				products_negative.push_back(temp);
			}
		}
		sort(coupons_positive.begin(),coupons_positive.end());
		sort(coupons_negative.begin(),coupons_negative.end());
		sort(products_positive.begin(),products_positive.end());
		sort(products_negative.begin(),products_negative.end());
		int sum = 0;
		int i = coupons_positive.size()-1,j = products_positive.size()-1;
		while(i >= 0 && j >= 0){
			sum+= coupons_positive[i] * products_positive[j];
			i--;
			j--;
		}
		i = 0;
		j = 0;
		while(i < coupons_negative.size() && j < products_negative.size()){
			sum += coupons_negative[i] * products_negative[j];
			i++;
			j++;
		}
		cout<<sum<<endl;
	}
	return 0;
}
```
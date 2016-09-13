[月饼 ](https://www.patest.cn/contests/pat-b-practise/1020)

**AC code:**

``` c++
#include<iostream>
#include<iomanip>
#include<algorithm>
#include<vector>
struct moon{
	double stock;
	double price;
	double unit_price;
};
using namespace std;
bool cmp(const moon & temp1,const moon & temp2){
	if(temp1.unit_price > temp2.unit_price){
		return true;
	}else{
		return false;
	}
}
int main(){
	for(int n,num;cin>>n>>num;){
		vector<moon>vec(n);
		for(int i = 0;i < n;i++){
			cin>>vec[i].stock;
		}
		for(int i = 0;i < n;i++){
			cin>>vec[i].price;
			vec[i].unit_price = vec[i].price/vec[i].stock;
		}
		sort(vec.begin(),vec.end(),cmp);
		double sum = 0.0;
		for(int i =0;i < n;i++){
			if(num >= vec[i].stock){
				sum+= vec[i].price;
				num-= vec[i].stock;
			}else{
				sum+= num*vec[i].unit_price;
				break;
			}
		}
		cout<<fixed<<setprecision(2)<<sum<<endl;
	}
	return 0;
}
/*
库存 售价 单价要用double类型否则有一个测试点过不了
*/
```
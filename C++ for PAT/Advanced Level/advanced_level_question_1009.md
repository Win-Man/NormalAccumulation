[Product of Polynomials ](https://www.patest.cn/contests/pat-a-practise/1009)

``` c++
#include<iostream>
#include<vector>
#include<map>
#include<iomanip>
using namespace std;
int main(){
	for(int n1,n2;cin>>n1;){
		map<int,double>ploy1;
		map<int,double>ploy2;
		map<int,double>res;
		for(int i = 0;i < n1;i++){
			int m;
			cin>>m;
			cin>>ploy1[m];
		}
		cin>>n2;
		for(int i = 0;i < n2;i++){
			int m;
			cin>>m;
			cin>>ploy2[m];
		}
		map<int,double>::iterator iter1 = ploy1.begin();
		map<int,double>::iterator iter2 = ploy2.begin();
		for(;iter1 != ploy1.end();iter1++){
			for(iter2 = ploy2.begin();iter2 != ploy2.end();iter2++){
				res[iter1->first + iter2->first] += iter1->second*iter2->second;
			}
		}
		for(iter1 = res.begin();iter1 != res.end();){
			if(iter1->second == 0.0){
				iter1 = res.erase(iter1++);
				continue;
			}
			iter1++;
		}
		cout<<res.size();
		if(res.size()){
			iter1 = res.end();
			iter1--;
			for(;iter1 != res.begin();iter1--){
				if(iter1->second != 0.0){
					cout<<" "<<iter1->first<<" "<<fixed<<setprecision(1)<<iter1->second;
				}
			}
			iter1 = res.begin();
			if(iter1->second != 0.0){
				cout<<" "<<iter1->first<<" "<<fixed<<setprecision(1)<<iter1->second;
			}
		}
		cout<<endl;
	}
	return 0;
}
/*考虑结果中系数为零的情况要不输出*/
```
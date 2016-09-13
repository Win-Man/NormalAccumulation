[ÈË¿ÚÆÕ²é](https://www.patest.cn/contests/pat-b-practise/1028)

``` c++
#include<iostream>
#include<vector>
#include<algorithm>
#include<string>
#include<stdio.h>
using namespace std;
struct people{
	char name[6];
	int year;
	int month;
	int day;
};
bool cmp(const people & peo1,const people & peo2){
	if(peo1.year != peo2.year){
		return peo1.year < peo2.year;
	}else{
		if(peo1.month != peo2.month){
			return peo1.month < peo2.month;
		}else{
			return peo1.day < peo2.day;
		}
	}
}
int main(){
	for(int n;scanf("%d",&n) != EOF;){
		vector<people>reasonable;
		for(int i =0;i < n;i++){
			people temp;
			scanf("%s %d/%d/%d",temp.name,&temp.year,&temp.month,&temp.day);
			if(temp.year < 2014 && temp.year >1814 ||(temp.year == 2014 && temp.month <9)
					||(temp.year == 2014 && temp.month==9&&temp.day<=6)||(temp.year == 1814 && temp.month > 9)
					||(temp.year==1814&&temp.month==9&&temp.day>=6)){
				reasonable.push_back(temp);
			}
		}
		sort(reasonable.begin(),reasonable.end(),cmp);
		if(reasonable.size()){
			cout<<reasonable.size()<<" "<<reasonable[0].name<<" "<<reasonable[reasonable.size()-1].name<<endl;
		}else{
			cout<<"0"<<endl;
		}
		
	}
	return 0;
}
```
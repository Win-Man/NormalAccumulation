[Mars Numbers ](https://www.patest.cn/contests/pat-a-practise/1100)

``` c++
#include<iostream>
#include<string>
#include<sstream>
#include<stdio.h>
#include<map>
using namespace std;
int main(){
	/*地球人的0被火星人称为tret。
	地球人数字1到12的火星文分别为：jan, feb, mar, apr, may, jun, jly, aug, sep, oct, nov, dec。
	火星人将进位以后的12个高位数字分别称为：tam, hel, maa, huh, tou, kes, hei, elo, syy, lok, mer, jou。*/
	string arr1[13]={"tret","jan","feb","mar","apr","may","jun","jly","aug","sep","oct","nov","dec"};
	string arr2[12]={"tam","hel","maa","huh","tou","kes","hei","elo","syy","lok","mer","jou"};
	map<int,string>low;
	map<int,string>high;
	for(int i = 0;i < 13;i++){
		low[i] = arr1[i];
	}
	for(int i =0;i < 12;i++){
		high[i] = arr2[i];
	}
	for(int n;scanf("%d\n",&n)!=EOF;){
		for(int i = 0;i < n;i++){
			string str;
			getline(cin,str);
			if(str[0] >= '0' && str[0] <= '9'){
				int num;
				stringstream ss;
				ss.clear();
				ss.str(str);
				while(1){
					ss>>num;
					if(ss.fail()){
						break;
					}
				}
				if(num >= 13 && num%13 != 0){
					cout<<high[num/13 - 1]<<" "<<low[num%13]<<endl;
				}else if(num >= 13 && num%13 ==0){
					cout<<high[num/13 - 1]<<endl;
				}else{
					cout<<low[num]<<endl;
				}
			}else{
				string mars[2];
				int index = 0;
				stringstream ss;
				ss.clear();
				ss.str(str);
				while(1){
					ss>>mars[index];
					if(ss.fail()){
						break;
					}
					index++;
				}
				int sum = 0;
				map<int,string>::iterator iter;
				for(iter = high.begin();iter != high.end();iter++){
					if(iter->second == mars[0]){
						sum+=(iter->first+1)*13;
						break;
					}
				}
				for(iter = low.begin();iter != low.end();iter++){
					if(iter->second == mars[0]){
						sum+=(iter->first);
						break;
					}
				}
				for(iter = low.begin();iter != low.end();iter++){
					if(iter->second == mars[1]){
						sum+=(iter->first);
						break;
					}
				}
				cout<<sum<<endl;
			}
		}
	}
	return 0;
}
```
[Longest Symmetric String ](https://www.patest.cn/contests/pat-a-practise/1040)

``` c++
#include<iostream>
#include<string>
using namespace std;
int main(){
	for(string str;getline(cin,str);){
		int max_length = 1;
		for(int i = 0; i < str.length();i++){
			int length1 = 1;
			int length2 = 1;
			//判断长度为偶数的情况
			int index_1 = i;
			int index_2 = i+1;
			if(index_1 < 0 || index_2 >= str.length()){
				continue;
			}
			if(str[index_1] == str[index_2]){
				length1 = 2;
				while(1){
					index_1--;
					index_2++;
					if(index_1 < 0  || index_2 >= str.length()){
						break;
					}
					if(str[index_1] == str[index_2]){
						length1+=2;
					}else{
						break;
					}
				}
			}
			//判断长度为奇数的情况
			index_1 = i-1;
			index_2 = i+1;
			if(index_1 < 0 || index_2 >= str.length()){
				continue;
			}
			if(str[index_1] == str[index_2]){
				length2 = 3;
				while(1){
					index_1--;
					index_2++;
					if(index_1 < 0 || index_2 >= str.length()){
						break;
					}
					if(str[index_1] == str[index_2]){
						length2+= 2;
					}else{
						break;
					}
				}
			}
			if(length1 > max_length){
				max_length = length1;
			}
			if(length2 > max_length){
				max_length = length2;
			}
		}
		cout<<max_length<<endl;
	}
	return 0;
}
```
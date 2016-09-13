[科学计数法 ](https://www.patest.cn/contests/pat-b-practise/1024)

``` c++
#include<iostream>
#include<string>
#include<vector>
#include<math.h>
using namespace std;
int main(){
	for(string str;cin>>str;){
		string num = str.substr(1,str.find_first_of('E')-1);
		string index = str.substr(str.find_first_of('E')+2,str.length()-str.find_first_of('E')-2);
		int pos_point = num.find_first_of('.');
		int index_num=0;
		if(str[0] == '-'){
			cout<<'-';
		}
		for(int i =0;i < index.length();i++){
			index_num+= pow(10,index.length() - i - 1)*(index[i] - '0');
		}
		char index_flag = str[str.find_first_of('E')+1];
		if(index_flag == '-'){
			for(int i = 0;i < index_num;i++){
				if(pos_point == 0){
					num[0] = '0';
					num = "." + num;
				}else{
					num[pos_point] = num[pos_point - 1];
					num[pos_point - 1] = '.';
					pos_point--;
				}
			}
			if(pos_point){
				cout<<num<<endl;
			}else{
				cout<<"0"<<num<<endl;
			}
		}else if(index_flag == '+'){
			for(int i =0;i < index_num;i++){
				if(pos_point < num.length() - 2){
					num[pos_point] = num[pos_point + 1];
					num[pos_point + 1] = '.';
					pos_point++;
				}else if(pos_point == num.length() - 2){
					num.erase(pos_point,1);
					pos_point = index_num + num.length();
				}else{
					num = num + "0";
				}
			}
			cout<<num<<endl;
		}
	}
	return 0;
}
```
[Sign In and Sign Out ](https://www.patest.cn/contests/pat-a-practise/1006)

``` c++
#include<iostream>
#include<stdio.h>
#include<string>
#include<vector>
#include<algorithm>
using namespace std;
struct person{
	string ID;
	int hour,minute,second;
};
bool cmp(const person& p1,const person& p2){
	if(p1.hour != p2.hour){
		return p1.hour < p2.hour;
	}else{
		if(p1.minute != p2.minute){
			return p1.minute < p2.minute;
		}else{
			return p1.second < p2.second;
		}
	}
}
int main(){
	for(int n;cin>>n;){
		vector<person>sign_in;
		vector<person>sign_out;
		for(int i = 0;i < n;i++){
			person in,out;
			string id ;
			cin>>id;
			in.ID = id;
			out.ID = id;
			scanf("%d:%d:%d",&in.hour,&in.minute,&in.second);
			scanf("%d:%d:%d",&out.hour,&out.minute,&out.second);
			sign_in.push_back(in);
			sign_out.push_back(out);
		}
		sort(sign_in.begin(),sign_in.end(),cmp);
		sort(sign_out.begin(),sign_out.end(),cmp);
		cout<<sign_in[0].ID<<" "<<sign_out[sign_out.size()-1].ID<<endl;
	}
	return 0;
}
```
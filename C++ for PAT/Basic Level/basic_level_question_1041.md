[¿¼ÊÔ×ùÎ»ºÅ](https://www.patest.cn/contests/pat-b-practise/1041)

``` c++
#include<iostream>
#include<string>
#include<vector>
using namespace std;
struct student{
	string kaohao;
	int shiji;
	int kaoshi;
};
int main(){
	for(int n;cin>>n;){
		vector<student>students(1005);
		for(int i = 0;i < n;i++){
			student temp;
			cin>>temp.kaohao>>temp.shiji>>temp.kaoshi;
			students[temp.shiji]=temp;
		}
		int m;
		cin>>m;
		for(int i =0;i < m;i++){
			int query;
			cin>>query;
			cout<<students[query].kaohao<<" "<<students[query].kaoshi<<endl;
		}
	}
	return 0;
}
```
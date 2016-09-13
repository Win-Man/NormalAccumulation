[成绩排名 ](https://www.patest.cn/contests/pat-b-practise/1004)

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<string>
#include<algorithm>
using namespace std;
struct student{
	string name;
	string xuehao;
	int grade;
};
bool cmp(const student & temp1,const student & temp2){
	if(temp1.grade < temp2.grade)
		return true;
	else
		return false;
}
int main(){
	int n;
	cin>>n;
	vector<student>temp;
	for(int i = 0;i<n;i++){
		student haha;
		cin>>haha.name>>haha.xuehao>>haha.grade;
		temp.push_back(haha);
	}
	sort(temp.begin(),temp.end(),cmp);
	cout<<temp[n-1].name<<" "<<temp[n-1].xuehao<<endl;
	cout<<temp[0].name<<" "<<temp[0].xuehao<<endl;
	return 0;
}
/*
写比较函数的时候，参数要写成const类型的引用，否则会出现编译错误
*/
```
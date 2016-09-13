[德才论 ](https://www.patest.cn/contests/pat-b-practise/1015)

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<algorithm>
#include<string>
using namespace std;
struct student{
	int kaohao;
	int de;
	int cai;
};
bool cmp(const student & stu1,const student & stu2){
	if(stu1.cai + stu1.de != stu2.de +stu2.cai){
		return stu1.cai + stu1.de < stu2.de +stu2.cai;
	}else{
		if(stu1.de != stu2.de){
			return stu1.de < stu2.de;
		}else{
			if(stu1.kaohao > stu2.kaohao){
				return true;
			}else{
				return false;
			}
		}
	}
	
}
int main(){
	for(int n,l,h;cin>>n>>l>>h;){
		vector<student>first;
		vector<student>second;
		vector<student>third;
		vector<student>forth;
		for(int i = 0;i < n;i++){
			student temp;
			cin>>temp.kaohao>>temp.de>>temp.cai;
			if(temp.de >= h && temp.cai >= h){
				first.push_back(temp);
			}else if(temp.de >= h && temp.cai < h && temp.cai >= l){
				second.push_back(temp);
			}else if(temp.cai >= l && temp.de >= l && temp.de >= temp.cai){
				third.push_back(temp);
			}else if(temp.cai >= l && temp.de >= l){
				forth.push_back(temp);
			}
		}
		if(first.size()){
			sort(first.begin(),first.end(),cmp);
		}
		if(second.size()){
			sort(second.begin(),second.end(),cmp);
		}
		if(third.size()){
			sort(third.begin(),third.end(),cmp);
		}
		if(forth.size()){
			sort(forth.begin(),forth.end(),cmp);
		}
		cout<<first.size() + second.size() + third.size() + forth.size()<<endl;
		for(int i = 0;i<first.size();i++){
			/*cout<<first[first.size()-i-1].kaohao<<" "<<first[first.size()-i-1].de
				<<" "<<first[first.size()-i-1].cai<<endl;*/
			printf("%d %d %d\n",first[first.size()-i-1].kaohao,first[first.size()-i-1].de,
				first[first.size()-i-1].cai);
		}
		for(int i = 0;i<second.size();i++){
			/*cout<<second[second.size()-i-1].kaohao<<" "<<second[second.size()-i-1].de
				<<" "<<second[second.size()-i-1].cai<<endl;*/
			printf("%d %d %d\n",second[second.size()-i-1].kaohao,second[second.size()-i-1].de,
				second[second.size()-i-1].cai);
		}
		for(int i = 0;i<third.size();i++){
			/*cout<<third[third.size()-i-1].kaohao<<" "<<third[third.size()-i-1].de
				<<" "<<third[third.size()-i-1].cai<<endl;*/
			printf("%d %d %d\n",third[third.size()-i-1].kaohao,third[third.size()-i-1].de,
				third[third.size()-i-1].cai);
		}
		for(int i = 0;i<forth.size();i++){
			/*cout<<forth[forth.size()-i-1].kaohao<<" "<<forth[forth.size()-i-1].de
				<<" "<<forth[forth.size()-i-1].cai<<endl;*/
			printf("%d %d %d\n",forth[forth.size()-i-1].kaohao,forth[forth.size()-i-1].de,
				forth[forth.size()-i-1].cai);
		}
	}
	return 0;
}
/*
输出要使用printf不然会超时
注意cmp函数的书写，不然会有operator<重载的错误
*/
```
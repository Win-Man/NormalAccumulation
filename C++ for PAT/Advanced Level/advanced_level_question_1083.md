[List Grades ](https://www.patest.cn/contests/pat-a-practise/1083)

``` c++
#include<iostream>
#include<vector>
#include<string>
#include<algorithm>
using namespace std;
struct student{
	string name;
	string ID;
	int grade;
};
bool cmp(const student& stu1,const student& stu2){
	if(stu1.grade < stu2.grade){
		return true;
	}else{
		return false;
	}
}
int main(){
	for(int n;cin>>n;){
		vector<student>students;
		for(int i = 0;i < n; i++){
			student temp;
			cin>>temp.name>>temp.ID>>temp.grade;
			students.push_back(temp);
		}
		sort(students.begin(),students.end(),cmp);
		int min,max;
		cin>>min>>max;
		int count = 0;
		for(int i = students.size()-1;i > -1;i--){
			if(students[i].grade >= min && students[i].grade <= max){
				cout<<students[i].name<<" "<<students[i].ID<<endl;
				count++;
			}
		}
		if(!count){
			cout<<"NONE"<<endl;
		}
	}
	return 0;
}
```
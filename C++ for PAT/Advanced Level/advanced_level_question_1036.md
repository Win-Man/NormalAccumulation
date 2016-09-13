[Boys vs Girls ](https://www.patest.cn/contests/pat-a-practise/1036)

``` c++
#include<iostream>
#include<string>
#include<vector>
#include<algorithm>
using namespace std;
struct student{
	string name,ID;
	char sex;
	int grade;
};
bool cmp(const student& stu1,const student& stu2){
	return stu1.grade < stu2.grade;
}
int main(){
	for(int n;cin>>n;){
		vector<student>female;
		vector<student>male;
		for(int i = 0;i < n;i++){
			student temp;
			cin>>temp.name>>temp.sex>>temp.ID>>temp.grade;
			switch (temp.sex)
			{
			case 'F':
				female.push_back(temp);
				break;
			case 'M':
				male.push_back(temp);
				break;
			default:
				break;
			}
		}
		sort(female.begin(),female.end(),cmp);
		sort(male.begin(),male.end(),cmp);
		if(female.size() == 0){
			cout<<"Absent"<<endl;
			cout<<male[0].name<<" "<<male[0].ID<<endl;
			cout<<"NA"<<endl;
		}else if(male.size() == 0){
			cout<<female[female.size()-1].name<<" "<<female[female.size()-1].ID<<endl;
			cout<<"Absent"<<endl<<"NA"<<endl;
		}else{
			cout<<female[female.size()-1].name<<" "<<female[female.size()-1].ID<<endl;
			cout<<male[0].name<<" "<<male[0].ID<<endl;
			cout<<female[female.size()-1].grade - male[0].grade<<endl;
		}
	}
	return 0;
}
```
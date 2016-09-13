[List Sorting ](https://www.patest.cn/contests/pat-a-practise/1028)

``` c++
#include<iostream>
#include<vector>
#include<string.h>
#include<algorithm>
#include<stdio.h>
using namespace std;
struct student{
	int ID;
	char name[9];
	int grade;
};
bool cmp1(const student& stu1,const student& stu2){
	return stu1.ID < stu2.ID;
}
bool cmp2(const student& stu1,const student& stu2){
	if(strcmp(stu1.name,stu2.name) != 0){
		if(strcmp(stu1.name,stu2.name) < 0){
			return true;
		}else{
			return false;
		}
	}else{
		return stu1.ID < stu2.ID;
	}
}
bool cmp3(const student& stu1,const student& stu2){
	if(stu1.grade != stu2.grade){
		return stu1.grade < stu2.grade;
	}else{
		return stu1.ID < stu2.ID;
	}
}
int main(){
	for(int n,c;scanf("%d%d",&n,&c)!=EOF;){
		vector<student>students;
		for(int i = 0;i < n;i++){
			student temp;
			scanf("%d%s%d",&temp.ID,temp.name,&temp.grade);
			students.push_back(temp);
		}
		if(c == 1){
			sort(students.begin(),students.end(),cmp1);
		}else if(c == 2){
			sort(students.begin(),students.end(),cmp2);
		}else if(c == 3){
			sort(students.begin(),students.end(),cmp3);
		}
		for(int i = 0;i < students.size();i++){
			printf("%06d %s %d\n",students[i].ID,students[i].name,students[i].grade);
		}
	}
	return 0;
}
/*需要用scanf cin会超时*/
```
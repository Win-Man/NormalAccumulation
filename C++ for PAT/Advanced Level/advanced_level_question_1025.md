[PAT Ranking ](https://www.patest.cn/contests/pat-a-practise/1025)

``` c++
#include<iostream>
#include<vector>
#include<string.h>
#include<algorithm>
#include<stdio.h>
#include<vector>
using namespace std;
struct student{
	char ID[15];
	int grade;
	int location;
	int location_rank;
};
bool cmp(const student&stu1,const student&stu2){
	if(stu1.grade == stu2.grade){
		if(strcmp(stu1.ID,stu2.ID) > 0){
			return true;
		}else{
			return false;
		}
	}else{
		return stu1.grade < stu2.grade;
	}
}
int main(){
	for(int n;scanf("%d",&n)!=EOF;){
		vector<student>vec_all;
		for(int i = 1;i <= n;i++){
			int k;
			scanf("%d",&k);
			vector<student>vec_local;
			while(k--){
				student temp;
				scanf("%s%d",temp.ID,&temp.grade);
				vec_local.push_back(temp);
			}
			sort(vec_local.begin(),vec_local.end(),cmp);
			//grade_flag ���ڼ�¼��һ��ѧ���ĳɼ���flag���ڼ�¼��һ��ѧ��������,�ɼ���ͬ��ѧ������һ����
			int grade_flag = vec_local[vec_local.size()-1].grade;
			int flag = 1;
			for(int j = vec_local.size()-1;j >= 0;j--){
				vec_local[j].location = i;
				//������ѧ������һ��ѧ���ĳɼ�һ���������ѧ������������һ��ѧ��������һ����
				if(vec_local[j].grade == grade_flag){
					vec_local[j].location_rank = flag;
				}else{
					//������ѧ���ĳɼ�����һ��ѧ���ĳɼ���һ������˵������������ǰ��ѧ��������+1
					grade_flag = vec_local[j].grade;
					vec_local[j].location_rank = vec_local.size() - j;
					flag = vec_local[j].location_rank;
				}
				vec_all.push_back(vec_local[j]);
			}
		}
		sort(vec_all.begin(),vec_all.end(),cmp);
		printf("%d\n",vec_all.size());
		int grade_flag = vec_all[vec_all.size()-1].grade;
		int flag = 1;
		for(int j = vec_all.size()-1;j >= 0;j--){
			printf("%s ",vec_all[j].ID);
			if(vec_all[j].grade == grade_flag){
				printf("%d ",flag);
			}else{
				grade_flag = vec_all[j].grade;
				flag = vec_all.size() - j;
				printf("%d ",flag);
			}
			printf("%d %d\n",vec_all[j].location,vec_all[j].location_rank);
		}
		
	}
	return 0;
}
```
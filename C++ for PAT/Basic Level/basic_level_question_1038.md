[统计同成绩学生](https://www.patest.cn/contests/pat-b-practise/1038)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<stdio.h>
using namespace std;
int main(){
	for(int n;scanf("%d",&n)!=EOF;){
		int grades[105]={0};
		for(int i = 0;i < n;i++){
			int stu;
			scanf("%d",&stu);
			grades[stu]++;
		}
		int m ;
		scanf("%d",&m);
		for(int i = 0;i < m-1;i++){
			int temp;
			scanf("%d",&temp);
			printf("%d ",grades[temp]);
		}
		scanf("%d",&m);
		printf("%d\n",grades[m]);
	}
	return 0;
}
```
[Course List for Student](https://www.patest.cn/contests/pat-a-practise/1039)


**解题思路：**
一开始以为是一题挺简单的题目，用一下c++中的map很轻松就能过吧，写了一下，发现最后一个case不能过，各种简单的优化都上了，还是没有用，最后网上看了下别人的解题思路，发现说，用map进行查询的时候检索名字比较慢，题目中给定名字的组成是三个大写字母加一个数字，总共26*26*26*10中可能性，用哈希来表示一下名字的话，索引的时候查询效率就变成O（1）了

**AC code:**

``` c++
#include<iostream>  
#include<map>  
#include<vector>  
#include<string>  
#include<stdio.h>  
#include<algorithm>  
using namespace std;  
int main(){  
    for (int query_num, course_num; scanf("%d%d",&query_num,&course_num)!=EOF;){  
        //构建名字的哈希表  
        vector<vector<int>>names(26*26*26*10 + 5);  
        while (course_num--){  
            int course, stu_num;  
            //cin >> course >> stu_num;  
            scanf("%d%d", &course, &stu_num);  
            char name[5];  
            while (stu_num--){  
                scanf("%s", name);  
                int index = (name[0] - 'A') * 26 * 26 * 10 + (name[1] - 'A') * 26 * 10 + (name[2] - 'A') * 10 + (name[3] - '0');  
                names[index].push_back(course);  
            }  
        }  
        while (query_num--){  
            char query_name[5];  
            scanf("%s", query_name);  
            int index = (query_name[0] - 'A') * 26 * 26 * 10 + (query_name[1] - 'A') * 26 * 10 + (query_name[2] - 'A') * 10 + (query_name[3] - '0');  
            sort(names[index].begin(),names[index].end());  
            printf("%s %d", query_name, names[index].size());  
            for (int i = 0; i < names[index].size(); i++){  
                printf(" %d", names[index][i]);  
            }  
            printf("\n");  
        };  
    }  
    return 0;  
}  
```
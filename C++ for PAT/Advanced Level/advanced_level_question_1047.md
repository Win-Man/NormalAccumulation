[Student List for Course](https://www.patest.cn/contests/pat-a-practise/1047)


**解题思路：**
上一题是有信息输出学生的情况，这一题是由信息输出课程的情况，一开始用string来表示学生的姓名，结果最后一个case又超时了，然后我就试着把用string表示名字再变成用数字来表示名字，用名字转换为哈希表的下标来表示名字，过了。最后输出之前，数字的排序比string的排序快吧。

**AC code:**

``` c++
#include<iostream>  
#include<vector>  
#include<map>  
#include<stdio.h>  
#include<algorithm>  
#include<string>  
using namespace std;  
int main(){  
    for (int student_num, course_num; scanf("%d%d", &student_num, &course_num) != EOF;){  
        vector<vector<int>>stu_list(course_num+1);  
        while (student_num--){  
            char name[5];  
            scanf("%s", name);  
            int index = (name[0] - 'A') * 26 * 26 * 10 + (name[1] - 'A') * 26 * 10 + (name[2] - 'A') * 10 + (name[3] - '0');  
            int register_num;  
            scanf("%d", &register_num);  
            int course_id;  
            while (register_num--){  
                scanf("%d", &course_id);  
                stu_list[course_id].push_back(index);  
            }  
        }  
        for (int i = 1; i < course_num + 1; i++){  
            cout << i << " " << stu_list[i].size() << endl;  
            sort(stu_list[i].begin(), stu_list[i].end());  
            for (int j = 0; j < stu_list[i].size(); j++){  
                printf("%c%c%c%c\n", stu_list[i][j] / (26 * 26 * 10) + 'A',  
                    (stu_list[i][j]%(26*26*10))/(26*10) + 'A',  
                    (stu_list[i][j]%(26*10))/10 +'A',  
                    (stu_list[i][j]%(10)) + '0');  
            }  
        }  
    }  
    return 0;  
}  

```
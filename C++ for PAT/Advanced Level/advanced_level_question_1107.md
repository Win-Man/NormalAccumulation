[Social Clusters](https://www.patest.cn/contests/pat-a-practise/1107)


**解题思路：**
这题考察的内容是并查集，还是比较简单的，只要基本掌握并查集便可。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<algorithm>
#include<string.h>
using namespace std;
int people[1005] = { 0 };
vector<vector<int>>hobbies(1005);
int find(int x){
    int temp = x;
    while (people[temp] != temp){
        temp = people[temp];
    }
    //压缩路径
    while (people[x] != x){
        people[x] = temp;
        x = people[x];
    }
    return temp;
}
void join(int x, int y){
    int tx = find(x);
    int ty = find(y);
    if (tx != ty){
        people[tx] = ty;
    }
}
int main(){
    for (int n; scanf("%d", &n) != EOF;){       
        for (int i = 1; i <= n; i++){
            //顺便初始化人员的标记数组
            people[i] = i;
            int k;
            scanf("%d:",&k);
            while (k--){
                int temp;
                scanf("%d", &temp);
                hobbies[temp].push_back(i);
            }
        }
        for (int i = 1; i < hobbies.size(); i++){
            if (hobbies[i].size()){
                int x = hobbies[i][0];
                for (int j = 1; j < hobbies[i].size(); j++){
                    int y = hobbies[i][j];
                    join(x, y);
                }
            }
        }
        for (int i = 1; i <= n; i++){
            find(i);
        }
        int ans[1005] = { 0 };
        for (int i = 1; i <= n; i++){
            ans[people[i]]++;
        }
        vector<int>group;
        for (int i = 1; i <= n; i++){
            if (ans[i]){
                group.push_back(ans[i]);
            }
        }
        printf("%d\n", group.size());
        sort(group.begin(), group.end());
        printf("%d", group[group.size() - 1]);
        for (int i = 1; i < group.size(); i++){
            printf(" %d", group[group.size() - 1 - i]);
        }
        printf("\n");
    }
    return 0;
}
```
[Mice and Rice](https://www.patest.cn/contests/pat-a-practise/1056)


**解题思路：**
这道题的题意有点混乱，第一行输入的是老鼠的数量，和每次匹配的数量，第二行输入的是编号0——n-1号老鼠的重量，第三行输入的是进行比赛的老鼠的编号的排列顺序。样例输入中第三行是6 0 8 7 10 5 9 1 4 2 3，就是6号，0号，8号老鼠一组。7号，10号，5号老鼠一组……，这题意混乱了我好久。看懂题目之后就比较简单了，模拟淘汰的顺序就好，只是在淘汰的过程中需要对老鼠进行排名，在某一轮中淘汰的老鼠的rank就是这一轮中晋级的老鼠的数量+1.

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
using namespace std;
int main(){
    for (int p, g; scanf("%d%d", &p, &g) != EOF;){
        int *mice = new int[p];
        int *rank = new int[p];
        for (int i = 0; i < p; i++){
            rank[i] = 1;
            scanf("%d", &mice[i]);
        }
        vector<int>compete;
        for (int i = 0; i < p; i++){
            int temp;
            scanf("%d", &temp);
            compete.push_back(temp);
        }
        while (compete.size() != 1){
            vector<int>next;
            int nextCount = compete.size() / g;
            if (compete.size() % g != 0){
                nextCount++;
            }
            for (int i = 0; i < compete.size(); i += g){
                int win = -1;
                int fatest = -1;
                for (int j = i; j < i+g && j < compete.size(); j++){
                    if (mice[compete[j]] > fatest){
                        win =compete[j];
                        fatest = mice[compete[j]];
                    }
                }
                if (win != -1){
                    next.push_back(win);
                    rank[win]--;
                }
                for (int j = i; j < i + g && j < compete.size(); j++){
                    if (compete[j] != win){
                        rank[compete[j]] = nextCount+1;
                    }
                }
            }
            compete = vector<int>(next);
        }
        rank[compete[0]] = 1;
        printf("%d", rank[0]);
        for (int i = 1; i < p; i++){
            printf(" %d", rank[i]);
        }
        printf("\n");
        delete[] mice;
        delete[] rank;
    }
    return 0;
}
```
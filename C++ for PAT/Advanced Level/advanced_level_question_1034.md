[Head of a Gang](https://www.patest.cn/contests/pat-a-practise/1034)


**解题思路：**
一道并查集的题目，名字用hash的办法变成int类型，所以开个26*26*26的s数组就行。用一个数组记录每个人总的通话时间，用于最后找出head。并查集找出团队，再在团队中找到通话时间最长的人就是结果。因为结果需要按字母序输出，所以我先将结果保存下来，排序之后再输出的。 
做题的时候并查集的压缩路径的地方写错了，导致测试点2 4一直答案错误，还是太粗心。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<string.h>
#include<vector>
#include<algorithm>
using namespace std;
int group[26 * 26 * 26];
int ti[26 * 26 * 26];
vector<int> groupCount[26 * 26 * 26];
struct res{
    int name;
    int size;
};
bool cmp(const res&r1, const res&r2){
    return r1.name < r2.name;
}
int find(int x){
    int res = x;
    while (group[res] != res){
        res = group[res];
    }
    while (x != res){
        int temp = group[x];
        group[x] = res;
        x = temp;
    }
    return res;
}
void join(int x, int y){
    int fx = find(x);
    int fy = find(y);
    if (fx != fy){
        group[fy] = fx;
    }
}
int main(){
    for (int n, k; scanf("%d%d", &n, &k) != EOF;){
        for (int i = 0; i < 26 * 26 * 26; i++){
            group[i] = i;
            ti[i] = 0;
            groupCount[i] = vector<int>();
        }
        for (int i = 0; i < n; i++){
            char p1[4];
            char p2[4];
            int temp;
            scanf("%s %s %d", p1, p2, &temp);
            int index1 = (p1[0] - 'A') * 26 * 26 + (p1[1] - 'A') * 26 + (p1[2] - 'A');
            int index2 = (p2[0] - 'A') * 26 * 26 + (p2[1] - 'A') * 26 + (p2[2] - 'A');
            join(index1, index2);
            ti[index1] += temp;
            ti[index2] += temp;
        }
        for (int i = 0; i < 26 * 26 * 26; i++){
            find(i);
        }
        for (int i = 0; i < 26 * 26 * 26; i++){
            if (ti[i]>0){
                groupCount[group[i]].push_back(i);
            }
        }
        vector<res>out;
        for (int i = 0; i < 26 * 26 * 26; i++){
            if (groupCount[i].size() > 2){
                int name = groupCount[i][0];
                int maxTime = ti[groupCount[i][0]];
                int totalTime = 0;
                for (int j = 0; j < groupCount[i].size(); j++){
                    totalTime += ti[groupCount[i][j]];
                    if (ti[groupCount[i][j]] >= maxTime){
                        maxTime = ti[groupCount[i][j]];
                        name = groupCount[i][j];
                    }
                }
                if (totalTime > k*2){
                    res temp;
                    temp.name = name;
                    temp.size = groupCount[i].size();
                    out.push_back(temp);
                }
            }
        }
        sort(out.begin(), out.end(), cmp);
        printf("%d\n", out.size());
        for (int i = 0; i < out.size(); i++){
            printf("%c%c%c %d\n", (char)(out[i].name / (26 * 26) + 'A'), (char)((out[i].name % (26 * 26)) / 26 + 'A'), (char)(out[i].name % 26 + 'A'), out[i].size);
        }
    }
    return 0;
}
```
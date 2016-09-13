[PAT Judge](https://www.patest.cn/contests/pat-a-practise/1075)


**解题思路：**
比较简单，但是比较繁琐的一道题，因为要考虑的情况很多，一直最后测试点过不了，理逻辑理了好久。

tips：

- 首先是没有提交和虽然提交过但是所有提交都是未通过编译的用户不用显示 
- 一道题如果只被未通过编译的提交过，那么最后显示的分数应该是0 
- 一道题如果没有被提交过显示的才是- 
- 一道题如果已经通过编译提交过，后面可能还会有未通过编译的提交，需要考虑，不能覆盖结果 
- 多次满分提交的情况也需要考虑。

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;
struct user{
    //id
    int id;
    //获取到的总分
    int score = 0;
    //获取满分的题数
    int per_solve = 0;
    //编译通过的提交数
    int pass = 0;
    //每道题获取到的分数
    int obtained[6] ;
};
bool cmp(const user u1, const user u2){
    if (u1.score > u2.score){
        return true;
    }
    else if (u1.score == u2.score){
        if (u1.per_solve > u2.per_solve){
            return true;
        }
        else if(u1.per_solve ==  u2.per_solve){
            if (u1.id < u2.id){
                return true;
            }
            else {
                return false;
            }
        }
        else{
            return false;
        }
    }
    else{
        return false;
    }
}
int main(){
    for (int n, k, m; scanf("%d%d%d", &n, &k, &m) != EOF;){
        int* arr = new int[k+1];
        for (int i = 1; i <= k; i++){
            scanf("%d", &arr[i]);
        }
        user* users = new user[n + 1];
        for (int i = 1; i <= n; i++){
            for (int j = 1; j <= k; j++){
                users[i].obtained[j] = -1;
            }
        }
        while (m--){
            int id;
            int problem;
            int get_score;
            scanf("%d%d%d", &id, &problem, &get_score);
            users[id].id = id;
            if (get_score != -1){
                users[id].pass++;
                if (get_score > users[id].obtained[problem]){
                    //users[id].score = users[id].score - users[id].obtained[problem] + get_score;
                    users[id].obtained[problem] = get_score;
                    //可能会有多次满分的提交
                    if (get_score == arr[problem]){
                        users[id].per_solve++;
                    }
                }
            }
            else{
                //之前通过编译的提交不能被后面未通过编译的提交覆盖
                if (users[id].obtained[problem] == -1)
                    users[id].obtained[problem] = 0;
            }
        }
        vector<user>comp;
        for (int i = 1; i <= n; i++){
            if (users[i].id == i && users[i].pass){
                for (int j = 1; j <= k; j++){
                    if (users[i].obtained[j]!= -1)
                        users[i].score += users[i].obtained[j];
                }
                comp.push_back(users[i]);
            }
        }
        sort(comp.begin(), comp.end(),cmp);
        int rank = 1;
        int per_score = comp[0].score;
        for (int i = 0; i < comp.size(); i++){
            if (comp[i].score < per_score){
                per_score = comp[i].score;
                rank = i + 1;
            }
            printf("%d %05d %d", rank, comp[i].id, comp[i].score);
            for (int j = 1; j <= k; j++){
                if (comp[i].obtained[j] != -1){
                    printf(" %d", comp[i].obtained[j]);
                }
                else{
                    printf(" -");
                }
            }
            printf("\n");
        }
    }
    return 0;
}
```
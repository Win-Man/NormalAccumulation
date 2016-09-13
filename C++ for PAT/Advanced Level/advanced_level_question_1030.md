[Travel Plan](https://www.patest.cn/contests/pat-a-practise/1030)


**解题思路：**
简单的最短路径问题，需要打印最短路径的dijkstra问题，只是需要用到两个二维数组分别记录，城市之间的距离和花费。还用了一个fa数组，用来记录，到当前城市的最短路径的上一个城市编号。用于最后方便输出路径。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
using namespace std;
#define INF 999999
int mp[505][505];
int mpc[505][505];
int flag[505];
int dist[505];
int cost[505];
int fa[505];
int n, m, s, d;
void dijkstra(int s){
    dist[s] = 0;
    cost[s] = 0;
    while (1){
        int u, mdist = INF;
        for (int i = 0; i < n; i++){
            if (flag[i] == 0 && mdist > dist[i]){
                mdist = dist[i];
                u = i;
            }
        }
        if (mdist == INF){
            break;
        }
        flag[u] = 1;
        for (int i = 0; i < n; i++){
            if (flag[i] == 0){
                if (dist[i] > dist[u] + mp[u][i]){
                    dist[i] = dist[u] + mp[u][i];
                    cost[i] = cost[u] + mpc[u][i];
                    fa[i] = u;
                }
                else if (dist[i] == dist[u] + mp[u][i]){
                    if (cost[i] > cost[u] + mpc[u][i]){
                        cost[i] = cost[u] + mpc[u][i];
                        fa[i] = u;
                    }
                }
            }
        }
    }
}
int main(){
    for (; scanf("%d%d%d%d", &n, &m, &s, &d) != EOF;){
        for (int i = 0; i < n; i++){
            dist[i] = INF;
            cost[i] = INF;
            flag[i] = 0;
            for (int j = 0; j < n; j++){
                mp[i][j] = INF;
                mpc[i][j] = INF;
            }
        }
        for (int i = 0; i < m; i++){
            int c1, c2, d, c;
            scanf("%d%d%d%d", &c1, &c2, &d, &c);
            mp[c1][c2] = mp[c2][c1] = d;
            mpc[c1][c2] = mpc[c2][c1] = c;
        }
        dijkstra(s);
        int temp = d;
        vector<int>res;
        while (temp != s){
            res.push_back(temp);
            temp = fa[temp];
        }
        res.push_back(s);
        for (int i = res.size()-1; i >= 0; i--){
            printf("%d ", res[i]);
        }
        printf("%d %d\n", dist[d], cost[d]);
    }
    return 0;
}
```
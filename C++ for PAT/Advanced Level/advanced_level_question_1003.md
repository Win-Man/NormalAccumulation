[Emergency](https://www.patest.cn/contests/pat-a-practise/1003)


**解题思路：**
最短路径的加强版，需要记录最短路径的数量，及最大权重，根据别人的代码写出来的。

**AC code:**

``` c++
#include<iostream>
using namespace std;
#define INF 999999
#define MX 501
//记录两个点之间的距离
int mp[MX][MX];
//标记各个点是否被访问过
int v[MX];
//记录从起始点到点i的距离
int dist[MX];
//记录从起始点到当前点的最大队伍数量
int amount[MX];
//记录每个城市的队伍数量
int teams[MX];
//记录从起始点到当前点的最短路径数量
int pathcount[MX];
int N, M, start, en;
void dijkstra(int s){
    //
    amount[s] = teams[s];
    dist[s] = 0;
    pathcount[s] = 1;

    while (1){
        int u,dmin = INF;
        //都是从0开始，因为第一遍循环会把起始点找出来
        for (int i = 0; i < N; i++){
            if (v[i] == 0 && dist[i] < dmin){
                dmin = dist[i];
                u = i;
            }
        }
        //dmin==INF表示所有点已经都遍历过了
        if (dmin == INF)
            break;
        //将找到的点标记一下，说明这个点已经访问过，之后就不要访问了
        v[u] = 1;
        for (int i = 0; i < N; i++){
            if (v[i] == 0){
                //经过当前点到达距离可以更近
                if (dist[i] > dist[u] + mp[u][i]){
                    //更新距离
                    dist[i] = dist[u] + mp[u][i];
                    //到达新的点的队伍数量也进行更新
                    amount[i] = amount[u] + teams[i];
                    /*到达新的路径点的数量就等于上一个点的数量，因为是>号，所以找不到
                    从其他点过来多条路径*/
                    pathcount[i] = pathcount[u];
                }
                else if (dist[i] == dist[u] + mp[u][i]){
                    //如果距离相等就，到达当前的最短路径数就加1
                    pathcount[i] += pathcount[u];
                    //如果之前的最大队伍数量小于现在的队伍数量，则进行更新
                    if (amount[i] < amount[u] + teams[i]){
                        amount[i] = amount[u] + teams[i];
                    }
                }
            }
        }
    }
}
int main(){
    for (; scanf("%d%d%d%d", &N, &M, &start, &en) != EOF;){
        for (int i = 0; i < N; i++){
            scanf("%d", &teams[i]);
        }
        for (int i = 0; i < N; i++){
            dist[i] = INF;
            for (int j = 0; j < N; j++){
                mp[i][j] = INF;
            }
        }
        for (int i = 0; i < M; i++){
            int c1, c2, L;
            scanf("%d%d%d", &c1, &c2, &L);
            mp[c1][c2] = mp[c2][c1] = L;
        }
        dijkstra(start);
        printf("%d %d\n", pathcount[en], amount[en]);
    }
    return 0;
}
```
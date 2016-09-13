[Forward on Weibo](https://www.patest.cn/contests/pat-a-practise/1076)


**解题思路：**
一道广搜的题目，还算简单，就是根据题目输入记录下每个编号的用户的粉丝，这个跟题目的输入有点不一样的，题目给的是每个编号的用户关注的人的编号。查询的时候，从查询的用户编号开始，广搜，中间记录一下层的深度，超过规定深度L的就不用计算了。 
最后一个点应该是很多重复计算的用户的测试用例，因为一开始我的代码最后一个点一直内存超限，因为在遍历的时候没有判断重复，不断的就放到vector中，后来直接在遍历的过程中进行时候访问过的判断就过了。

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<set>
#include<string.h>
using namespace std;
int n, l;
vector<int>* followed;
int flag[1005];
int getTriger(int s){
    memset(flag, 0, sizeof(flag));
    int res = 0;
    int levelCount = 0;
    vector<int>vec;
    vec.push_back(s);
    while (levelCount < l && (vec.size() > 0)){
        vector<int>temp;
        for (int i = 0; i < vec.size(); i++){
            for (int j = 0; j < followed[vec[i]].size(); j++){
                if (followed[vec[i]][j] != s){

                    if (flag[followed[vec[i]][j]] == 0){
                        res++;
                        flag[followed[vec[i]][j]] = 1;
                        temp.push_back(followed[vec[i]][j]);
                    }
                }
            }
        }
        vec.clear();
        vec = temp;
        levelCount++;
    }
    return res;
}
int main(){
    for (; scanf("%d%d", &n, &l) != EOF;){
        followed = new vector<int>[n + 1];
        for (int i = 0; i <= n; i++){
            followed[i] = vector<int>();
        }
        for (int i = 1; i <= n; i++){
            int size;
            scanf("%d", &size);
            while (size--){
                int temp;
                scanf("%d", &temp);
                followed[temp].push_back(i);
            }
        }
        int k;
        scanf("%d", &k);
        while (k--){
            int temp;
            scanf("%d", &temp);
            printf("%d\n", getTriger(temp));
        }
    }
    return 0;
}
```
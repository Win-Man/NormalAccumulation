[Sharing](https://www.patest.cn/contests/pat-a-practise/1032)


**解题思路：**
可能不止有两个单词的节点，可能存在一些混淆点，所以使用了标记数组，从root1开始遍历，将单词一所有的节点进行标记，然后从root2开始进行遍历，遇到节点是已经被标记过的，那就记下地址，然后跳出循环。最后输出结果。

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<algorithm>
#include<string.h>
using namespace std;
struct node{
    char key;
    int address;
    int next;
};

int flag [100005];
node nodes[100005];
int main(){ 
    for (int root1, root2, n; scanf("%d%d%d", &root1, &root2, &n) != EOF;){
        memset(flag, 0, sizeof(int)* 100005);
        while (n--){
            int address;
            scanf("%d", &address);
            scanf(" %c %d", &nodes[address].key, &nodes[address].next);
            nodes[address].address = address;
        }
        int index = -1;
        while (root1 != -1){
            flag[root1]++;
            root1 = nodes[root1].next;
        }
        while (root2 != -1){
            if (flag[root2]){
                index = nodes[root2].address;
                break;
            }
            root2 = nodes[root2].next;
        }
        if (index != -1){
            printf("%05d\n", index);
        }
        else{
            printf("-1\n");
        }
    }
    return 0;
}
```
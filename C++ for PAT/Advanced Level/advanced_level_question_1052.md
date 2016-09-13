[Linked List Sorting](https://www.patest.cn/contests/pat-a-practise/1052)


**解题思路：**
注意两个情况：不知一条链表，还有链表为空的情况，所以next属性是有用的。

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;
struct node{
    int key;
    int address;
    int next;
};
bool cmp(const node n1, const node n2){
    return n1.key < n2.key;
}
node ns[100005];
int main(){ 
    for (int n, root; scanf("%d%d", &n, &root) != EOF;){
        vector<node>nodes;
        for (int i = 0; i < n; i++){
            int address;
            scanf("%d", &address);
            scanf("%d %d", &ns[address].key, &ns[address].next);
            ns[address].address = address;
        }
        while (root != -1){
            nodes.push_back(ns[root]);
            root = ns[root].next;
        }
        sort(nodes.begin(), nodes.end(),cmp);
        if (nodes.size()){
            printf("%d %05d\n", nodes.size(), nodes[0].address);
            for (int i = 0; i < nodes.size() - 1; i++){
                printf("%05d %d %05d\n", nodes[i].address, nodes[i].key, nodes[i + 1].address);
            }
            printf("%05d %d -1\n", nodes[nodes.size() - 1].address, nodes[nodes.size() - 1].key);
        }
        else{
            printf("%d -1\n", nodes.size());
        }
    }
    return 0;
}
```
[Deduplication on a Linked List](https://www.patest.cn/contests/pat-a-practise/1097)


**解题思路：**
比较简单，只要注意可能removed列表可能是空的

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<string.h>
#include<math.h>
using namespace std;
struct node{
    int key;
    int address;
    int next;
};
int num[10005];
node addr[100005];
int main(){
    for (int first, n; scanf("%d %d", &first, &n) != EOF;){
        memset(num, 0, sizeof(int)* 10005);
        for (int i = 0; i < n; i++){
            int temp;
            scanf("%d", &temp);
            addr[temp].address = temp;
            scanf("%d %d", &addr[temp].key, &addr[temp].next);
        }
        int index = first;
        vector<node>output;
        vector<node>remove;
        while (index != -1){
            int temp = abs(addr[index].key);
            if (num[temp]){
                remove.push_back(addr[index]);
            }
            else{
                num[temp] = 1;
                output.push_back(addr[index]);
            }
            index = addr[index].next;
        }
        if (output.size()){
            for (int i = 0; i < output.size() - 1; i++){
                printf("%05d %d %05d\n", output[i].address, output[i].key, output[i + 1].address);
            }
            printf("%05d %d -1\n", output[output.size() - 1].address, output[output.size() - 1].key);
        }
        if (remove.size()){
            for (int i = 0; i < remove.size() - 1; i++){
                printf("%05d %d %05d\n", remove[i].address, remove[i].key, remove[i + 1].address);
            }
            printf("%05d %d -1\n", remove[remove.size() - 1].address, remove[remove.size() - 1].key);
        }
    }
    return 0;
}
```
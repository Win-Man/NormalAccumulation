[Shuffling Machine](https://www.patest.cn/contests/pat-a-practise/1042)


**解题思路：**

**AC code:**

``` c++
#include<iostream>  
#include<string>  
using namespace std;  
int main(){  
    string cards[55] = { "0", "S1", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S10", "S11", "S12", "S13"  
        , "H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8", "H9", "H10", "H11", "H12", "H13"  
        , "C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "C10", "C11", "C12", "C13"  
        , "D1", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "D10", "D11", "D12", "D13"  
        , "J1", "J2" };  
    for (int n; cin >> n;){  
        string res[55] = { "" };  
        int position[55] = { 0 };  
        for (int i = 1; i < 55; i++){  
            cin >> position[i];  
        }  
        //找到经过N次之后牌放的位置  
        for (int i = 1; i < 55; i++){  
            int index = i;  
            for (int j = 0; j < n; j++){  
                index = position[index];  
            }  
            res[index] = cards[i];  
        }  
        cout << res[1];  
        for (int i = 2; i < 55; i++){  
            cout << " " << res[i];  
        }  
        cout << endl;  
    }  
    return 0;  
}  
```
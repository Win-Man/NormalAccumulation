[Have Fun with Numbers](https://www.patest.cn/contests/pat-a-practise/1023)

**解题思路：**
一开始记录各个数字出现的次数，然后double一下，看原先的数字是否够新的数字的组成。

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<algorithm>
#include<stdio.h>
using namespace std;
int main(){
    for (string ss; cin >> ss;){
        int flag[10] = { 0 };
        for (int i = 0; i < ss.length(); i++){
            flag[ss[i] - '0'] ++;
        }
        reverse(ss.begin(), ss.end());
        string double_s = "";
        int carry = 0;
        for (int i = 0; i < ss.length(); i++){  
            double_s = double_s + (char)(((ss[i] - '0') * 2 + carry) % 10 + '0');
            carry = ((ss[i] - '0') * 2 + carry) / 10;
        }
        if (carry){
            double_s += (char)(carry + '0');
        }
        reverse(double_s.begin(), double_s.end());
        int test = 1;
        for (int i = 0; i < double_s.length(); i++){
            if (flag[double_s[i] - '0']){
                flag[double_s[i] - '0']--;
            }
            else{
                test = 0;
                break;
            }
        }
        if (test){
            printf("Yes\n");
            cout << double_s << endl;
        }
        else{
            printf("No\n");
            cout << double_s << endl;
        }
    }
    return 0;
}
```
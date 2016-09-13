[Palindromic Number](https://www.patest.cn/contests/pat-a-practise/1024)


**解题思路：**
不能是使用long类型 使用long类型也会超出范围

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<algorithm>
using namespace std;
bool isPalindromic(string num);
string addReverse(string num);
int main(){
    string num;
    for (int k; cin >> num >> k;){
        int step = 0;
        if (!isPalindromic(num)){
            for (int i = 0; i < k; i++){
                step++;
                num = addReverse(num);
                if (isPalindromic(num)){
                    break;
                }
            }
        }
        cout << num << endl << step << endl;
    }
    return 0;
}
bool isPalindromic(string num){
    bool res = true;
    for (int i = 0; i < num.length() / 2; i++){
        if (num[i] != num[num.length() - 1 - i]){
            res = false;
            break;
        }
    }
    return res;
}
string addReverse(string num){
    string res = num;
    reverse(num.begin(), num.end());
    int carry = 0;
    for (int i = 0; i < num.length(); i++){
        int temp = res[res.length() - 1 - i] - '0' + num[num.length() - 1 - i] - '0' + carry;
        //cout << "temp:" << temp << endl;
        carry = temp / 10;
        res[res.length() - 1 - i] = (temp % 10) + '0';
    }
    if (carry){
        res = "1" + res;
    }
    return res;
}
```
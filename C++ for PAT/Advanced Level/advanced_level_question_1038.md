[Recover the Smallest Number](https://www.patest.cn/contests/pat-a-practise/1038)


**解题思路：**
就是尽量将小的数字开始的数放在前面，让小的数字占据高位，最后形成的数也就越小。
注意点：一开始写程序的时候，我在输入的时候就将以0为开头的数字直接将0去掉之后存放在向量中，因为我认为是每个数字实际的值在进行组合，但是后来提交一直错误，尝试改了一下，发现是这个问题，这个题对于 2 01 02 的输入正确输出应该是102而不是我之前想的12.

**AC code:**

``` c++
#include<iostream>  
#include<string>  
#include<vector>  
#include<algorithm>  
using namespace std;  
bool cmp(const string ss1, const string ss2){  
    int index = 0;  
    string s1 = ss1,s2 = ss2;  
    if (s1.length() < s2.length()){  
        while (s1.length() < s2.length()){  
            s1 = s1 + ss1;  
        }  
    }  
    else if (s1.length() > s2.length()){  
        while (s2.length() < s1.length()){  
            s2 = s2 + ss2;  
        }  
    }  
    while (index < s1.length() && index < s2.length()){  
        if (s1[index] < s2[index]){  
            return true;  
        }  
        else if (s1[index] > s2[index]){  
            return false;  
        }  
        else{  
            index++;  
        }  
    }  
    return true;  
}  
int main(){  
    for (int n; cin >> n;){  
        vector<string>vec;  
        while (n--){  
            string s;  
            cin >> s;  
            vec.push_back(s);  
        }  
        sort(vec.begin(), vec.end(), cmp);  
        string res = "";  
        res += vec[0];  
        for (int i = 1; i < vec.size(); i++){  
            res += vec[i];  
        }  
        if (res.find_first_not_of('0') != -1){  
            cout << res.substr(res.find_first_not_of('0'))<<endl;  
        }  
        else{  
            cout << "0" << endl;  
        }  
    }  
    return 0;  
}  
```
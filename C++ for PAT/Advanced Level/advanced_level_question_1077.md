[Kuchiguse](https://www.patest.cn/contests/pat-a-practise/1077)


**解题思路：**
英语差是硬伤啊，做成了求最长公共字串，结果发现只是求最长的末尾后缀。

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<algorithm>
#include<string>
using namespace std;
int main(){
    for (int n; cin >> n;){
        vector<string>vec(n);
        //读取换行
        getline(cin, vec[0]);
        for (int i = 0; i < n; i++){
            getline(cin, vec[i]);
            reverse(vec[i].begin(), vec[i].end());
        }
        sort(vec.begin(), vec.end());
        string res = "";
        int flag = 1;
        for (int i = 0; i < vec[0].length(); i++){
            for (int j = 1; j < vec.size(); j++){
                if (vec[0][i] != vec[j][i]){
                    flag = 0;
                    break;
                }
            }
            if (flag){
                res += vec[0][i];
            }
            else{
                break;
            }
        }
        if (res.length()){
            reverse(res.begin(), res.end());
            cout << res << endl;
        }
        else{
            cout << "nai" << endl;
        }
    }
    return 0;
}
```
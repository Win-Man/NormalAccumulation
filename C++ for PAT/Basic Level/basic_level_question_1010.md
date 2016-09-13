[一元多项式求导 ](https://www.patest.cn/contests/pat-b-practise/1010)

**AC code:**

``` c++
#include<iostream>
#include<string>
#include<sstream>
#include<vector>
using namespace std;
int main(){
  for(string str;getline(cin,str);){
    stringstream ss;
    vector<int>temp;
    ss.clear();
    ss.str(str);
    while(1){
      int a,b;
      ss>>a>>b;
      if(ss.fail()){
        break;
      }
      if(b == 0){
      }else{
        temp.push_back(a*b);
        temp.push_back(b-1);
      }
    }
    if(temp.size())
    {
      cout<<temp[0];
      for(int i =1;i<temp.size();i++){
        cout<<" "<<temp[i];
      }
    }else
    {
      cout<<"0 0";
    }
    cout<<endl;
  }
  return 0;
}
```
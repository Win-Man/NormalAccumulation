[A+B Format ](https://www.patest.cn/contests/pat-a-practise/1001)

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<stdio.h>
using namespace std;
int main(){
  for(int a,b;cin>>a>>b;){
    vector<int>vec;
    int sum = a+b;
    if(sum >= 0){
      if(sum < 1000){
        cout<<sum<<endl;
      }else{
        while(sum){
          vec.push_back(sum%1000);
          sum /= 1000;
        }
        cout<<vec[vec.size()-1];
        for(int i =vec.size()-2;i > -1;i--){
			printf(",%03d",vec[i]);
        }
        cout<<endl;
      }
    }else{
      sum = -sum;
      if(sum < 1000){
        cout<<-sum<<endl;
      }else{
        while(sum){
          vec.push_back(sum%1000);
          sum /= 1000;
        }
        cout<<"-";
        cout<<vec[vec.size()-1];
        for(int i =vec.size()-2;i > -1;i--){
            printf(",%03d",vec[i]);
        }
        cout<<endl;
      }
    }
  }
}
```
[害死人不偿命的(3n+1)猜想 ](https://www.patest.cn/contests/pat-b-practise/1001)

**AC code:**

``` c++
#include<iostream>
using namespace std;
int main(){
  int n;
  while(cin>>n){
    int count = 0;
    while(n!=1){
      if(n&1){
        n = (n*3+1)>>1;
      }else{
        n = n>>1;
      }
      count++;
    }
    cout<<count<<endl;
  }
  return 0;
}
```
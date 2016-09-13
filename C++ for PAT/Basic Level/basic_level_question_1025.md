[反转链表 ](https://www.patest.cn/contests/pat-b-practise/1025)

**AC code:**

``` c++
//一开始自己的一段代码，其它点都过了，就一个点超时，改了输入输出方式还是不行。
//于是上网搜索了一下别人的思路，调了很久终于AC了，不容易
#include<iostream>
#include<algorithm>
#include<vector>
#include<string>
#include<stdio.h>
#include<math.h>
using namespace std;
struct point{
	int addr;
	int num;
	int next_addr;
};
int main(){
	/*！= EOF一定要加不然会错误，以c++的习惯不是很习惯c的输入输出*/
	for(int first,n,k;scanf("%d%d%d",&first,&n,&k) != EOF;){
		//如果是用数组的话，编译时就会出现栈溢出
		vector<point>arr(100005);
		vector<point>list;
		vector<point>order_list;
		/*下一个节点的地址*/
		int next = first;
		for(int i = 0;i < n;i++){
			point temp;
			scanf("%d%d%d",&temp.addr,&temp.num,&temp.next_addr);
			arr[temp.addr] = temp;
		}
		while(next != -1){
			list.push_back(arr[next]);
			next = arr[next].next_addr;
		}
		/*while(list.size() >= k){
			for(int i = 0;i < k;i++){
				order_list.push_back(list[k-i-1]);
			}
			list.erase(list.begin(),list.begin()+k);
		}
		for(int i = 0;i<list.size();i++){
			order_list.push_back(list[i]);
		}*/
		for(int i = 0;i+k <= list.size();i = i + k){
			for(int j = i+k-1;j >= i;j--){
				order_list.push_back(list[j]);
			}
		}
		for(int i =list.size()-list.size()%k;i<list.size();i++){
			order_list.push_back(list[i]);
		}
		for(int i = 0;i < order_list.size()-1;i++){
			printf("%05d %d %05d\n",order_list[i].addr,order_list[i].num,order_list[i+1].addr);
		}
		printf("%05d %d -1\n",order_list[order_list.size()-1].addr,order_list[order_list.size()-1].num);
	}
	return 0;
}

/*之前一个点超时的代码*/
#include<iostream>
#include<algorithm>
#include<vector>
#include<string>
#include<math.h>
using namespace std;
struct point{
  string addr;
  int num;
  string next_addr;
};
int main(){
  string first;
  for(int n,k;cin>>first>>n>>k;){
    vector<point>list;
    vector<point>order_list;
    string next = first;
    for(int i =0;i < n;i++){
      point temp;
      cin>>temp.addr>>temp.num>>temp.next_addr;
      list.push_back(temp);
    }
    while(next != "-1"){
      for(int i = 0;i < list.size();i++){
        if(next == list[i].addr){
          next = list[i].next_addr;
          order_list.push_back(list[i]);
          list.erase(list.begin()+i,list.begin()+i+1);
          break;
        }
      }
    }
    list.clear();
    while(order_list.size() >= k){
      for(int i = 0; i < k;i++){
        list.push_back(order_list[k-i-1]);
      }
      order_list.erase(order_list.begin(),order_list.begin()+k);
    }
    for(int i = 0;i < order_list.size();i++){
      list.push_back(order_list[i]);
    }
    for(int i =0;i<list.size()-1;i++){
      cout<<list[i].addr<<" "<<list[i].num<<" "
        <<list[i+1].addr<<endl;
    }
    cout<<list[list.size()-1].addr<<" "<<list[list.size()-1].num<<" -1"<<endl;
  }
  return 0;
}
```
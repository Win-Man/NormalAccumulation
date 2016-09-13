[Ëµ·´»° ](https://www.patest.cn/contests/pat-b-practise/1009)

``` c++
#include<iostream>
#include<string>
#include<sstream>
#include<vector>
using namespace std;
int main(){
	for(string str;getline(cin,str);){
		stringstream ss;
		vector<string>temp;
		ss.clear();
		ss.str(str);
		while(1){
			string s;
			ss>>s;
			if(ss.fail()){
				break;
			}
			temp.push_back(s);
		}
		if(temp.size() == 1){
			cout<<temp[0]<<endl;
		}else{
			cout<<temp[temp.size()-1];
			for(int i = 1;i<temp.size();i++){
				cout<<" "<<temp[temp.size()-i-1];
			}
			cout<<endl;
		}
	}
	return 0;
}
```
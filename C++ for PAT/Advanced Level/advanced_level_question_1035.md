[Password ](https://www.patest.cn/contests/pat-a-practise/1035)

``` c++
#include<iostream>
#include<vector>
#include<string>
using namespace std;
struct user{
	string name,pwd;
};
int main(){
	for(int n;cin>>n;){
		vector<user>vec(n);
		vector<user>res;
		for(int i = 0;i < n;i++){
			cin>>vec[i].name>>vec[i].pwd;
		}
		for(int i = 0;i < n;i++){
			int flag = 0;
			for(int j =0; j < vec[i].pwd.length();j++){
				switch ((vec[i].pwd)[j])
				{
				case '1':
					(vec[i].pwd)[j] = '@';
					flag = 1;
					break;
				case '0':
					(vec[i].pwd)[j] = '%';
					flag = 1;
					break;
				case 'l':
					(vec[i].pwd)[j] = 'L';
					flag = 1;
					break;
				case 'O':
					(vec[i].pwd)[j] = 'o';
					flag = 1;
					break;
				default:
					break;
				}
			}
			if(flag){
				res.push_back(vec[i]);
			}
		}
		if(res.size()){
			cout<<res.size()<<endl;
			for(int i = 0; i < res.size();i++){
				cout<<res[i].name<<" "<<res[i].pwd<<endl;
			}
		}else{
			if(n == 1){
				cout<<"There is 1 account and no account is modified"<<endl;
			}else{
				cout<<"There are "<<n<<" accounts and no account is modified"<<endl;
			}
		}
	}
	return 0;
}
```
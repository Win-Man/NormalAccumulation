[The Best Rank ](https://www.patest.cn/contests/pat-a-practise/1012)

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<algorithm>
#include<string>
using namespace std;
struct stu{
	string ID;
	int grade;
};
bool cmp(const stu& stu1,const stu& stu2){
	if(stu1.grade > stu2.grade){
		return true;
	}else{
		return false;
	}
}
int main(){
	for(int n,m;cin>>n>>m;){
		vector<stu>vec_C,vec_M,vec_E,vec_A;
		for(int i = 0;i < n; i++){
			stu temp;
			cin>>temp.ID>>temp.grade;
			vec_C.push_back(temp);
			cin>>temp.grade;
			vec_M.push_back(temp);
			cin>>temp.grade;
			vec_E.push_back(temp);
			temp.grade = (int)((vec_C[i].grade + vec_M[i].grade + vec_E[i].grade)*1.0/3.0+0.5);
			vec_A.push_back(temp);
		}
		sort(vec_C.begin(),vec_C.end(),cmp);
		sort(vec_M.begin(),vec_M.end(),cmp);
		sort(vec_E.begin(),vec_E.end(),cmp);
		sort(vec_A.begin(),vec_A.end(),cmp);
		for(int i = 0;i < m;i++){
			string id;
			cin>>id;
			int flag = 0;
			int res[4]={0};
			for(int i = 0;i < n;i++){
				if(vec_A[i].ID == id){
					int o = i;
					if(i > 0){
						while(o){
							if(vec_A[o].grade == vec_A[o-1].grade){
								o--;
							}else{
								break;
							}
						}
					}
					res[0]=o+1;
					flag = 1;
				}
				if(vec_C[i].ID == id){
					int o = i;
					if(i > 0){
						while(o >0){
							if(vec_C[o].grade == vec_C[o-1].grade){
								o--;
							}else{
								break;
							}
						}
					}
					res[1]=o+1;
					flag = 1;
				}
				if(vec_M[i].ID == id){
					int o = i;
					if(i > 0){
						while(o > 0){
							if(vec_M[o].grade == vec_M[o-1].grade){
								o--;
							}else{
								break;
							}
						}
					}
					res[2]=o+1;
					flag = 1;
				}
				if(vec_E[i].ID == id){
					int o = i;
					if(i > 0){
						while(o > 0){
							if(vec_E[o].grade == vec_E[o-1].grade){
								o--;
							}else{
								break;
							}
						}
					}
					res[3] = o+1;
					flag = 1;
				}
			}
			if(!flag){
				cout<<"N/A"<<endl;
				continue;
			}
			if(res[0] <= res[1] && res[0] <= res[2] && res[0] <= res[3]){
				cout<<res[0]<<" A"<<endl;
			}else if(res[1] < res[0] && res[1] <= res[2] && res[1] <= res[3]){
				cout<<res[1]<<" C"<<endl;
			}else if(res[2] < res[0] && res[2] < res[1] && res[2] <= res[3]){
				cout<<res[2]<<" M"<<endl;
			}else if(res[3] < res[0] && res[3] < res[1] && res[3] < res[2]){
				cout<<res[3]<<" E"<<endl;
			}
		}
	}
	return 0;
}
```
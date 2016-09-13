[Êý×Ö·ÖÀà ](https://www.patest.cn/contests/pat-b-practise/1012)

``` c++
#include<iostream>
#include<vector>
#include<iomanip>
#include<algorithm>
using namespace std;
int main(){
	for(int n;cin>>n;){
		int count_1=0,count_2=0,count_3=0,count_4=0,count_5=0;
		vector<int>vec1;
		vector<int>vec2;
		vector<int>vec3;
		vector<int>vec4;
		vector<int>vec5;
		for(int i = 0;i < n;i++){
			int temp;
			cin>>temp;
			switch (temp % 5)
			{
			case 0:
				if(!(temp&1)){
					vec1.push_back(temp);
				}
				break;
			case 1:
				vec2.push_back(temp);
				break;
			case 2:
				vec3.push_back(temp);
				break;
			case 3:
				vec4.push_back(temp);
				break;
			case 4:
				vec5.push_back(temp);
				break;
			default:
				break;
			}
		}
		if(vec1.size()){
			int sum = 0;
			for(int i =0;i<vec1.size();i++){
				if( ! (vec1[i] & 1)){
					sum+=vec1[i];
				}
			}
			cout<<sum;
		}else{
			cout<<"N";
		}
		if(vec2.size()){
			int sum = 0;
			for(int i = 0;i<vec2.size();i++){
				if(i&1){
					sum-=vec2[i];
				}else{
					sum+=vec2[i];
				}
			}
			cout<<" "<<sum;
		}else
		{
			cout<<" N";
		}
		if(vec3.size()){
			cout<<" "<<vec3.size();
		}else{
			cout<<" N";
		}
		if(vec4.size()){
			double sum = 0;
			for(int i = 0;i<vec4.size();i++){
				sum+=vec4[i];
			}
			sum = sum/vec4.size()*1.0;
			cout<<" "<<fixed<<setprecision(1)<<sum;
		}else{
			cout<<" N";
		}
		if(vec5.size()){
			sort(vec5.begin(),vec5.end());
			cout<<" "<<vec5[vec5.size()-1]<<endl;
		}else
		{
			cout<<" N"<<endl;
		}
	}
	return 0;
}
```
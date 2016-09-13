[´¸×Ó¼ôµ¶²¼ ](https://www.patest.cn/contests/pat-b-practise/1018)

``` c++
#include<iostream>
#include<string>
#include<math.h>
#include<stdio.h>
using namespace std;
int main(){
	for(int n;cin>>n;){
		int win_A=0,win_B=0,pin=0,lose_A=0,lose_B=0;
		//012 BCJ
		int log_A[3]={0,0,0};
		int log_B[3]={0,0,0};
		for(int i = 0;i < n;i++){
			char a,b;
			cin>>a>>b;
			if(a == b){
				pin++;
			}else if(a == 'C' && b == 'J'){
				win_A++;
				lose_B++;
				log_A[1]++;
			}else if(a == 'C' && b == 'B'){
				win_B++;
				lose_A++;
				log_B[0]++;
			}else if(a == 'J' && b == 'C'){
				win_B++;
				lose_A++;
				log_B[1]++;
			}else if(a == 'J' && b =='B'){
				win_A++;
				lose_B++;
				log_A[2]++;
			}else if(a == 'B' && b == 'J'){
				win_B++;
				lose_A++;
				log_B[2]++;
			}else if(a == 'B' && b == 'C'){
				win_A++;
				lose_B++;
				log_A[0]++;
			}
		}
		printf("%d %d %d\n",win_A,pin,lose_A);
		printf("%d %d %d\n",win_B,pin,lose_B);
		if(log_A[0] > log_A[1] && log_A[0] > log_A[2]){
			printf("%c ",'B');
		}else if(log_A[1] > log_A[0] && log_A[1] > log_A[2]){
			printf("%c ",'C');
		}else if(log_A[2] > log_A[0] && log_A[2] > log_A[1]){
			printf("%c ",'J');
		}else if(log_A[0] == log_A[1] && log_A[0] > log_A[2]){
			printf("%c ",'B');
		}else if(log_A[0] == log_A[2] && log_A[0] > log_A[1]){
			printf("%c ",'B');
		}else if(log_A[1] == log_A[2] && log_A[1] > log_A[0]){
			printf("%c ",'C');
		}else if(log_A[0] == log_A[1] && log_A[1] == log_A[2]){
			printf("%c ",'B');
		}
		if(log_B[0] > log_B[1] && log_B[0] > log_B[2]){
			printf("%c\n",'B');
		}else if(log_B[1] > log_B[0] && log_B[1] > log_B[2]){
			printf("%c\n",'C');
		}else if(log_B[2] > log_B[0] && log_B[2] > log_B[1]){
			printf("%c\n",'J');
		}else if(log_B[0] == log_B[1] && log_B[0] > log_B[2]){
			printf("%c\n",'B');
		}else if(log_B[0] == log_B[2] && log_B[0] > log_B[1]){
			printf("%c\n",'B');
		}else if(log_B[1] == log_B[2] && log_B[1] > log_B[0]){
			printf("%c\n",'C');
		}else if(log_B[0] == log_B[1] && log_B[1] == log_B[2]){
			printf("%c\n",'B');
		}
	}
	return 0 ;
}
```
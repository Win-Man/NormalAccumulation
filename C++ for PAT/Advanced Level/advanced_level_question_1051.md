[Pop Sequence ](https://www.patest.cn/contests/pat-a-practise/1051)

**AC code:**

``` c++
#include<iostream>
#include<stack>
#include<stdio.h>
using namespace std;
int main(){
	for(int m,n,k;scanf("%d%d%d",&m,&n,&k)!= EOF;){
		while(k--){
			//����һ����ջ add��ģ��ջѹ���ʱ����� flag��������Ƿ����
			stack<int>sta;
			int flag = 1;
			int add = 1;
			for(int i = 0; i < n;i++){
				int num;
				scanf("%d",&num);
				if(flag){
					//��ջΪ��ջ����˵��ջ��Ԫ�ػ������ڶ�ȡ�����ֵ�ʱ��һֱ��ջ����push�������м��ж�ջ�����Ƿ񳬹�
					while(sta.size() == 0 || sta.top() != num){
						sta.push(add);
						if(sta.size() > m){
							flag = 0;
							break;
						}
						add++;
					}
					if(flag && sta.size() >= 1 && sta.top() == num){
						sta.pop();
					}
				}
			}
			if(flag){
				printf("YES\n");
			}else{
				printf("NO\n");
			}
		}
	}
	return 0;
}
```
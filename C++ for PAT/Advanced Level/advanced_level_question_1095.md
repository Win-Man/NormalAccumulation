[Cars on Campus](https://www.patest.cn/contests/pat-a-practise/1095)


**解题思路：**
首先对所有的记录进行筛选，保留能配对的记录，配对原则是选择时间距离最短的进行in-out匹配。对于查询，我是维护了一个数组，用于记录以秒为单位的每个事件点的状态，比如时间点x上有车进来，则对这一时刻进行++操作，有车出去则进行–操作。当查询的时候，从0开始直至查询时刻进行求和，就是当时在学校内停靠的车的数量。 
查询的时候不必每次都从0开始进行求和操作，因为题目说查询的时间是按照升序进行输入的，所以保留每次求和的结果，可以被下一次查询使用，这样可以减少时间消耗，不然测试点4会运行超时。 
保留求和操作的时候要注意一个点就是，第一次求和的时候是从0开始进行求和的

**AC code:**

``` c++
#include<iostream>
#include<vector>
#include<algorithm>
#include<map>
#include<string.h>
#include<string>
using namespace std;
struct record{
    //char plate[8];
    string plate;
    int hour;
    int minute;
    int second;
    string status;
    int totalTime = 0;
    //char status[4];
};
bool cmp(const record& r1, const record&r2){
    return r1.hour * 3600 + r1.minute * 60 + r1.second < r2.hour * 3600 + r2.minute * 60 + r2.second;
}
int carCount[100000];
int main(){
    for (int n, k; scanf("%d%d", &n, &k) != EOF;){
        memset(carCount, 0, sizeof(carCount));
        vector<string>longestCar;
        int maxTime = 0;
        map<string, vector<record>>cars;
        for (int i = 0; i < n; i++){
            record temp;
            cin >> temp.plate;
            scanf("%d:%d:%d", &temp.hour, &temp.minute, &temp.second);
            cin >> temp.status;
            cars[temp.plate].push_back(temp);
        }
        map<string, vector<record>>::iterator iter;
        for (iter = cars.begin(); iter != cars.end(); iter++){
            sort(iter->second.begin(), iter->second.end(), cmp);
            vector<record>temp;
            //printf("%s\n", iter->first);
            int tempTime = 0;
            for (int i = 0; i < iter->second.size(); i++){
                if (iter->second[i].status[0] == 'i'){
                    record start = iter->second[i];
                    for (i = i + 1; i < iter->second.size(); i++){
                        if (iter->second[i].status[0] == 'i'){
                            start = iter->second[i];
                        }
                        else{
                            temp.push_back(start);
                            carCount[start.hour * 3600 + start.minute * 60 + start.second]++;
                            record end = iter->second[i];
                            carCount[end.hour * 3600 + end.minute * 60 + end.second]--;
                            temp.push_back(end);
                            tempTime += (end.hour * 3600 + end.minute * 60 + end.second - start.hour * 3600 - start.minute * 60 - start.second);
                            break;
                        }
                    }
                }
            }
            if (tempTime > maxTime){
                longestCar.clear();
                longestCar.push_back(iter->first);
                maxTime = tempTime;
            }
            else if (tempTime == maxTime){
                longestCar.push_back(iter->first);
            }
            iter->second = temp;
        }
        int lTime = 0;
        int sum = 0;
        for (int i = 0; i < k; i++){
            int hour;
            int minute;
            int second;
            scanf("%d:%d:%d", &hour, &minute, &second);
            int qTime = hour * 3600 + minute * 60 + second;
            int j;
            if (i == 0){
                j = 0;
            }
            else{
                j = lTime + 1;
            }
            for (; j <= qTime; j++){
                sum += carCount[j];
            }
            lTime = qTime;
            printf("%d\n", sum);
        }
        sort(longestCar.begin(),longestCar.end());
        for (int i = 0; i < longestCar.size(); i++){
            cout << longestCar[i] << " ";
        }
        printf("%02d:%02d:%02d\n", maxTime / 3600, (maxTime % 3600) / 60, maxTime % 60);
    }
    return 0;
}
```
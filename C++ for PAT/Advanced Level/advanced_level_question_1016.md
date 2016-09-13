[Phone Bills](https://www.patest.cn/contests/pat-a-practise/1016)


**解题思路：**
将所有的记录进行排序，先按名字的字母序排，相同之后时间从小打到排，这样同一个人的所有的记录都是在一起的方便处理。然后将同一个人的所有记录提取出来，进行配对计算输出，配对原则是按照距离最近的on-off配对。 
题目坑爹的一点是说每个输入一定至少有一对配对的，但是并不是每个人的记录都一zing至少有一对配对的，对于某些一对配对都没有的顾客不用输出。 
最后还是吐槽一下傻逼的自己，cmp函数居然都能写错 
return r1.minute < r2.minute;写成return r1.minute < r1.minute; 最后发现错误的我差点哭了。

**AC code:**

``` c++
#include<iostream>
#include<stdio.h>
#include<vector>
#include<algorithm>
#include<string.h>
using namespace std;
struct record{
    char name[22];
    int month;
    int day;
    int hour;
    int minute;
    char status[10];
};
int tolls[24];
bool cmp(const record & r1, const record & r2){
    if (strcmp(r1.name, r2.name) < 0){
        return true;
    }
    else if (strcmp(r1.name, r2.name) == 0){
        if (r1.day < r2.day){
            return true;
        }
        else if (r1.day == r2.day){
            if (r1.hour < r2.hour){
                return true;
            }
            else if (r1.hour == r2.hour){
                return r1.minute < r2.minute;
            }
            else{
                return false;
            }
        }
        else{
            return false;
        }
    }
    else{
        return false;
    }
}
double show_record(record start, record end){
    record s = start;
    double sum = 0;
    int count = 0;
    /*while (start.day < end.day || start.hour < end.hour || start.minute < end.minute){

        start.minute++;
        count++;
        sum += tolls[start.hour] * 1.0 / 100.0;
        if (start.minute == 60){
            start.minute = 0;
            start.hour++;
        }
        if (start.hour == 24){
            start.hour = 0;
            start.day++;
        }

    }*/
    int count_start = start.minute + start.hour * 60 + start.day * 24 * 60;
    int count_end = end.minute + end.hour * 60 + end.day*24*60;
    int cost_start = start.minute * tolls[start.hour];
    int cost_end = end.minute * tolls[end.hour];
    for (int i = 0; i < start.hour; i++){
        cost_start += tolls[i] * 60;
    }
    for (int i = 0; i < end.hour; i++){
        cost_end += tolls[i] * 60;
    }
    int temp = 0;
    for (int i = 0; i < 24; i++){
        temp += tolls[i];
    }
    cost_start += start.day*temp * 60;
    cost_end += end.day*temp * 60;
    sum = (cost_end - cost_start)*1.0/100.0;
    printf("%02d:%02d:%02d %02d:%02d:%02d %d $%.2lf\n", s.day, s.hour, s.minute, end.day, end.hour, end.minute, count_end-count_start, (cost_end-cost_start)*1.0/100.0);
    return sum;
}
void show(vector<record>vec){
    record start;
    record end;
    double total = 0.0;

    vector<record>haha;
    for (int i = 0; i < vec.size(); i++){
        if (strcmp(vec[i].status, "on-line") == 0){
            start = vec[i];
            for (i = i + 1; i < vec.size(); i++){
                if (strcmp(vec[i].status, "on-line") == 0){
                    start = vec[i];
                }
                if (strcmp(vec[i].status, "off-line") == 0){
                    end = vec[i];
                    haha.push_back(start);
                    haha.push_back(end);
                    break;
                }
            }
        }
    }
    if (haha.size()){
        printf("%s %02d\n", haha[0].name, haha[0].month);
        for (int i = 0; i < haha.size(); i += 2){
            total += show_record(haha[i], haha[i + 1]);
        }
        printf("Total amount: $%.2lf\n", total);
    }

}
int main(){

    for (; scanf("%d", &tolls[0]) != EOF;){
        for (int i = 1; i < 24; i++){
            scanf("%d", &tolls[i]);
        }
        int n;
        scanf("%d", &n);
        vector<record>vec(n);
        for (int i = 0; i < n; i++){
            scanf("%s %d:%d:%d:%d %s", vec[i].name, &vec[i].month, &vec[i].day, &vec[i].hour, &vec[i].minute, vec[i].status);
        }
        sort(vec.begin(), vec.end(), cmp);
        record start;
        record end;
        char *name = vec[0].name;
        vector<record>out;
        out.push_back(vec[0]);
        for (int i = 1; i < n; i++){
            if (strcmp(name, vec[i].name) == 0){
                out.push_back(vec[i]);
            }
            else{
                name = vec[i].name;
                show(out);
                out.clear();
                out.push_back(vec[i]);
            }
        }
        if (out.size()){
            show(out);
        }
    }
    return 0;
}
```
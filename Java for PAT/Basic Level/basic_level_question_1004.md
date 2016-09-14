[成绩排名](https://www.patest.cn/contests/pat-b-practise/1004)

**解题思路：**

**AC code:**

``` java
import java.util.*;

/**
 * Created by sg on 2016/9/14.
 */
public class Main {

    public static void main(String[] args){
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        List<Student> students = new ArrayList<Student>();
        for(int i =0;i < n;i++){
            String name = sc.next();
            String code = sc.next();
            int grade = sc.nextInt();
            students.add(new Student(name,code,grade));
        }
        Collections.sort(students, new Comparator<Student>() {
            public int compare(Student o1, Student o2) {
                return o1.getGrade() - o2.getGrade();
            }
        });
        System.out.println(String.format("%s %s",students.get(n-1).getName(),students.get(n-1).getCode()));
        System.out.println(String.format("%s %s",students.get(0).getName(),students.get(0).getCode()));
    }
}
class Student{
    private String name;
    private String code;
    private int grade;

    public Student(){

    }

    public Student(String name,String code,int grade){
        this.name = name;
        this.code = code;
        this.grade = grade;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getCode() {
        return code;
    }

    public void setCode(String code) {
        this.code = code;
    }

    public int getGrade() {
        return grade;
    }

    public void setGrade(int grade) {
        this.grade = grade;
    }

    public String toString(){
        return String.format("姓名：%s  学号:%s  成绩：%d",getName(),getCode(),getGrade());
    }

}
```
package Test;

import com.itheima.dao.ICourseMapper;
import com.itheima.dao.ISchoolMapper;
import com.itheima.pojo.ICourse;
import com.itheima.pojo.ISchool;
import com.itheima.utils.MyBatisUtils;
import org.apache.ibatis.session.SqlSession;
import org.junit.Test;

public class MyBatisTest {
    @Test
    public void findIStudentByIdTest() {
        // 1.通过工具类获取SqlSession对象
        SqlSession session = MyBatisUtils.getSession();
        ICourseMapper mapperCourse = session.getMapper(ICourseMapper.class);
        ISchoolMapper mapperSchool = session.getMapper(ISchoolMapper.class);

        // 查询 id=2 的课程信息；
        ICourse course1 = mapperCourse.SelectCourse(2);
        System.out.println(course1.toString());
        // 查询出所有计算机学院开设的课程信息 ；
        ISchool school1 = mapperSchool.SelectSchool(1);
        System.out.println(school1.toString());
        // 将 id=4 这⻔课程的课时数修改为 32+8=40；
        int course2 = mapperCourse.UpdateCourseHours(32 + 8, 4);

        // 插⼊⼀条新的课程记录： names=”⼤数据存储“，hours=32，schools =1；
        int course3 = mapperCourse.InsertCourse("大数据存储",32,1);

        // 3.关闭SqlSession
        session.close();
    }
}


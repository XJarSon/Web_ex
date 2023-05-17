package com.itheima.dao;

import com.itheima.pojo.ICourse;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;
import org.apache.ibatis.annotations.Update;

public interface ICourseMapper {
    @Select("select * from c_course where id = #{id}")
    ICourse SelectCourse(Integer id);
    @Insert("insert into c_course  (name,hours,sid) values (#{name}, #{hours}, #{sid})")
    Integer InsertCourse(@Param("name") String name,@Param("hours") Integer hours,@Param("sid") Integer sid);
    @Update("update c_course set hours = #{hours} where id = #{id}")
    Integer UpdateCourseHours(@Param("hours") Integer hours, @Param("id") Integer id);
}
package com.itheima.dao;

import com.itheima.pojo.ISchool;
import org.apache.ibatis.annotations.Select;

public interface ISchoolMapper {
    @Select("select * from s_school where id = #{id}")
    ISchool SelectSchool(Integer id);

}

package com.itheima.pojo;

public class ISchool {
    private Integer id;
    private String school_name;

    public String getSchool_name() {
        return school_name;
    }

    public Integer getId() {
        return id;
    }

    public void setSchool_name(String school_name) {
        this.school_name = school_name;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    @Override
    public String toString() {
        return "ISchool{" +
                "id=" + id +
                ", school_name='" + school_name + '\'' +
                '}';
    }
}

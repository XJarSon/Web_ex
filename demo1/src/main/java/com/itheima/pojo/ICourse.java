package com.itheima.pojo;

public class ICourse {
    private Integer id;
    private String name;
    private Integer hours;
    private Integer sid;

    public void setName(String name) {
        this.name = name;
    }
    public void setId(Integer id) {
        this.id = id;
    }

    public void setSid(Integer sid) {
        this.sid = sid;
    }

    public void setHours(Integer hours) {
        this.hours = hours;
    }

    public Integer getSid() {
        return sid;
    }

    public Integer getHours() {
        return hours;
    }

    public String getName() {
        return name;
    }

    public Integer getId() {
        return id;
    }

    @Override
    public String toString() {
        return "ICourse{" +
                "id=" + id +
                ", name=" + name +
                ", hours=" + hours +
                ", sid=" + sid +
                '}';
    }
}

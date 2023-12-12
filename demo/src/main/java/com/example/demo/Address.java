package com.example.demo;


@Entity
public class Address {
    private String street;
    private String code;
    private String city;
    private String country;

    public static Address address(String street, String code, String city, String country) {
        code.isB
        return new Address(street, code, city, country);
    }

    private Address(String street, String code, String city, String country) {
        this.street = street;
        this.code = code;
        this.city = city;
        this.country = country;
    }

    public String getStreet() {
        return street;
    }

    public String getCode() {
        return code;
    }

    public String getCity() {
        return city;
    }

    public String getCountry() {
        return country;
    }
}

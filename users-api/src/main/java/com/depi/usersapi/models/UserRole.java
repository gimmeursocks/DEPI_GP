package com.depi.usersapi.models;

public enum UserRole {
    USER("user"),
    ADMIN("admin");

    private String description;
    UserRole(final String description) {
        this.description = description;
    }


    @Override
    public String toString() {
        return this.description;
    }
}

package com.depi.usersapi.repository;

import org.springframework.data.repository.CrudRepository;

import com.depi.usersapi.models.User;

public interface UserRepository extends CrudRepository<User, String> {
    User findOneByUsername(String username);
    User findByUsername(String username);
    User getByUsername(String username);
}

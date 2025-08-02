package org.jodyvole.userservice.repository;

import org.jodyvole.userservice.entity.User;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface UserRepository extends CrudRepository<User, Long> {

  Optional<User> findByUsername(String username);

}

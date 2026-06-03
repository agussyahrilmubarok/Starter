package io.github.agussyahrilmubarok.backend.infrastructure.persistence.specification;

import io.github.agussyahrilmubarok.backend.domain.User;
import jakarta.persistence.criteria.Predicate;
import org.springframework.data.jpa.domain.Specification;

import java.util.ArrayList;
import java.util.List;

public class UserSpecification {

    public static Specification<User> hasSearch(String search) {
        return (root, query, criteriaBuilder) -> {
            if (search == null || search.isBlank()) {
                return criteriaBuilder.conjunction();
            }

            String likeSearch = "%" + search.toLowerCase() + "%";

            List<Predicate> predicates = new ArrayList<>();
            predicates.add(criteriaBuilder.like(criteriaBuilder.lower(root.get("name")), likeSearch));
            predicates.add(criteriaBuilder.like(criteriaBuilder.lower(root.get("email")), likeSearch));

            return criteriaBuilder.or(predicates.toArray(new Predicate[0]));
        };
    }
}
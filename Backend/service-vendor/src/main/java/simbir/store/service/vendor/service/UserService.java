package simbir.store.service.vendor.service;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import simbir.store.service.vendor.db.entity.User;
import simbir.store.service.vendor.db.jpaRepository.UserRepository;


import java.util.Date;
import java.util.List;

@Service
@RequiredArgsConstructor
public class UserService {

    private  UserRepository userRepository;

    /**
     * Создание нового пользователя.
     * Роль должна быть пустой или USER.
     */
    public User createUser(User user) {
        user.setId(null);
        user.setRoles(null); // или "USER"
        user.setCreateAt((java.sql.Date) new Date());

        User saved = userRepository.save(user);


        return saved;
    }

    /**
     * Внешний сервис одобрил заявку на регистрацию продавца.
     */
    public User approveVendorRegistration(Long userId) {
        User user = userRepository.findById(userId)
                .orElseThrow(() -> new RuntimeException("User not found"));

        user.setRoles("VENDOR");
        userRepository.save(user);

        return user;
    }

    /**
     * Получить всех продавцов
     */
    public List<User> getVendors() {
        return userRepository.findAll()
                .stream()
                .filter(u -> "VENDOR".equals(u.getRoles()))
                .toList();
    }
}

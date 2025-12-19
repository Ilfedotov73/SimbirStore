package simbir.store.service.vendor.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;
import simbir.store.service.vendor.db.entity.User;
import simbir.store.service.vendor.service.UserService;

import java.util.List;

@RestController
@RequestMapping("/api/users")
@RequiredArgsConstructor
public class UserController {

    private  UserService userService;

    /**
     * Создание нового пользователя
     */
    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public User createUser(@RequestBody User user) {
        return userService.createUser(user);
    }

    /**
     * Одобрение регистрации продавца
     */
    @PostMapping("/{userId}/approve")
    public User approveVendor(@PathVariable Long userId) {
        return userService.approveVendorRegistration(userId);
    }

    /**
     * Получение всех продавцов
     */
    @GetMapping("/vendors")
    public List<User> getVendors() {
        return userService.getVendors();
    }
}

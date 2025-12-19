package simbir.store.service.vendor.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import simbir.store.service.vendor.db.entity.User;
import simbir.store.service.vendor.service.UserService;

import java.util.List;

import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyLong;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

class UserControllerTest {

    private MockMvc mockMvc;

    private ObjectMapper objectMapper = new ObjectMapper();

    @Mock
    private UserService userService;

    @InjectMocks
    private UserController userController;

    @BeforeEach
    void setup() {
        MockitoAnnotations.openMocks(this);
        mockMvc = MockMvcBuilders.standaloneSetup(userController).build();
    }

    @Test
    void testCreateUser() throws Exception {
        User user = new User();
        user.setId(1);
        user.setLogin("user1");

        when(userService.createUser(any(User.class))).thenReturn(user);

        mockMvc.perform(post("/api/users")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(user)))
                .andExpect(status().isCreated())
                .andExpect(jsonPath("$.id").value(1))
                .andExpect(jsonPath("$.login").value("user1"));
    }

    @Test
    void testApproveVendor() throws Exception {
        User user = new User();
        user.setId(1);
        user.setRoles("VENDOR");

        when(userService.approveVendorRegistration(anyLong())).thenReturn(user);

        mockMvc.perform(post("/api/users/1/approve"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.roles").value("VENDOR"));
    }

    @Test
    void testGetVendors() throws Exception {
        User user = new User();
        user.setId(1);
        user.setRoles("VENDOR");

        when(userService.getVendors()).thenReturn(List.of(user));

        mockMvc.perform(get("/api/users/vendors"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$[0].roles").value("VENDOR"));
    }
}

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
import org.springframework.web.client.RestTemplate;
import simbir.store.service.vendor.dto.NoticesDto;

import java.sql.Date;

import static org.mockito.Mockito.*;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

class NoticeControllerTest {

    private MockMvc mockMvc;
    private ObjectMapper objectMapper = new ObjectMapper();

    @Mock
    private RestTemplate restTemplate;

    @InjectMocks
    private NoticeController noticeController;

    private NoticesDto notice;

    @BeforeEach
    void setup() {
        MockitoAnnotations.openMocks(this);
        mockMvc = MockMvcBuilders.standaloneSetup(noticeController).build();

        notice = new NoticesDto();
        notice.setId(1);
        notice.setEntityId(1);
        notice.setText("Test notice");
        notice.setCreateAt(Date.valueOf("2025-12-19"));
    }

    @Test
    void testReceiveNotice() throws Exception {
        mockMvc.perform(post("/api/notices")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(notice)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.text").value("Test notice"));
    }

}

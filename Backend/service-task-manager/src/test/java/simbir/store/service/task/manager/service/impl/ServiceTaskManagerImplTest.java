package simbir.store.service.task.manager.service.impl;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import simbir.store.service.task.manager.db.entity.AdminTaskProductRegistration;
import simbir.store.service.task.manager.db.entity.AdminTaskVendorRegistration;
import simbir.store.service.task.manager.db.entity.TaskProductRegistration;
import simbir.store.service.task.manager.db.entity.TaskVendorRegistration;
import simbir.store.service.task.manager.db.jpaRepository.AdminsTasksQueueProductsRegistrationRepository;
import simbir.store.service.task.manager.db.jpaRepository.AdminsTasksQueueVendorRegistrationRepository;
import simbir.store.service.task.manager.db.jpaRepository.TasksQueueProductsRegistrationRepository;
import simbir.store.service.task.manager.db.jpaRepository.TasksQueueVendorRegistrationRepository;
import simbir.store.service.task.manager.dto.TaskProductsRegistrationDto;
import simbir.store.service.task.manager.dto.TaskVendorRegistrationDto;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class ServiceTaskManagerImplTest {

    @InjectMocks
    private ServiceTaskManagerImpl service;

    @Mock
    private TasksQueueVendorRegistrationRepository vendorRegRepo;

    @Mock
    private TasksQueueProductsRegistrationRepository productsRegRepo;

    @Mock
    private AdminsTasksQueueVendorRegistrationRepository adminsVendorRepo;

    @Mock
    private AdminsTasksQueueProductsRegistrationRepository adminsProductsRepo;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    // -----------------------------
    // Тест создания задач
    // -----------------------------

    @Test
    void testCreateVendorRegistrationTask() {
        TaskVendorRegistrationDto dto = new TaskVendorRegistrationDto();
        dto.setTaskName("Vendor Task");
        dto.setUserId(1L);
        dto.setDeadline(LocalDateTime.now().toString());
        dto.setStatus("NEW");
        dto.setCreatAt(LocalDateTime.now().toLocalDate().toString());

        TaskVendorRegistrationDto result = service.createVendorRegistrationTask(dto);

        assertEquals(dto, result);
        verify(vendorRegRepo, times(1)).save(any(TaskVendorRegistration.class));
    }

    @Test
    void testCreateProductsRegistrationTask() {
        TaskProductsRegistrationDto dto = new TaskProductsRegistrationDto();
        dto.setTaskName("Product Task");
        dto.setUserId(1L);
        dto.setDeadline(LocalDateTime.now().toLocalDate().toString());
        dto.setStatus("NEW");
        dto.setCreatAt(LocalDateTime.now().toLocalDate().toString());
        dto.setContainer("Container1");

        TaskProductsRegistrationDto result = service.createProductsRegistrationTask(dto);

        assertEquals(dto, result);
        verify(productsRegRepo, times(1)).save(any(TaskProductRegistration.class));
    }

    // -----------------------------
    // Тест получения задач
    // -----------------------------

    @Test
    void testGetVendorRegistrationTasks() {
        when(vendorRegRepo.findAllByOrderByCreatedAtAsc()).thenReturn(List.of(new TaskVendorRegistration()));

        List<TaskVendorRegistration> tasks = service.getVendorRegistrationTasks();

        assertFalse(tasks.isEmpty());
        verify(vendorRegRepo, times(1)).findAllByOrderByCreatedAtAsc();
    }

    @Test
    void testGetProductsRegistrationTasks() {
        when(productsRegRepo.findAllByOrderByCreatedAtAsc()).thenReturn(List.of(new TaskProductRegistration()));

        List<TaskProductRegistration> tasks = service.getProductsRegistrationTasks();

        assertFalse(tasks.isEmpty());
        verify(productsRegRepo, times(1)).findAllByOrderByCreatedAtAsc();
    }

    @Test
    void testGetVendorRegistrationTaskById_Found() {
        TaskVendorRegistration task = new TaskVendorRegistration();
        when(vendorRegRepo.findById(1L)).thenReturn(Optional.of(task));

        TaskVendorRegistration result = service.getVendorRegistrationTaskById(1L);

        assertEquals(task, result);
    }

    @Test
    void testGetVendorRegistrationTaskById_NotFound() {
        when(vendorRegRepo.findById(1L)).thenReturn(Optional.empty());

        Exception exception = assertThrows(RuntimeException.class, () -> service.getVendorRegistrationTaskById(1L));
        assertEquals("Vendor registration task not found", exception.getMessage());
    }

    // -----------------------------
    // Тест назначения задач администратору
    // -----------------------------

    @Test
    void testAssignVendorRegistrationTask() {
        service.assignVendorRegistrationTask(10L, 20L);
        verify(adminsVendorRepo, times(1)).save(any(AdminTaskVendorRegistration.class));
    }

    @Test
    void testAssignProductsRegistrationTask() {
        service.assignProductsRegistrationTask(10L, 20L);
        verify(adminsProductsRepo, times(1)).save(any(AdminTaskProductRegistration.class));
    }

    // -----------------------------
    // Тест обновления статуса
    // -----------------------------

    @Test
    void testUpdateVendorRegistrationTaskStatus() {
        TaskVendorRegistration task = new TaskVendorRegistration();
        when(vendorRegRepo.findById(1L)).thenReturn(Optional.of(task));

        service.updateVendorRegistrationTaskStatus(1L, 10L, "COMPLETED");

        assertEquals("COMPLETED", task.getStatus());
        verify(vendorRegRepo, times(1)).save(task);
        verify(adminsVendorRepo, times(1)).save(any(AdminTaskVendorRegistration.class));
    }

    @Test
    void testUpdateProductsRegistrationTaskStatus() {
        TaskProductRegistration task = new TaskProductRegistration();
        when(productsRegRepo.findById(1L)).thenReturn(Optional.of(task));

        service.updateProductsRegistrationTaskStatus(1L, 10L, "COMPLETED");

        assertEquals("COMPLETED", task.getStatus());
        verify(productsRegRepo, times(1)).save(task);
        verify(adminsProductsRepo, times(1)).save(any(AdminTaskProductRegistration.class));
    }
}

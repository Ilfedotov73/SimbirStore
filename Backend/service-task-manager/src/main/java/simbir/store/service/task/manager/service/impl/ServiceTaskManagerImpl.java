package simbir.store.service.task.manager.service.impl;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
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
import simbir.store.service.task.manager.service.ServiceTaskManager;

import java.sql.Date;
import java.time.LocalDateTime;
import java.util.List;

@Service
@RequiredArgsConstructor
public class ServiceTaskManagerImpl implements ServiceTaskManager {

    private final TasksQueueVendorRegistrationRepository vendorRegRepo;
    private final TasksQueueProductsRegistrationRepository productsRegRepo;

    private final AdminsTasksQueueVendorRegistrationRepository adminsVendorRepo;
    private final AdminsTasksQueueProductsRegistrationRepository adminsProductsRepo;

    // -----------------------------
    // Создание задач
    // -----------------------------

    @Override
    public TaskVendorRegistrationDto createVendorRegistrationTask(TaskVendorRegistrationDto dto) {
        TaskVendorRegistration task = new TaskVendorRegistration();
        task.setTaskName(dto.getTaskName());
        task.setUserId(dto.getUserId());
        task.setDeadline(LocalDateTime.parse(dto.getDeadline()));
        task.setStatus(dto.getStatus());
        task.setCreatedAt(Date.valueOf(dto.getCreatAt()));
        vendorRegRepo.save(task);
        return dto;
    }

    @Override
    public TaskProductsRegistrationDto createProductsRegistrationTask(TaskProductsRegistrationDto dto) {
        TaskProductRegistration task = new TaskProductRegistration();
        task.setTaskName(dto.getTaskName());
        task.setUserId(dto.getUserId());
        task.setProductContainer(dto.getContainer());
        task.setStatus(dto.getStatus());
        task.setDeadline(Date.valueOf(dto.getDeadline()));
        task.setCreatedAt(Date.valueOf(dto.getCreatAt()));
        productsRegRepo.save(task);
        return dto;
    }

    // -----------------------------
    // Получение задач (для админов)
    // -----------------------------

    @Override
    public List<TaskVendorRegistration> getVendorRegistrationTasks() {
        return vendorRegRepo.findAllByOrderByCreatedAtAsc();
    }

    @Override
    public List<TaskProductRegistration> getProductsRegistrationTasks() {
        return productsRegRepo.findAllByOrderByCreatedAtAsc();
    }

    @Override
    public TaskVendorRegistration getVendorRegistrationTaskById(Long taskId) {
        return vendorRegRepo.findById(taskId)
                .orElseThrow(() -> new RuntimeException("Vendor registration task not found"));
    }

    @Override
    public TaskProductRegistration getProductsRegistrationTaskById(Long taskId) {
        return productsRegRepo.findById(taskId)
                .orElseThrow(() -> new RuntimeException("Product registration task not found"));
    }

    // -----------------------------
    // Админ берёт задачу (assign)
    // -----------------------------

    @Override
    public void assignVendorRegistrationTask(Long adminId, Long taskId) {
        AdminTaskVendorRegistration entry = new AdminTaskVendorRegistration();
        entry.setAdminId(adminId);
        entry.setTaskId(taskId);
        adminsVendorRepo.save(entry);

    }

    @Override
    public void assignProductsRegistrationTask(Long adminId, Long taskId) {
        AdminTaskProductRegistration entry = new AdminTaskProductRegistration();
        entry.setAdminId(adminId);
        entry.setTaskId(taskId);
        adminsProductsRepo.save(entry);
    }

    // -----------------------------
    // Обновление статуса задачи
    // -----------------------------

    @Override
    public void updateVendorRegistrationTaskStatus(Long taskId, Long adminId, String status) {
        TaskVendorRegistration task = vendorRegRepo.findById(taskId)
                .orElseThrow(() -> new RuntimeException("Task not found"));
        task.setStatus(status);
        vendorRegRepo.save(task);
    }

    @Override
    public void updateProductsRegistrationTaskStatus(Long taskId, Long adminId, String status) {
        TaskProductRegistration task = productsRegRepo.findById(taskId)
                .orElseThrow(() -> new RuntimeException("Task not found"));
        task.setStatus(status);
        productsRegRepo.save(task);
    }
}

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
import simbir.store.service.task.manager.service.ServiceTaskManager;

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
    public Long createVendorRegistrationTask(Long userId) {
        TaskVendorRegistration task = new TaskVendorRegistration();
        task.setUserId(userId);
        task.setStatus("не обработано");
        vendorRegRepo.save(task);
        return task.getId();
    }

    @Override
    public Long createProductsRegistrationTask(Long userId, String xmlData) {
        TaskProductRegistration task = new TaskProductRegistration();
        task.setUserId(userId);
        task.setProductContainer(xmlData);
        task.setStatus("не обработано");
        productsRegRepo.save(task);
        return task.getId();
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
    public void updateVendorRegistrationTaskStatus(Long taskId, String status) {
        TaskVendorRegistration task = vendorRegRepo.findById(taskId)
                .orElseThrow(() -> new RuntimeException("Task not found"));
        task.setStatus(status);
        vendorRegRepo.save(task);
    }

    @Override
    public void updateProductsRegistrationTaskStatus(Long taskId, String status) {
        TaskProductRegistration task = productsRegRepo.findById(taskId)
                .orElseThrow(() -> new RuntimeException("Task not found"));
        task.setStatus(status);
        productsRegRepo.save(task);
    }
}

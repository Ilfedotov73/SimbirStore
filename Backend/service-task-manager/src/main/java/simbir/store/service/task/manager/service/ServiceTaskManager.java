package simbir.store.service.task.manager.service;

import simbir.store.service.task.manager.db.entity.TaskProductRegistration;
import simbir.store.service.task.manager.db.entity.TaskVendorRegistration;
import simbir.store.service.task.manager.dto.TaskProductsRegistrationDto;
import simbir.store.service.task.manager.dto.TaskVendorRegistrationDto;

import java.util.List;

public interface ServiceTaskManager {

    TaskVendorRegistrationDto createVendorRegistrationTask(TaskVendorRegistrationDto dto);

    TaskProductsRegistrationDto createProductsRegistrationTask(TaskProductsRegistrationDto dto);

    List<TaskVendorRegistration> getVendorRegistrationTasks();

    List<TaskProductRegistration> getProductsRegistrationTasks();

    TaskVendorRegistration getVendorRegistrationTaskById(Long taskId);
    TaskProductRegistration getProductsRegistrationTaskById(Long taskId);

    void assignVendorRegistrationTask(Long adminId, Long taskId);

    void assignProductsRegistrationTask(Long adminId, Long taskId);

    void updateVendorRegistrationTaskStatus(Long taskId, Long adminId, String status);

    void updateProductsRegistrationTaskStatus(Long taskId, Long adminId, String status);
}

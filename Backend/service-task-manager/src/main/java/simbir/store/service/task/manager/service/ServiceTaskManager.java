package simbir.store.service.task.manager.service;

import simbir.store.service.task.manager.db.entity.TaskProductRegistration;
import simbir.store.service.task.manager.db.entity.TaskVendorRegistration;

import java.util.List;

public interface ServiceTaskManager {

    Long createVendorRegistrationTask(Long vendorId);

    Long createProductsRegistrationTask(Long vendorId, String xmlPath);

    List<TaskVendorRegistration> getVendorRegistrationTasks();

    List<TaskProductRegistration> getProductsRegistrationTasks();

    void assignVendorRegistrationTask(Long adminId, Long taskId);

    void assignProductsRegistrationTask(Long adminId, Long taskId);

    void updateVendorRegistrationTaskStatus(Long taskId, String status);

    void updateProductsRegistrationTaskStatus(Long taskId, String status);
}

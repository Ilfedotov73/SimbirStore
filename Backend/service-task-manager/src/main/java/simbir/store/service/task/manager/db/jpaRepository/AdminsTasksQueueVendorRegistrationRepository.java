package simbir.store.service.task.manager.db.jpaRepository;

import org.springframework.data.jpa.repository.JpaRepository;
import simbir.store.service.task.manager.db.entity.AdminTaskVendorRegistration;

public interface AdminsTasksQueueVendorRegistrationRepository
        extends JpaRepository<AdminTaskVendorRegistration, Long> {
}

package simbir.store.service.task.manager.db.jpaRepository;

import org.springframework.data.jpa.repository.JpaRepository;
import simbir.store.service.task.manager.db.entity.AdminTaskProductRegistration;

public interface AdminsTasksQueueProductsRegistrationRepository
        extends JpaRepository<AdminTaskProductRegistration, Long> {
}

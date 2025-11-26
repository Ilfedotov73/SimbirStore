package simbir.store.service.task.manager.db.jpaRepository;

import simbir.store.service.task.manager.db.entity.TaskVendorRegistration;

import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface TasksQueueVendorRegistrationRepository
        extends JpaRepository<TaskVendorRegistration, Long> {

    List<TaskVendorRegistration> findAllByOrderByCreatedAtAsc();
}

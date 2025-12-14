package simbir.store.service.task.manager.db.jpaRepository;

import org.springframework.data.jpa.repository.JpaRepository;
import simbir.store.service.task.manager.db.entity.TaskProductRegistration;

import java.util.List;

public interface TasksQueueProductsRegistrationRepository
        extends JpaRepository<TaskProductRegistration, Long> {

    List<TaskProductRegistration> findAllByOrderByCreatedAtAsc();
}

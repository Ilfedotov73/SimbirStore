package simbir.store.service.vendor.db.jpaRepository;

import org.springframework.data.jpa.repository.JpaRepository;
import simbir.store.service.vendor.db.entity.Notices;

import java.util.List;

public interface NoticesRepository extends JpaRepository<Notices, Integer> {

    List<Notices> findByEntityId(Integer entityId);
}

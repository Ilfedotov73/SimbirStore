package simbir.store.service.task.manager.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.sql.Date;
import java.time.LocalDateTime;

@Entity
@Getter
@Setter
@Table(name = "tasks_queue_vendor_registration")
public class TaskVendorRegistration {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String taskName;

    private Long userId;

    @Enumerated(EnumType.STRING)
    private String status;

    private Date createdAt;

    private LocalDateTime deadline;


}

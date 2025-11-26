package simbir.store.service.task.manager.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.sql.Date;
import java.time.LocalDateTime;

@Entity
@Table(name = "tasks_queue_product_registration")
@Getter
@Setter
public class TaskProductRegistration {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private Long userId;

    private String productContainer;

    @Enumerated(EnumType.STRING)
    private String status;

    private Date createdAt;
    private Date deadline;
}

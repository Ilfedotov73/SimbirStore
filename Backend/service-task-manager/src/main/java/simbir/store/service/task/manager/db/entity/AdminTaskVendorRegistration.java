package simbir.store.service.task.manager.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.util.Date;

@Entity
@Table(name = "admins_tasks_vendor_registration")
@Getter
@Setter
public class AdminTaskVendorRegistration {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private Long adminId;
    private Long taskId;

    private Date assignedAt;
}

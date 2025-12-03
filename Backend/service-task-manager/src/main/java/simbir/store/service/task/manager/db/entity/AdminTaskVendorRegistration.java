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

    public void setAdminId(Long adminId) {
        this.adminId = adminId;
    }

    public void setAssignedAt(Date assignedAt) {
        this.assignedAt = assignedAt;
    }

    public void setTaskId(Long taskId) {
        this.taskId = taskId;
    }

    public void setId(Long id) {
        this.id = id;
    }

    private Long adminId;

    public Long getId() {
        return id;
    }

    public Date getAssignedAt() {
        return assignedAt;
    }

    public Long getTaskId() {
        return taskId;
    }

    public Long getAdminId() {
        return adminId;
    }

    private Long taskId;

    private Date assignedAt;
}

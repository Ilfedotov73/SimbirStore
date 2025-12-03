package simbir.store.service.task.manager.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.util.Date;

@Entity
@Table(name = "admins_tasks_product_registration")
@Getter
@Setter
public class AdminTaskProductRegistration {

    public void setId(Long id) {
        this.id = id;
    }

    public void setAdminId(Long adminId) {
        this.adminId = adminId;
    }

    public void setTaskId(Long taskId) {
        this.taskId = taskId;
    }

    public void setAssignedAt(Date assignedAt) {
        this.assignedAt = assignedAt;
    }

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    public Long getId() {
        return id;
    }

    public Long getTaskId() {
        return taskId;
    }

    public Long getAdminId() {
        return adminId;
    }

    public Date getAssignedAt() {
        return assignedAt;
    }

    private Long adminId;
    private Long taskId;

    private Date assignedAt;
}

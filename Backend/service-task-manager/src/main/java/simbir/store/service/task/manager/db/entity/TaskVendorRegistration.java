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

    public void setTaskName(String taskName) {
        this.taskName = taskName;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }

    public void setDeadline(LocalDateTime deadline) {
        this.deadline = deadline;
    }

    private String taskName;

    public String getTaskName() {
        return taskName;
    }

    public Long getId() {
        return id;
    }

    public Long getUserId() {
        return userId;
    }

    public String getStatus() {
        return status;
    }

    public Date getCreatedAt() {
        return createdAt;
    }

    public LocalDateTime getDeadline() {
        return deadline;
    }

    private Long userId;

    @Enumerated(EnumType.STRING)
    private String status;

    private Date createdAt;

    private LocalDateTime deadline;


}

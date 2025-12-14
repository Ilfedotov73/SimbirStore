package simbir.store.service.task.manager.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.sql.Date;

@Entity
@Table(name = "tasks_queue_product_registration")
@Getter
@Setter
public class TaskProductRegistration {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private Long userId;

    private String taskName;

    private String productContainer;

    @Enumerated(EnumType.STRING)
    private String status;

    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public void setTaskName(String taskName) {
        this.taskName = taskName;
    }

    public void setProductContainer(String productContainer) {
        this.productContainer = productContainer;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public void setDeadline(Date deadline) {
        this.deadline = deadline;
    }

    private Date createdAt;

    public Date getDeadline() {
        return deadline;
    }

    public Long getId() {
        return id;
    }

    public Long getUserId() {
        return userId;
    }

    public String getTaskName() {
        return taskName;
    }

    public String getProductContainer() {
        return productContainer;
    }

    public String getStatus() {
        return status;
    }

    public Date getCreatedAt() {
        return createdAt;
    }

    private Date deadline;
}

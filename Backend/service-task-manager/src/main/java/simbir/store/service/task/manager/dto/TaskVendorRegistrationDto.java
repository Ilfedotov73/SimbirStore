package simbir.store.service.task.manager.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class TaskVendorRegistrationDto {
    private Long id;

    public Long getId() {
        return id;
    }

    public String getDeadline() {
        return deadline;
    }

    public String getTaskName() {
        return taskName;
    }

    public Long getUserId() {
        return userId;
    }

    public String getStatus() {
        return status;
    }

    public String getCreatAt() {
        return creatAt;
    }

    private String taskName;

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public void setTaskName(String taskName) {
        this.taskName = taskName;
    }

    public void setDeadline(String deadline) {
        this.deadline = deadline;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public void setCreatAt(String creatAt) {
        this.creatAt = creatAt;
    }

    private Long userId;
    private String deadline;
    private String status;
    private String creatAt;


}

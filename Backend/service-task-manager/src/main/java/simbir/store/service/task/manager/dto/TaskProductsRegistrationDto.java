package simbir.store.service.task.manager.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class TaskProductsRegistrationDto {
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

    public void setDeadline(String deadline) {
        this.deadline = deadline;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public void setCreatAt(String creatAt) {
        this.creatAt = creatAt;
    }

    public void setContainer(String container) {
        this.container = container;
    }

    private String taskName;

    public Long getId() {
        return id;
    }

    public String getTaskName() {
        return taskName;
    }

    public String getDeadline() {
        return deadline;
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

    public String getContainer() {
        return container;
    }

    private Long userId;
    private String deadline;
    private String status;
    private String creatAt;
    private String container;


    public TaskProductsRegistrationDto() {
    }

    public TaskProductsRegistrationDto(String taskName, Long id, Long userId, String deadline, String status, String creatAt, String container) {
        this.taskName = taskName;
        this.id = id;
        this.userId = userId;
        this.deadline = deadline;
        this.status = status;
        this.creatAt = creatAt;
        this.container = container;
    }
}

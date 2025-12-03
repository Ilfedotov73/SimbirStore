package simbir.store.service.task.manager.dto;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class TaskProductsRegistrationDto {
    private Long Id;
    private String TaskName;
    private Long UserId;
    private String Deadline;
    private String Status;
    private String CreatAt;
    private String Container;
}

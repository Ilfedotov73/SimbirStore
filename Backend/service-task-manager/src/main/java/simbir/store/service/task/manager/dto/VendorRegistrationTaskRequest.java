package simbir.store.service.task.manager.dto;

import lombok.Getter;
import lombok.Setter;

import java.sql.Date;

@Getter
@Setter
public class VendorRegistrationTaskRequest {
    private Long regId;
    private String JsonVendor;
    private Date createAt;
}

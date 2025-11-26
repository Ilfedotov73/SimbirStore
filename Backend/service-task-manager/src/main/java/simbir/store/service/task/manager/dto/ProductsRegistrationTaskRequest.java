package simbir.store.service.task.manager.dto;

import lombok.Getter;
import lombok.Setter;

import java.sql.Date;

@Getter
@Setter
public class ProductsRegistrationTaskRequest {
    private Long vendorId;
    private Date createAt;
    private String xmlProductContainer;
}

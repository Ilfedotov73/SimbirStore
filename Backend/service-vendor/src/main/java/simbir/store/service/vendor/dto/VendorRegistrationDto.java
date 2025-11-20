package simbir.store.service.vendor.dto;

import lombok.Getter;
import lombok.Setter;

import java.sql.Date;

@Getter
@Setter
public class VendorRegistrationDto {
    private int id;
    private String task;
    private String jsonVendor;
    private Date deadline;
    private String status;
    private Date createAt;
}

package simbir.store.service.vendor.dto;

import lombok.Getter;
import lombok.Setter;

import java.sql.Date;

@Getter
@Setter
public class NoticesDto {
    private int id;
    private String text;
    private int entityrId;
    private Date createAt;
}

package simbir.store.service.vendor.dto;

import lombok.Getter;
import lombok.Setter;

import java.sql.Date;

@Getter
@Setter
public class ProductDto {
    private int id;
    private String name;
    private float price;
    private String photoUrl;
    private Date CreateAt;
    private String characteristics;
    private float productRating;
}
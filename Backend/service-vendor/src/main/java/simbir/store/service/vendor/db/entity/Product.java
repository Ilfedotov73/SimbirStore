package simbir.store.service.vendor.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.util.Date;
import java.util.List;

@Entity
@Table(name = "products")
@Getter
@Setter
public class Product {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    private String name;
    private Float price;

    private String photoUrl;

    private Date createAt;

    private String characteristics;

    private Float productRating;

    /** связь с VendorsProducts (m:m через entity) */
    @OneToMany(mappedBy = "product")
    private List<VendorsProducts> vendorLinks;

    /** список отзывов (m:m через ProductsReviews) */
    @OneToMany(mappedBy = "product")
    private List<ProductsReviews> reviews;
}

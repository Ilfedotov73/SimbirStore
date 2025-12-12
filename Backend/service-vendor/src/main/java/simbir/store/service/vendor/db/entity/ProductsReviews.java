package simbir.store.service.vendor.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

@Entity
@Table(name = "products_reviews")
@Getter
@Setter
public class ProductsReviews {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    /** FK → User.Id (Customer) */
    @ManyToOne
    @JoinColumn(name = "customerId")
    private User customer;

    /** FK → Products.Id */
    @ManyToOne
    @JoinColumn(name = "productId")
    private Product product;

    private String review;

    private Float rating;
}

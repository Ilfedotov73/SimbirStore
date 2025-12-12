package simbir.store.service.vendor.db.entity;

import jakarta.persistence.*;
import lombok.Getter;
import lombok.Setter;

import java.sql.Date;
import java.util.List;

@Entity
@Table(name = "users")
@Getter
@Setter
public class User {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    private String firstName;
    private String lastName;
    private String phoneNumber;

    private String photoUrl;

    private String userTelegramId;

    private Date createAt;

    @Column(unique = true)
    private String login;

    private String email;

    private String password;

    /** roles[] на ER → строка */
    private String roles;

    /** Notices: User 1 → m Notices */
    @OneToMany(mappedBy = "user")
    private List<Notices> notices;

    /** Vendor — имеет много продуктов через VendorsProducts */
    @OneToMany(mappedBy = "vendor")
    private List<VendorsProducts> vendorProducts;

    /** Reviews left by this customer */
    @OneToMany(mappedBy = "customer")
    private List<ProductsReviews> reviews;
}

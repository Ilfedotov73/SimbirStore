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

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public void setPhotoUrl(String photoUrl) {
        this.photoUrl = photoUrl;
    }

    public void setPhoneNumber(String phoneNumber) {
        this.phoneNumber = phoneNumber;
    }

    public void setUserTelegramId(String userTelegramId) {
        this.userTelegramId = userTelegramId;
    }

    public void setCreateAt(Date createAt) {
        this.createAt = createAt;
    }

    public void setLogin(String login) {
        this.login = login;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public void setRoles(String roles) {
        this.roles = roles;
    }

    public void setNotices(List<Notices> notices) {
        this.notices = notices;
    }

    public void setVendorProducts(List<VendorsProducts> vendorProducts) {
        this.vendorProducts = vendorProducts;
    }

    public void setReviews(List<ProductsReviews> reviews) {
        this.reviews = reviews;
    }

    public String getFirstName() {
        return firstName;
    }

    public String getLastName() {
        return lastName;
    }

    public String getPhoneNumber() {
        return phoneNumber;
    }

    public String getPhotoUrl() {
        return photoUrl;
    }

    public String getUserTelegramId() {
        return userTelegramId;
    }

    public Date getCreateAt() {
        return createAt;
    }

    public String getLogin() {
        return login;
    }

    public String getEmail() {
        return email;
    }

    public String getPassword() {
        return password;
    }

    public String getRoles() {
        return roles;
    }

    public List<Notices> getNotices() {
        return notices;
    }

    public List<VendorsProducts> getVendorProducts() {
        return vendorProducts;
    }

    public List<ProductsReviews> getReviews() {
        return reviews;
    }

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
